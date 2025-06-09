package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	libCtx "github.com/nocturna-ta/golib/context"
	"github.com/nocturna-ta/golib/custerr"
	"github.com/nocturna-ta/golib/fileutils"
	"github.com/nocturna-ta/golib/log"
	response2 "github.com/nocturna-ta/golib/response"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"github.com/nocturna-ta/ums/internal/interfaces/dao"
	"github.com/nocturna-ta/ums/internal/interfaces/jwtsvc"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
	"github.com/nocturna-ta/ums/pkg/common"
	"github.com/nocturna-ta/ums/pkg/constants"
	"github.com/nocturna-ta/ums/pkg/constants/errorcode"
	"github.com/nocturna-ta/ums/pkg/roles"
	"strings"
	"time"
)

func (m *Module) RegisterUser(ctx context.Context, req *request.UserRegistrationRequest) (*response.UserRegistrationResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCases.RegisterUser")
	defer span.End()

	var (
		user                *model.User
		pendingRegistration *model.PendingRegistration
		KTPPhotoPath        string
	)

	transaction := func(txCtx context.Context) (any, error) {
		user = model.ConstructUserRegistration(req)

		errTx := m.userRepo.Insert(txCtx, user)
		if errTx != nil {
			if errors.Is(errTx, dao.ErrDuplicate) {
				return nil, &custerr.ErrChain{
					Message: errorcode.UserAlreadyExists.Message,
					Code:    errorcode.UserAlreadyExists.Code,
					Type:    response2.ErrBadRequest,
					Cause:   errTx,
				}
			}

			return nil, errTx
		}

		var entityData interface{}
		switch req.Role {
		case roles.RoleKPUProvinsi:
			entityData = &model.KPUProvinsiData{
				Username:  req.KPUUsername,
				Name:      req.KPUName,
				Address:   req.Address,
				Region:    req.Region,
				Telephone: req.Telephone,
			}
		case roles.RoleKPUKota:
			entityData = &model.KPUKotaData{
				Username:  req.KPUUsername,
				Name:      req.KPUName,
				Address:   req.Address,
				Region:    req.Region,
				Telephone: req.Telephone,
			}
		case roles.RoleVoter:
			if req.KTPPhotoName == "" || req.KTPPhotoFile == nil {
				fileConfigKTP := fileutils.DefaultConfig()
				fileConfigKTP.SetAllowedImageExtension()
				fileConfigKTP.EntityType = "KTP"

				KTPPhotoPath, errTx = fileutils.StoreFile(txCtx, req.KTPPhotoFile, req.KTPPhotoName, fileConfigKTP)
				if errTx != nil {
					_ = fileutils.DeleteFile(txCtx, KTPPhotoPath)
					log.WithFields(log.Fields{
						"error": errTx,
					}).ErrorWithCtx(txCtx, "[UserUseCases.RegisterUser] Failed to store KTP photo")
					return nil, &custerr.ErrChain{
						Message: "Failed to store KTP photo",
						Code:    500,
						Type:    response2.ErrInternalServerError,
						Cause:   errTx,
					}
				}
			} else {
				KTPPhotoPath = ""
			}

			entityData = &model.VoterData{
				NIK:          req.NIK,
				Fullname:     req.FullName,
				Gender:       req.Gender,
				BirthPlace:   req.BirthPlace,
				BirthDate:    req.BirthDate,
				Residential:  req.ResidentialAddress,
				VoterAddress: req.Address,
				Region:       req.Region,
				KTPPhotoPath: KTPPhotoPath,
			}
		}

		var errReg error
		pendingRegistration, errReg = model.NewPendingRegist(
			user.ID,
			req.Role,
			entityData,
		)
		if errReg != nil {
			return nil, &custerr.ErrChain{
				Message: "Failed to create pending registration",
				Code:    500,
				Type:    response2.ErrInternalServerError,
				Cause:   errReg,
			}
		}

		errTx = m.pendingRegRepo.Insert(txCtx, pendingRegistration)
		if errTx != nil {
			if errors.Is(errTx, dao.ErrDuplicate) {
				_ = fileutils.DeleteFile(txCtx, KTPPhotoPath)
				return nil, &custerr.ErrChain{
					Message: "Pending registration already exists",
					Code:    400,
					Type:    response2.ErrBadRequest,
					Cause:   errTx,
				}
			}
			return nil, errTx
		}

		errTx = m.publisher.Publish(txCtx, m.topics.MasterDataUser.Value, user.ID.String(), user.ToMessageModel(), map[string]any{
			constants.MetaDataOperation: constants.Create,
		})

		if errTx != nil {
			return nil, errTx
		}

		return nil, nil

	}
	_, err := m.txMgr.Execute(ctx, transaction, nil)
	if err != nil {
		return nil, err
	}

	return &response.UserRegistrationResponse{
		ID:                 user.ID.String(),
		Email:              user.Email,
		VerificationStatus: user.VerificationStatus,
		RequestedRole:      user.RequestedRole,
		Message:            "Registration successful. Your account is pending verification by an admin.",
	}, nil
}

