package dao

import (
	"context"
	sql2 "database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/golib/txmanager/utils"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"github.com/nocturna-ta/ums/pkg/constants"
	"net/url"
	"time"
)

type UserRepository struct {
	db *sql.Store
}

type OptsUserRepository struct {
	DB *sql.Store
}

func NewUserRepository(opts *OptsUserRepository) repository.UserRepository {
	return &UserRepository{
		db: opts.DB,
	}
}

const (
	insertUser = `INSERT INTO users (id, email, password, password_salt, role, requested_role, is_active, verification_status, created_at, updated_at) 
                 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	selectUser = `SELECT %s FROM users WHERE TRUE %s`
	updateUser = `UPDATE users SET %s WHERE TRUE %s`
)

func (repo *UserRepository) Insert(ctx context.Context, user *model.User) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserRepository.Insert")
	defer span.End()

	var (
		err error
	)

	sqlTrx := utils.GetSqlTx(ctx)

	if sqlTrx != nil {
		_, err = sqlTrx.ExecContext(ctx, insertUser,
			user.ID,
			user.Email,
			user.Password,
			user.PasswordSalt,
			user.Role,
			user.RequestedRole,
			user.IsActive,
			user.VerificationStatus,
			user.CreatedAt,
			user.UpdatedAt)
	} else {
		_, err = repo.db.GetMaster().ExecContext(ctx, insertUser,
			user.ID,
			user.Email,
			user.Password,
			user.PasswordSalt,
			user.Role,
			user.RequestedRole,
			user.IsActive,
			user.VerificationStatus,
			user.CreatedAt,
			user.UpdatedAt)
	}

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code {
			case "23505":
				log.WithFields(log.Fields{
					"error": err,
					"users": user,
				}).ErrorWithCtx(ctx, "[UserRepository.Insert] User already exists")
				return ErrDuplicate
			}
		}
		log.WithFields(log.Fields{
			"error": err,
			"users": user,
		}).ErrorWithCtx(ctx, "[UserRepository.Insert] Failed to insert user")
		return err
	}
	return nil
}

func (repo *UserRepository) GetById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserRepository.GetById")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	var (
		err  error
		user model.User
		args []any
	)

	selectQuery := "users.id, users.email, users.password, users.password_salt, users.role, users.requested_role, users.is_active, users.verification_status, users.created_at, users.updated_at"
	whereQuery := " AND users.id = $1 AND users.is_deleted = FALSE"
	args = append(args, id)

	query := fmt.Sprintf(selectUser, selectQuery, whereQuery)

	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &user, query, args...)
	} else {
		err = repo.db.GetMaster().GetContext(ctx, &user, query, args...)
	}

	if err != nil {
		if errors.Is(err, sql2.ErrNoRows) {
			return nil, ErrNoResult
		}
		log.WithFields(log.Fields{
			"error": err,
			"id":    id,
		}).ErrorWithCtx(ctx, "[UserRepository.GetById] Failed to get user by id")
	}
	return &user, nil
}

func (repo *UserRepository) ChangePassword(ctx context.Context, id uuid.UUID, newPass string) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserRepository.ChangePassword")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	var (
		err  error
		args []any
	)

	if newPass == constants.EmptyString {
		return ErrNilParam
	}

	setQuery := "password = $1, updated_at = $2"
	whereQuery := " AND id = $3 AND is_deleted = FALSE"

	args = append(args, newPass, time.Now(), id)
	query := fmt.Sprintf(updateUser, setQuery, whereQuery)

	if sqlTrx != nil {
		_, err = sqlTrx.ExecContext(ctx, query, args...)
	} else {
		_, err = repo.db.GetMaster().ExecContext(ctx, query, args...)
	}

	if err != nil {

		log.WithFields(log.Fields{
			"error": err,
			"id":    id,
		}).ErrorWithCtx(ctx, "[UserRepository.ChangePassword] Failed to change password")
		return err
	}

	return nil
}

func (repo *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserRepository.GetByEmail")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	var (
		err  error
		user model.User
		args []any
	)

	decodedEmail, err := url.QueryUnescape(email)
	if err != nil {
		decodedEmail = email
	}

	selectQuery := "users.id, users.email, users.password, users.password_salt, users.role, users.requested_role, users.is_active, users.verification_status, users.created_at, users.updated_at"
	whereQuery := " AND users.email = $1 AND users.is_deleted = FALSE"
	args = append(args, decodedEmail)

	query := fmt.Sprintf(selectUser, selectQuery, whereQuery)

	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &user, query, args...)
	} else {
		err = repo.db.GetMaster().GetContext(ctx, &user, query, args...)
	}

	if err != nil {
		if errors.Is(err, sql2.ErrNoRows) {
			return nil, ErrNoResult
		}
		log.WithFields(log.Fields{
			"error": err,
			"email": email,
		}).ErrorWithCtx(ctx, "[UserRepository.GetByEmail] Failed to get user by email")
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) UpdateVerificationStatus(ctx context.Context, id uuid.UUID, status string, role string) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserRepository.UpdateVerificationStatus")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	var (
		err  error
		args []any
	)

	setQuery := "verification_status = $1, role = $2, is_active = $3, updated_at = $4"
	whereQuery := " AND id = $5 AND is_deleted = FALSE"

	args = append(args, status, role, true, time.Now(), id)
	query := fmt.Sprintf(updateUser, setQuery, whereQuery)
	if sqlTrx != nil {
		_, err = sqlTrx.ExecContext(ctx, query, args...)
	} else {
		_, err = repo.db.GetMaster().ExecContext(ctx, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"id":     id,
			"status": status,
		}).ErrorWithCtx(ctx, "[UserRepository.UpdateVerificationStatus] Failed to update verification status")
		return err
	}

	return nil
}

func (repo *UserRepository) GetPendingVerificationUsers(ctx context.Context) ([]model.User, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserRepository.GetPendingVerificationUsers")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	var (
		users []model.User
		args  []any
	)

	selectQuery := "users.id, users.email, users.password, users.password_salt, users.role, users.requested_role, users.is_active, users.verification_status, users.created_at, users.updated_at"
	whereQuery := " AND users.verification_status = $1 AND users.is_deleted = FALSE"
	args = append(args, model.VerificationStatusPending)

	query := fmt.Sprintf(selectUser, selectQuery, whereQuery)

	if sqlTrx != nil {
		err := sqlTrx.SelectContext(ctx, &users, query, args...)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).ErrorWithCtx(ctx, "[UserRepository.GetPendingVerificationUsers] Failed to get pending verification users")
			return nil, err
		}
	} else {
		err := repo.db.GetMaster().SelectContext(ctx, &users, query, args...)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).ErrorWithCtx(ctx, "[UserRepository.GetPendingVerificationUsers] Failed to get pending verification users")
			return nil, err
		}
	}

	return users, nil
}

func (repo *UserRepository) GetPendingVerificationUsersByRequestedRole(ctx context.Context, requestedRole string) ([]model.User, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserRepository.GetPendingVerificationUsersByRequestedRole")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	var (
		users []model.User
		args  []any
	)

	selectQuery := "users.id, users.email, users.password, users.password_salt, users.role, users.requested_role, users.is_active, users.verification_status, users.created_at, users.updated_at"
	whereQuery := " AND users.verification_status = $1 AND users.requested_role = $2 AND users.is_deleted = FALSE"
	args = append(args, model.VerificationStatusPending, requestedRole)

	query := fmt.Sprintf(selectUser, selectQuery, whereQuery)

	if sqlTrx != nil {
		err := sqlTrx.SelectContext(ctx, &users, query, args...)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
				"role":  requestedRole,
			}).ErrorWithCtx(ctx, "[UserRepository.GetPendingVerificationUsersByRequestedRole] Failed to get pending verification users by role")
			return nil, err
		}
	} else {
		err := repo.db.GetMaster().SelectContext(ctx, &users, query, args...)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
				"role":  requestedRole,
			}).ErrorWithCtx(ctx, "[UserRepository.GetPendingVerificationUsersByRequestedRole] Failed to get pending verification users by role")
			return nil, err
		}
	}

	return users, nil
}
