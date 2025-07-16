package dao

import (
	"context"
	sql2 "database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/golib/ethereum"
	"github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/golib/txmanager/utils"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	utils2 "github.com/nocturna-ta/ums/pkg/utils"
)

type VoterRepository struct {
	client ethereum.Client
	db     *sql.Store
}

type OptsVoterRepository struct {
	Client ethereum.Client
	DB     *sql.Store
}

func NewVoterRepository(opts *OptsVoterRepository) repository.VoterRepository {
	return &VoterRepository{
		client: opts.Client,
		db:     opts.DB,
	}
}

const (
	insertVoter = `INSERT INTO voters (id, user_id, nik, full_name, gender, birth_place, 
                    birth_date, residential_address, region, voter_address, is_registered, 
                    ktp_photo_path,telephone,has_voted, voted_at, last_login, created_at, updated_at)
					VALUES($1, $2, $3, $4, $5, $6, $7,$8,$9, $10, $11, $12, $13, $14, $15, $16, $17,$18)`
	selectVoter = `SELECT %s FROM voters %s WHERE TRUE %s`
	updateVoter = `UPDATE voters SET %s WHERE TRUE %s`
)

func (v *VoterRepository) InsertVoter(ctx context.Context, voter *model.Voter) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "VoterRepository.InsertVoter")
	defer span.End()

	var err error

	sqlTrx := utils.GetSqlTx(ctx)

	if sqlTrx != nil {
		_, err = sqlTrx.ExecContext(ctx, insertVoter,
			voter.ID,
			voter.UserID,
			voter.NIK,
			voter.FullName,
			voter.Gender,
			voter.BirthPlace,
			voter.BirthDate,
			voter.ResidentialAddress,
			voter.Region,
			voter.VoterAddress,
			voter.IsRegistered,
			voter.KTPPhotoPath,
			voter.Telephone,
			voter.HasVoted,
			voter.VotedAt,
			voter.LastLogin,
			voter.CreatedAt,
			voter.UpdatedAt,
		)
	} else {
		_, err = v.db.GetMaster().ExecContext(ctx, insertVoter,
			voter.ID,
			voter.UserID,
			voter.NIK,
			voter.FullName,
			voter.Gender,
			voter.BirthPlace,
			voter.BirthDate,
			voter.ResidentialAddress,
			voter.Region,
			voter.VoterAddress,
			voter.IsRegistered,
			voter.KTPPhotoPath,
			voter.Telephone,
			voter.HasVoted,
			voter.VotedAt,
			voter.LastLogin,
			voter.CreatedAt,
			voter.UpdatedAt,
		)

	}

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code {
			case "23505":
				log.WithFields(log.Fields{
					"error": err,
					"voter": voter,
				}).ErrorWithCtx(ctx, "[VoterRepository.InsertVoter] Duplicate entry")
				return ErrDuplicate
			}
		}

		log.WithFields(log.Fields{
			"error": err,
			"voter": voter,
		}).ErrorWithCtx(ctx, "[VoterRepository.InsertVoter] Failed to insert entry")
		return err
	}

	return nil
}

func (v *VoterRepository) SendTxVoterBlockchain(ctx context.Context, signedTransaction string) (string, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "VoterRepository.SendTxVoterBlockchain")
	defer span.End()

	tx, err := utils2.StringToTx(signedTransaction)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[VoterRepository.SendTxVoterBlockchain] Failed to convert string to transaction")
		return "", err
	}

	txHash, err := v.client.SendTransaction(ctx, tx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"tx":    tx,
		}).ErrorWithCtx(ctx, "[VoterRepository.SendTxVoterBlockchain] Failed to send transaction")
		return "", err
	}

	return txHash, nil
}

func (v *VoterRepository) GetAllVoter(ctx context.Context) ([]model.Voter, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "VoterRepository.GetAllVoter")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)
	var (
		voters []model.Voter
		err    error
	)

	selectQuery := `id, user_id, nik, full_name, gender, birth_place, 
			birth_date, residential_address, region, voter_address, is_registered, ktp_photo_path, telephone,
			has_voted, voted_at, last_login, created_at, updated_at`
	whereQuery := " AND voters.is_deleted = false"
	joinQuery := ""

	query := fmt.Sprintf(selectVoter, selectQuery, joinQuery, whereQuery)
	if sqlTrx != nil {
		err = sqlTrx.SelectContext(ctx, &voters, query)
	} else {
		err = v.db.GetMaster().SelectContext(ctx, &voters, query)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[VoterRepository.GetAllVoter] Failed to get all voters")
		return nil, err
	}

	return voters, nil
}

func (v *VoterRepository) GetVoterByNIK(ctx context.Context, nik string) (*model.Voter, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "VoterRepository.GetVoterByNIK")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)
	var (
		voter model.Voter
		args  []any
		err   error
	)

	selectQuery := `id, user_id, nik, full_name, gender, birth_place, 
			birth_date, residential_address, region, voter_address, is_registered, ktp_photo_path, telephone,
			has_voted, voted_at, last_login, created_at, updated_at`

	whereQuery := " AND voters.is_deleted = false AND voters.nik = $1"
	joinQuery := ""
	args = append(args, nik)

	query := fmt.Sprintf(selectVoter, selectQuery, joinQuery, whereQuery)
	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &voter, query, args...)
	} else {
		err = v.db.GetMaster().GetContext(ctx, &voter, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[VoterRepository.GetVoterByNIK] Failed to get voter by nik")
		return nil, err
	}

	return &voter, nil
}