func (m *Module) GetUserByEmail(ctx context.Context, email string) (*response.UserResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCases.GetUserByEmail")
	defer span.End()

	user, err := m.userRepo.GetByEmail(ctx, email)
	if err != nil {
		log.WithFields(log.Fields{
			"email": email,
			"error": err,
		}).ErrorWithCtx(ctx, "[UserUseCases.GetUserByEmail] Failed to get user by email")
		return nil, err
	}

	return &response.UserResponse{
		ID:    user.ID.String(),
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

func (m *Module) GetByID(ctx context.Context) (*response.UserResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCases.GetByID")
	defer span.End()

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := m.userRepo.GetById(ctx, reqCtx.GetUserId())
	if err != nil {
		log.WithFields(log.Fields{
			"id":    reqCtx.UserId,
			"error": err,
		}).ErrorWithCtx(ctx, "[UserUseCases.GetByID] Failed to get user by id")
		return nil, err
	}

	return &response.UserResponse{
		ID:    user.ID.String(),
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

func (m *Module) ChangePassword(ctx context.Context, req *request.UserChangePasswordRequest) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCases.ChangePassword")
	defer span.End()

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return err
	}

	existing, err := m.userRepo.GetById(ctx, reqCtx.GetUserId())
	if err != nil {
		log.WithFields(log.Fields{
			"id":    reqCtx.UserId,
			"error": err,
		}).ErrorWithCtx(ctx, "[UserUseCases.ChangePassword] Failed to get user by id")
		return err
	}

	oldHash := common.PasswordHash(req.Old, existing.PasswordSalt)
	newPassword := common.PasswordHash(req.Confirm, existing.PasswordSalt)
	if oldHash != existing.Password {
		return &custerr.ErrChain{
			Message: errorcode.WrongPassword.Message,
			Code:    errorcode.WrongPassword.Code,
		}
	}

	err = m.userRepo.ChangePassword(ctx, existing.ID, newPassword)

	if err != nil {
		log.WithFields(log.Fields{
			"id":      reqCtx.UserId,
			"err":     err,
			"request": req,
		}).ErrorWithCtx(ctx, "[UserUseCases.ChangePassword] Failed to change password")
		return err
	}

	return nil
}

func (m *Module) LoginUser(ctx context.Context, req *request.UserLoginRequest) (*response.UserLoginResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCases.LoginUser")
	defer span.End()

	existing, err := m.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		log.WithFields(log.Fields{
			"request": req,
			"error":   err,
		}).ErrorWithCtx(ctx, "[UserUseCases.LoginUser] Failed to get user by email")
		return nil, err
	}

	password := common.PasswordHash(req.Password, existing.PasswordSalt)
	if password != existing.Password {
		return nil, &custerr.ErrChain{
			Message: errorcode.WrongPassword.Message,
			Code:    errorcode.WrongPassword.Code,
		}
	}

	response := &response.UserLoginResponse{
		IsActive:           existing.IsActive,
		VerificationStatus: existing.VerificationStatus,
		RequestedRole:      existing.RequestedRole,
	}

	switch existing.VerificationStatus {
	case model.VerificationStatusPending:
		response.Message = "Your account is pending verification by an admin. Please wait for approval."
	case model.VerificationStatusRejected:
		response.Message = "Your account verification was rejected."
	}

	var publicAddress string
	switch existing.Role {
	case roles.RoleKPUProvinsi:
		kpuProvinsi, err := m.kpuProvinsiRepo.GetKPUProvinsiByUserID(ctx, existing.ID)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
				"id":    existing.ID,
			}).ErrorWithCtx(ctx, "[UserUseCases.LoginUser] Failed to get KPU Provinsi by user ID")
			return nil, err
		}
		publicAddress = kpuProvinsi.Address
	case roles.RoleVoter:
		voter, err := m.voterRepo.GetVoterByUserID(ctx, existing.ID)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
				"id":    existing.ID,
			}).ErrorWithCtx(ctx, "[UserUseCases.LoginUser] Failed to get voter by user ID")
			return nil, err
		}
		publicAddress = voter.VoterAddress
	case roles.RoleKPUKota:
		kpuKota, err := m.kpuKotaRepo.GetKPUKotaByUserID(ctx, existing.ID)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
				"id":    existing.ID,
			}).ErrorWithCtx(ctx, "[UserUseCases.LoginUser] Failed to get KPU Kota by user ID")
			return nil, err
		}
		publicAddress = kpuKota.Address
	}

	if existing.IsActive {
		expiresAt := jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour))
		token, err := m.jwtSvc.GenerateToken(ctx, jwtsvc.AccessClaims{
			RegisteredClaims: &jwt.RegisteredClaims{
				ExpiresAt: expiresAt,
				ID:        uuid.New().String(),
			},
			JwtData: &jwtsvc.JwtData{
				UserID:        existing.ID.String(),
				Role:          existing.Role,
				PublicAddress: publicAddress,
			},
		})

		if err != nil {
			log.WithFields(log.Fields{
				"error":   err,
				"request": req,
			}).ErrorWithCtx(ctx, "[UserUseCases.LoginUser] Failed to generate token")
			return nil, err
		}

		response.Token = token
		response.ExpiresAt = &expiresAt.Time
		response.Message = "Login successful."
	} else {

	}

	return response, nil
}

func (m *Module) GetPendingVerificationUsers(ctx context.Context) (*[]response.UserVerificationResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCases.GetPendingVerificationUsers")
	defer span.End()

	users, err := m.userRepo.GetPendingVerificationUsers(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserUseCases.GetPendingVerificationUsers] Failed to get pending verification users")
		return nil, err
	}

	var result []response.UserVerificationResponse
	for _, user := range users {
		result = append(result, response.UserVerificationResponse{
			ID:                 user.ID.String(),
			Email:              user.Email,
			RequestedRole:      user.RequestedRole,
			VerificationStatus: user.VerificationStatus,
			CreatedAt:          user.CreatedAt,
		})
	}

	return &result, nil
}

func (m *Module) GetVerificationDetails(ctx context.Context, userIDStr string) (*response.UserVerificationDetailsResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCases.GetVerificationDetails")
	defer span.End()

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, &custerr.ErrChain{
			Message: errorcode.InvalidUUID.Message,
			Code:    errorcode.InvalidUUID.Code,
			Type:    response2.ErrBadRequest,
		}
	}

	user, err := m.userRepo.GetById(ctx, userID)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    userID,
		}).ErrorWithCtx(ctx, "[UserUseCases.GetVerificationDetails] Failed to get user by id")
		return nil, err
	}

	if user.VerificationStatus != model.VerificationStatusPending {
		return nil, &custerr.ErrChain{
			Message: "User is not in pending verification status",
			Code:    400,
			Type:    response2.ErrBadRequest,
		}
	}

	pendingReg, err := m.pendingRegRepo.GetByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, dao.ErrNoResult) {
			return &response.UserVerificationDetailsResponse{
				ID:                 user.ID.String(),
				Email:              user.Email,
				RequestedRole:      user.RequestedRole,
				VerificationStatus: user.VerificationStatus,
				CreatedAt:          user.CreatedAt,
			}, nil
		}

		log.WithFields(log.Fields{
			"error":  err,
			"userID": userID,
		}).ErrorWithCtx(ctx, "[UserUseCases.GetVerificationDetails] Failed to get pending registration")
		return nil, err
	}

	var entityData map[string]interface{}
	if err := json.Unmarshal(pendingReg.EntityData, &entityData); err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    userID,
		}).ErrorWithCtx(ctx, "[UserUseCases.GetVerificationDetails] Failed to unmarshal entity data")
		return nil, &custerr.ErrChain{
			Message: "Failed to parse entity data",
			Code:    500,
			Type:    response2.ErrInternalServerError,
			Cause:   err,
		}
	}

	return &response.UserVerificationDetailsResponse{
		ID:                 user.ID.String(),
		Email:              user.Email,
		RequestedRole:      user.RequestedRole,
		VerificationStatus: user.VerificationStatus,
		CreatedAt:          user.CreatedAt,
		EntityData:         entityData,
	}, nil
}