func (v *VoterRepository) GetVoterByAddress(ctx context.Context, address string) (*model.Voter, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "VoterRepository.GetVoterByAddress")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)
	var (
		voter model.Voter
		args  []any
		err   error
	)

	selectQuery := `id, user_id, nik, full_name, gender, birth_place, 
			birth_date, residential_address, region, voter_address, is_registered, ktp_photo_path, telephone,
			has_voted, voted_at, last_login, created_at, updated_at`

	whereQuery := " AND voters.is_deleted = false AND voters.voter_address = $1"
	joinQuery := ""
	args = append(args, address)

	query := fmt.Sprintf(selectVoter, selectQuery, joinQuery, whereQuery)
	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &voter, query, args...)
	} else {
		err = v.db.GetMaster().GetContext(ctx, &voter, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[VoterRepository.GetVoterByAddress] Failed to get voter by address")
		return nil, err
	}

	return &voter, nil
}

func (v *VoterRepository) GetVoterByRegion(ctx context.Context, region string) ([]model.Voter, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "VoterRepository.GetVoterByRegion")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	var (
		voters []model.Voter
		args   []any
		err    error
	)

	selectQuery := `id, user_id, nik, full_name, gender, birth_place, 
			birth_date, residential_address, region, voter_address, is_registered, ktp_photo_path, telephone,
			has_voted, voted_at, last_login, created_at, updated_at`

	whereQuery := " AND is_deleted = false AND region = $1"
	joinQuery := ""
	args = append(args, region)

	query := fmt.Sprintf(selectVoter, selectQuery, joinQuery, whereQuery)
	if sqlTrx != nil {
		err = sqlTrx.SelectContext(ctx, &voters, query, args...)
	} else {
		err = v.db.GetMaster().SelectContext(ctx, &voters, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[VoterRepository.GetVoterByRegion] Failed to get voter by region")
		return nil, err
	}

	return voters, nil
}

func (v *VoterRepository) GetVoterByID(ctx context.Context, id uuid.UUID) (*model.Voter, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "VoterRepository.GetVoterByID")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)
	var (
		voter model.Voter
		err   error
		args  []any
	)

	selectQuery := `id, user_id, nik, full_name, gender, birth_place, 
			birth_date, residential_address, region, voter_address, is_registered, ktp_photo_path, telephone,
			has_voted, voted_at, last_login, created_at, updated_at`

	whereQuery := " AND voters.is_deleted = false AND voters.id = $1"
	joinQuery := ""
	args = append(args, id)

	query := fmt.Sprintf(selectVoter, selectQuery, joinQuery, whereQuery)
	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &voter, query, args...)
	} else {
		err = v.db.GetMaster().GetContext(ctx, &voter, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    id,
		}).ErrorWithCtx(ctx, "[VoterRepository.GetVoterByID] Failed to get voter by ID")
		return nil, err
	}

	return &voter, nil
}

func (v *VoterRepository) UpdateVoter(ctx context.Context, voter *model.Voter) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "VoterRepository.UpdateVoter")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	var (
		err    error
		args   []any
		result sql2.Result
	)

	setQuery := `has_voted = $1, voted_at = $2, updated_at = $3`
	whereQuery := " AND voters.is_deleted = false AND voters.id = $4"
	args = append(args, voter.HasVoted, voter.VotedAt, voter.UpdatedAt, voter.ID)
	query := fmt.Sprintf(updateVoter, setQuery, whereQuery)

	if sqlTrx != nil {
		result, err = sqlTrx.ExecContext(ctx, query, args...)
	} else {
		result, err = v.db.GetMaster().ExecContext(ctx, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    voter.ID,
		}).ErrorWithCtx(ctx, "[VoterRepository.UpdateVoter] Failed to update voter")
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"voter": voter,
		}).ErrorWithCtx(ctx, "[VoterRepository.UpdateVoter] Failed to get rows affected")
		return err
	}

	if rowsAffected == 0 {
		return ErrNoUpdateHappened
	}

	return nil
}

func (v *VoterRepository) GetVoterByUserID(ctx context.Context, userID uuid.UUID) (*model.Voter, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "VoterRepository.GetVoterByUserID")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)
	var (
		voter model.Voter
		err   error
		args  []any
	)

	selectQuery := `id, user_id, nik, full_name, gender, birth_place, 
			birth_date, residential_address, region, voter_address, is_registered, ktp_photo_path, telephone,
			has_voted, voted_at, last_login, created_at, updated_at`

	whereQuery := " AND voters.is_deleted = false AND voters.user_id = $1"
	joinQuery := ""
	args = append(args, userID)

	query := fmt.Sprintf(selectVoter, selectQuery, joinQuery, whereQuery)
	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &voter, query, args...)
	} else {
		err = v.db.GetMaster().GetContext(ctx, &voter, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    userID,
		}).ErrorWithCtx(ctx, "[VoterRepository.GetVoterByID] Failed to get voter by ID")
		return nil, err
	}

	return &voter, nil
}