func (m *Module) ApproveUserVerification(ctx context.Context, req *request.UserVerificationRequest) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCases.ApproveUserVerification")
	defer span.End()

	err := req.ValidateVerificationRequest()
	if err != nil {
		return err
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return &custerr.ErrChain{
			Message: errorcode.InvalidUUID.Message,
			Code:    errorcode.InvalidUUID.Code,
			Type:    response2.ErrBadRequest,
		}
	}

	user, err := m.userRepo.GetById(ctx, userID)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    userID,
		}).ErrorWithCtx(ctx, "[UserUseCases.ApproveUserVerification] Failed to get user by id")
		return err
	}

	if user.VerificationStatus != model.VerificationStatusPending {
		return &custerr.ErrChain{
			Message: "User is not in pending verification status",
			Code:    400,
			Type:    response2.ErrBadRequest,
		}
	}

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return &custerr.ErrChain{
			Message: "Failed to get request context",
			Code:    500,
			Type:    response2.ErrInternalServerError,
		}
	}

	verifierRole := reqCtx.GetRole()

	if !roles.CanVerify(verifierRole, user.RequestedRole) {
		return &custerr.ErrChain{
			Message: fmt.Sprintf("A %s cannot verify a %s. Please follow the verification hierarchy.",
				roles.GetRoleDisplayName(verifierRole),
				roles.GetRoleDisplayName(user.RequestedRole)),
			Code: 403,
			Type: response2.ErrForbiddenResource,
		}
	}

	pendingReg, err := m.pendingRegRepo.GetByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, dao.ErrNoResult) {
			return &custerr.ErrChain{
				Message: "No pending registration data found for this user",
				Code:    404,
				Type:    response2.ErrNotFound,
			}
		}
		log.WithFields(log.Fields{
			"error":  err,
			"userID": userID,
		}).ErrorWithCtx(ctx, "[UserUseCases.ApproveUserVerification] Failed to get pending registration")
		return err
	}

	if !roles.IsValidRole(pendingReg.Role) {
		return &custerr.ErrChain{
			Message: "Invalid requested role",
			Code:    400,
			Type:    response2.ErrBadRequest,
		}
	}

	transaction := func(txCtx context.Context) (any, error) {
		errTx := m.userRepo.UpdateVerificationStatus(txCtx, userID, model.VerificationStatusApproved, pendingReg.Role)
		if errTx != nil {
			return nil, errTx
		}

		switch pendingReg.Role {
		case roles.RoleKPUProvinsi:
			kpuProvinsiData, errData := pendingReg.GetKPUProvinsiData()
			if errData != nil {
				return nil, &custerr.ErrChain{
					Message: "Failed to extract KPU Provinsi data",
					Code:    500,
					Type:    response2.ErrInternalServerError,
					Cause:   errData,
				}
			}

			kpuProvinsi := &model.KPUProvinsi{
				BaseModel: model.BaseModel{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					IsDeleted: false,
				},
				ID:           uuid.New(),
				UserID:       userID,
				Username:     kpuProvinsiData.Username,
				Name:         kpuProvinsiData.Name,
				Address:      kpuProvinsiData.Address,
				Region:       kpuProvinsiData.Region,
				Telephone:    kpuProvinsiData.Telephone,
				IsActive:     true,
				RegisteredAt: time.Now(),
			}

			errTx := m.kpuProvinsiRepo.InsertKPUProvinsi(txCtx, kpuProvinsi)
			if errTx != nil {
				return nil, errTx
			}

			txHash, errTx := m.kpuProvinsiRepo.SendTxKPUProvinsiBlockchain(txCtx, req.SignedTransaction)

			if txHash == "" {
				return nil, err
			}

			errTx = m.publisher.Publish(txCtx, m.topics.MasterDataKPUProvinsi.Value, kpuProvinsi.ID.String(), kpuProvinsi.ToMessageModel(), map[string]any{
				constants.MetaDataOperation: constants.Create,
			})
			if errTx != nil {
				return nil, errTx
			}

		case roles.RoleKPUKota:
			kpuKotaData, errData := pendingReg.GetKPUKotaData()
			if errData != nil {
				return nil, &custerr.ErrChain{
					Message: "Failed to extract KPU Kota data",
					Code:    500,
					Type:    response2.ErrInternalServerError,
					Cause:   errData,
				}
			}

			kpuKota := &model.KPUKota{
				BaseModel: model.BaseModel{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					IsDeleted: false,
				},
				ID:           uuid.New(),
				UserID:       userID,
				Username:     kpuKotaData.Username,
				Name:         kpuKotaData.Name,
				Address:      kpuKotaData.Address,
				Region:       kpuKotaData.Region,
				Telephone:    kpuKotaData.Telephone,
				IsActive:     true,
				RegisteredAt: time.Now(),
			}

			errTx := m.kpuKotaRepo.InsertKPUKota(txCtx, kpuKota)
			if errTx != nil {
				return nil, errTx
			}

			txHash, errTx := m.kpuKotaRepo.SendTxKPUKotaBlockchain(txCtx, req.SignedTransaction)
			if errTx != nil {
				return nil, errTx
			}

			if txHash == "" {
				return nil, err
			}

			errTx = m.publisher.Publish(txCtx, m.topics.MasterDataKPUKota.Value, kpuKota.ID.String(), kpuKota.ToMessageModel(), map[string]any{
				constants.MetaDataOperation: constants.Create,
			})
			if errTx != nil {
				return nil, errTx
			}

		case roles.RoleVoter:
			voterData, errData := pendingReg.GetVoterData()
			if errData != nil {
				return nil, &custerr.ErrChain{
					Message: "Failed to extract Voter data",
					Code:    500,
					Type:    response2.ErrInternalServerError,
					Cause:   errData,
				}
			}

			layout := "2006-01-02"
			now := time.Now()
			birthDate, err := time.Parse(layout, voterData.BirthDate)
			if err != nil {
				return nil, err
			}

			voter := &model.Voter{
				BaseModel: model.BaseModel{
					CreatedAt: now,
					UpdatedAt: time.Now(),
					IsDeleted: false,
				},
				ID:                 uuid.New(),
				UserID:             userID,
				NIK:                voterData.NIK,
				FullName:           voterData.Fullname,
				Gender:             voterData.Gender,
				BirthPlace:         voterData.BirthPlace,
				BirthDate:          birthDate,
				ResidentialAddress: voterData.Residential,
				VoterAddress:       voterData.VoterAddress,
				Region:             voterData.Region,
				KTPPhotoPath:       voterData.KTPPhotoPath,
				IsRegistered:       true,
				HasVoted:           false,
				VotedAt:            nil,
				LastLogin:          &now,
			}

			errTx := m.voterRepo.InsertVoter(txCtx, voter)
			if errTx != nil {
				if errors.Is(errTx, dao.ErrDuplicate) {
					_ = fileutils.DeleteFile(txCtx, voter.KTPPhotoPath)
					return nil, &custerr.ErrChain{
						Message: "User already exists",
						Code:    400,
						Type:    response2.ErrBadRequest,
						Cause:   errTx,
					}
				}
				return nil, errTx
			}

			txHash, errTx := m.voterRepo.SendTxVoterBlockchain(txCtx, req.SignedTransaction)
			if errTx != nil {
				return nil, errTx
			}

			if txHash == "" {
				return nil, err
			}

			errTx = m.publisher.Publish(txCtx, m.topics.MasterDataVoter.Value, voter.ID.String(), voter.ToMessageModel(), map[string]any{
				constants.MetaDataOperation: constants.Create,
			})
			if errTx != nil {
				return nil, errTx
			}
		}

		errTx = m.pendingRegRepo.Delete(txCtx, pendingReg.ID)
		if errTx != nil {
			log.WithFields(log.Fields{
				"error": errTx,
				"id":    pendingReg.ID,
			}).WarnWithCtx(txCtx, "[UserUseCases.ApproveUserVerification] Failed to delete pending registration")
		}

		errTx = m.publisher.Publish(txCtx, m.topics.MasterDataUser.Value, userID.String(), user.ToMessageModel(), map[string]any{
			constants.MetaDataOperation: constants.Update,
		})
		if errTx != nil {
			return nil, errTx
		}

		return nil, nil
	}

	_, err = m.txMgr.Execute(ctx, transaction, nil)
	return err
}

func (m *Module) RejectUserVerification(ctx context.Context, req *request.UserVerificationRequest) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCases.RejectUserVerification")
	defer span.End()

	err := req.ValidateVerificationRequest()
	if err != nil {
		return err
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return &custerr.ErrChain{
			Message: errorcode.InvalidUUID.Message,
			Code:    errorcode.InvalidUUID.Code,
			Type:    response2.ErrBadRequest,
		}
	}

	user, err := m.userRepo.GetById(ctx, userID)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    userID,
		}).ErrorWithCtx(ctx, "[UserUseCases.RejectUserVerification] Failed to get user by id")
		return err
	}

	if user.VerificationStatus != model.VerificationStatusPending {
		return &custerr.ErrChain{
			Message: "User is not in pending verification status",
			Code:    400,
			Type:    response2.ErrBadRequest,
		}
	}

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return &custerr.ErrChain{
			Message: "Failed to get request context",
			Code:    500,
			Type:    response2.ErrInternalServerError,
		}
	}

	verifierRole := reqCtx.GetRole()

	if !roles.CanVerify(verifierRole, user.RequestedRole) {
		return &custerr.ErrChain{
			Message: fmt.Sprintf("A %s cannot verify a %s. Please follow the verification hierarchy.",
				roles.GetRoleDisplayName(verifierRole),
				roles.GetRoleDisplayName(user.RequestedRole)),
			Code: 403,
			Type: response2.ErrForbiddenResource,
		}
	}

	pendingReg, err := m.pendingRegRepo.GetByUserID(ctx, userID)
	if err != nil && !errors.Is(err, dao.ErrNoResult) {
		log.WithFields(log.Fields{
			"error":  err,
			"userID": userID,
		}).ErrorWithCtx(ctx, "[UserUseCases.RejectUserVerification] Failed to get pending registration")
	}

	transaction := func(txCtx context.Context) (any, error) {
		errTx := m.userRepo.UpdateVerificationStatus(txCtx, userID, model.VerificationStatusRejected, "unverified")
		if errTx != nil {
			return nil, errTx
		}

		if pendingReg != nil {
			errTx = m.pendingRegRepo.Delete(txCtx, pendingReg.ID)
			if errTx != nil {
				log.WithFields(log.Fields{
					"error": errTx,
					"id":    pendingReg.ID,
				}).WarnWithCtx(txCtx, "[UserUseCases.RejectUserVerification] Failed to delete pending registration")
			}
		}

		metadata := map[string]any{
			constants.MetaDataOperation: constants.Update,
		}

		if req.AdminReason != "" {
			metadata["rejection_reason"] = req.AdminReason
		}

		errTx = m.publisher.Publish(txCtx, m.topics.MasterDataUser.Value, user.ID.String(), user.ToMessageModel(), metadata)
		if errTx != nil {
			return nil, errTx
		}

		return nil, nil
	}

	_, err = m.txMgr.Execute(ctx, transaction, nil)
	return err
}

func (m *Module) CheckVerificationStatus(ctx context.Context, email string) (*response.UserVerificationResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCases.CheckVerificationStatus")
	defer span.End()

	user, err := m.userRepo.GetByEmail(ctx, email)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"email": email,
		}).ErrorWithCtx(ctx, "[UserUseCases.CheckVerificationStatus] Failed to get user by email")
		return nil, err
	}

	return &response.UserVerificationResponse{
		ID:                 user.ID.String(),
		Email:              user.Email,
		RequestedRole:      user.RequestedRole,
		VerificationStatus: user.VerificationStatus,
		CreatedAt:          user.CreatedAt,
	}, nil
}

func (m *Module) GetMyVerificationStatus(ctx context.Context) (*response.UserVerificationStatusResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCases.GetMyVerificationStatus")
	defer span.End()

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return nil, err
	}

	userID := reqCtx.GetUserId()

	user, err := m.userRepo.GetById(ctx, userID)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    userID,
		}).ErrorWithCtx(ctx, "[UserUseCases.GetMyVerificationStatus] Failed to get user by id")
		return nil, err
	}

	response := &response.UserVerificationStatusResponse{
		Email:              user.Email,
		RequestedRole:      user.RequestedRole,
		Role:               user.Role,
		VerificationStatus: user.VerificationStatus,
		IsActive:           user.IsActive,
		CreatedAt:          user.CreatedAt,
	}

	switch user.VerificationStatus {
	case model.VerificationStatusPending:
		response.Message = "Your account is pending verification by an admin. Please wait for approval."
	case model.VerificationStatusApproved:
		response.Message = "Your account has been verified and is active."
	case model.VerificationStatusRejected:
		response.Message = "Your account verification was rejected."
	}

	return response, nil
}

func (m *Module) GetPendingVerificationsByRole(ctx context.Context) (*[]response.UserVerificationResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCases.GetPendingVerificationsByRole")
	defer span.End()

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return nil, &custerr.ErrChain{
			Message: "Failed to get request context",
			Code:    500,
			Type:    response2.ErrInternalServerError,
		}
	}

	currentRole := reqCtx.GetRole()
	currentUserID := reqCtx.GetUserId()

	users, err := m.userRepo.GetPendingVerificationUsers(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserUseCases.GetPendingVerificationsByRole] Failed to get pending verification users")
		return nil, err
	}

	var filteredUsers []model.User

	if currentRole == roles.RoleKPUProvinsi {
		kpuProvinsi, err := m.kpuProvinsiRepo.GetKPUProvinsiByUserID(ctx, currentUserID)
		if err != nil {
			log.WithFields(log.Fields{
				"error":  err,
				"userID": currentUserID,
			}).ErrorWithCtx(ctx, "[UserUseCases.GetPendingVerificationsByRole] Failed to get KPU Provinsi data")
			return nil, err
		}

		for _, user := range users {
			if roles.CanVerify(currentRole, user.RequestedRole) {
				if user.RequestedRole == roles.RoleKPUKota {
					pendingReg, err := m.pendingRegRepo.GetByUserID(ctx, user.ID)
					if err != nil {
						log.WithFields(log.Fields{
							"error":  err,
							"userID": user.ID,
						}).WarnWithCtx(ctx, "[UserUseCases.GetPendingVerificationsByRole] Failed to get pending registration for region check")
						continue
					}

					var entityData map[string]interface{}
					if err := json.Unmarshal(pendingReg.EntityData, &entityData); err != nil {
						log.WithFields(log.Fields{
							"error": err,
							"id":    user.ID,
						}).WarnWithCtx(ctx, "[UserUseCases.GetPendingVerificationsByRole] Failed to unmarshal entity data")
						continue
					}

					kotaRegion, ok := entityData["region"].(string)
					if !ok {
						log.WithFields(log.Fields{
							"userID": user.ID,
						}).WarnWithCtx(ctx, "[UserUseCases.GetPendingVerificationsByRole] Region not found in entity data")
						continue
					}

					isAllowed, err := m.isRegionAllowedByWilayahAPI(ctx, kpuProvinsi.Region, kotaRegion)
					if err != nil {
						log.WithFields(log.Fields{
							"error":    err,
							"province": kpuProvinsi.Region,
							"city":     kotaRegion,
							"userID":   user.ID,
						}).WarnWithCtx(ctx, "[UserUseCases.GetPendingVerificationsByRole] Failed to check region via wilayah.id API")

						filteredUsers = append(filteredUsers, user)
						continue
					}

					if isAllowed {
						filteredUsers = append(filteredUsers, user)
					} else {
						log.WithFields(log.Fields{
							"province":  kpuProvinsi.Region,
							"city":      kotaRegion,
							"userID":    user.ID,
							"userEmail": user.Email,
						}).InfoWithCtx(ctx, "[UserUseCases.GetPendingVerificationsByRole] User filtered out due to region mismatch")
					}
				} else {
					filteredUsers = append(filteredUsers, user)
				}
			}
		}
	} else {
		for _, user := range users {
			if roles.CanVerify(currentRole, user.RequestedRole) {
				filteredUsers = append(filteredUsers, user)
			}
		}
	}

	var result []response.UserVerificationResponse
	for _, user := range filteredUsers {
		pendingReg, err := m.pendingRegRepo.GetByUserID(ctx, user.ID)
		if err != nil {
			log.WithFields(log.Fields{
				"error":  err,
				"userID": user.ID,
			}).ErrorWithCtx(ctx, "[UserUseCases.GetPendingVerificationsByRole] Failed to get pending registration")
			return nil, err
		}

		var entityData map[string]interface{}
		if err := json.Unmarshal(pendingReg.EntityData, &entityData); err != nil {
			log.WithFields(log.Fields{
				"error": err,
				"id":    user.ID,
			}).ErrorWithCtx(ctx, "[UserUseCases.GetPendingVerificationsByRole] Failed to unmarshal entity data")
			return nil, &custerr.ErrChain{
				Message: "Failed to parse entity data",
				Code:    500,
				Type:    response2.ErrInternalServerError,
				Cause:   err,
			}
		}

		result = append(result, response.UserVerificationResponse{
			ID:                 user.ID.String(),
			Email:              user.Email,
			EntityData:         entityData,
			RequestedRole:      user.RequestedRole,
			VerificationStatus: user.VerificationStatus,
			CreatedAt:          user.CreatedAt,
		})
	}

	log.WithFields(log.Fields{
		"totalUsers":    len(users),
		"filteredUsers": len(filteredUsers),
		"currentRole":   currentRole,
		"currentUserID": currentUserID,
	}).InfoWithCtx(ctx, "[UserUseCases.GetPendingVerificationsByRole] Successfully filtered pending verifications")

	return &result, nil
}

func (m *Module) GetEnhancedVerificationStatus(ctx context.Context) (*response.EnhancedUserVerificationStatusResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCases.GetEnhancedVerificationStatus")
	defer span.End()

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return nil, err
	}

	userID := reqCtx.GetUserId()

	user, err := m.userRepo.GetById(ctx, userID)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    userID,
		}).ErrorWithCtx(ctx, "[UserUseCases.GetEnhancedVerificationStatus] Failed to get user by id")
		return nil, err
	}

	response := &response.EnhancedUserVerificationStatusResponse{
		Email:              user.Email,
		RequestedRole:      user.RequestedRole,
		Role:               user.Role,
		VerificationStatus: user.VerificationStatus,
		IsActive:           user.IsActive,
		CreatedAt:          user.CreatedAt,
		VerifierRole:       roles.GetVerifierRoleFor(user.RequestedRole),
		HierarchyLevel:     roles.GetRoleLevel(user.Role),
	}

	switch user.VerificationStatus {
	case model.VerificationStatusPending:
		verifierRole := roles.GetVerifierRoleFor(user.RequestedRole)
		if verifierRole != "" {
			response.Message = fmt.Sprintf("Your account is pending verification by %s. Please wait for approval.",
				roles.GetRoleDisplayName(verifierRole))
		} else {
			response.Message = "Your account is pending verification. Please wait for approval."
		}
	case model.VerificationStatusApproved:
		response.Message = "Your account has been verified and is active."
	case model.VerificationStatusRejected:
		response.Message = "Your account verification was rejected."
	}

	return response, nil
}

func (m *Module) isRegionAllowedByWilayahAPI(ctx context.Context, provinceName, regencyName string) (bool, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCases.isRegionAllowedByWilayahAPI")
	defer span.End()

	if !m.regionCache.IsExpired() {
		if allowed, found := m.checkRegionInCache(provinceName, regencyName); found {
			log.WithFields(log.Fields{
				"province": provinceName,
				"regency":  regencyName,
				"allowed":  allowed,
				"source":   "cache",
			}).InfoWithCtx(ctx, "[UserUseCases.isRegionAllowedByWilayahAPI] Using cached result")
			return allowed, nil
		}
	}

	allowed, err := m.wilayahAPIClient.IsRegencyInProvince(ctx, provinceName, regencyName)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err,
			"province": provinceName,
			"regency":  regencyName,
		}).ErrorWithCtx(ctx, "[UserUseCases.isRegionAllowedByWilayahAPI] API call failed")

		return true, nil
	}

	go m.updateRegionCache(context.Background(), provinceName)

	log.WithFields(log.Fields{
		"province": provinceName,
		"regency":  regencyName,
		"allowed":  allowed,
		"source":   "api",
	}).InfoWithCtx(ctx, "[UserUseCases.isRegionAllowedByWilayahAPI] Using API result")

	return allowed, nil
}

func (m *Module) checkRegionInCache(provinceName, regencyName string) (bool, bool) {
	provinceCode, provinceExists := m.regionCache.GetProvinceCode(provinceName)
	if !provinceExists {
		return false, false
	}

	regencies, regenciesExist := m.regionCache.GetRegencies(provinceCode)
	if !regenciesExist {
		return false, false
	}

	normalizedRegencyName := normalizeRegionName(regencyName)
	for _, regency := range regencies {
		if normalizeRegionName(regency) == normalizedRegencyName {
			return true, true
		}
	}

	return false, true
}

func (m *Module) updateRegionCache(ctx context.Context, provinceName string) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCases.updateRegionCache")
	defer span.End()

	provinceCode, err := m.wilayahAPIClient.FindProvinceByName(ctx, provinceName)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err,
			"province": provinceName,
		}).ErrorWithCtx(ctx, "[UserUseCases.updateRegionCache] Failed to find province code")
		return
	}

	regencies, err := m.wilayahAPIClient.GetRegenciesByProvinceCode(ctx, provinceCode)
	if err != nil {
		log.WithFields(log.Fields{
			"error":        err,
			"province":     provinceName,
			"provinceCode": provinceCode,
		}).ErrorWithCtx(ctx, "[UserUseCases.updateRegionCache] Failed to get regencies")
		return
	}

	m.regionCache.SetProvinceMapping(provinceName, provinceCode)

	regencyNames := make([]string, len(regencies))
	for i, regency := range regencies {
		regencyNames[i] = regency.Name
	}
	m.regionCache.SetRegencies(provinceCode, regencyNames)

	log.WithFields(log.Fields{
		"province":     provinceName,
		"provinceCode": provinceCode,
		"regencyCount": len(regencyNames),
	}).InfoWithCtx(ctx, "[UserUseCases.updateRegionCache] Cache updated successfully")
}

func normalizeRegionName(name string) string {
	normalized := strings.TrimSpace(name)
	normalized = strings.ToLower(normalized)

	prefixes := []string{"kota ", "kabupaten ", "kab. ", "kab ", "kot. ", "kot ", "provinsi "}
	for _, prefix := range prefixes {
		if strings.HasPrefix(normalized, prefix) {
			normalized = strings.TrimPrefix(normalized, prefix)
			break
		}
	}

	return normalized
}
