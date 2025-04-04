package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lib/pq"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/golib/ethereum"
	"github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/golib/txmanager/utils"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	utils2 "github.com/nocturna-ta/ums/pkg/utils"
	"github.com/nocturna-ta/votechain-contract/binding"
)

type VoterRepository struct {
	client   ethereum.Client
	contract *binding.Votechain
	db       *sql.Store
}

type OptsVoterRepository struct {
	Client          ethereum.Client
	DB              *sql.Store
	ContractAddress common.Address
}

func NewVoterRepository(opts *OptsVoterRepository) repository.VoterRepository {
	contract, err := binding.NewVotechain(opts.ContractAddress, opts.Client.GetEthClient())
	if err != nil {
		return nil
	}
	return &VoterRepository{
		client:   opts.Client,
		contract: contract,
		db:       opts.DB,
	}
}

const (
	insertVoter = `INSERT INTO voters (id, nik, voter_address, is_registered, has_voted, voted_at, region, last_login, created_at, updated_at)
								VALUES($1, $2, $3, $4, $5, $6, $7,$8,$9, $10)`
	selectVoter = `SELECT %s FROM voters %s WHERE TRUE %s`
	updateVoter = `UPDATE voters SET %s WHERE TRUE %s`
)

func (v *VoterRepository) InsertVoter(ctx context.Context, voter *model.Voter, signedTransaction string) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "VoterRepository.InsertVoter")
	defer span.End()

	tx, err := utils2.StringToTx(signedTransaction)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[VoterRepository.InsertVoter] Failed to convert signed transaction to transaction")
		return err
	}

	sqlTrx := utils.GetSqlTx(ctx)

	var ownTransaction bool
	if sqlTrx == nil {
		sqlTrx, err = v.db.GetMaster().BeginTxx(ctx, nil)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).ErrorWithCtx(ctx, "[VoterRepository.InsertVoter] Failed to begin transaction")
			return err
		}
		ownTransaction = true

		defer func() {
			if err != nil && ownTransaction {
				rollbackErr := sqlTrx.Rollback()
				if rollbackErr != nil {
					log.WithFields(log.Fields{
						"error": rollbackErr,
					}).ErrorWithCtx(ctx, "[VoterRepository.InsertVoter] Failed to rollback transaction")
				}
			}
		}()
	}

	_, err = sqlTrx.ExecContext(ctx, insertVoter,
		voter.ID,
		voter.NIK,
		voter.VoterAddress,
		voter.IsRegistered,
		voter.HasVoted,
		voter.VotedAt,
		voter.Region,
		voter.LastLogin,
		voter.CreatedAt,
		voter.UpdatedAt,
	)

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
	err = v.client.SendTransaction(ctx, tx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[VoterRepository.InsertVoter] Failed to send transaction")

		if ownTransaction {
			rollbackErr := sqlTrx.Rollback()
			if rollbackErr != nil {
				log.WithFields(log.Fields{
					"error": rollbackErr,
				}).ErrorWithCtx(ctx, "[VoterRepository.InsertVoter] Failed to rollback transaction")
			}
		}
		return err
	}

	if ownTransaction {
		if err := sqlTrx.Commit(); err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).ErrorWithCtx(ctx, "[VoterRepository.InsertVoter] Failed to commit transaction")
			return err
		}
	}

	return nil
}

func (v *VoterRepository) GetAllVoter(ctx context.Context) ([]model.Voter, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "VoterRepository.GetAllVoter")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)
	var (
		votersModel []model.Voter
	)

	voters, err := v.contract.GetAllVoter(nil)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[VoterRepository.GetAllVoter] Failed to get all voters")
	}

	selectQuery := `
			   voters.id,
               voters.nik,
		       voters.voter_address,
			   voters.region,
		       voters.is_registered,
		       voters.has_voted,
		       voters.voted_at,
		       voters.last_login,
		       voters.created_at,
		       voters.updated_at`

	whereQuery := " AND voters.is_deleted = false"
	joinQuery := ""

	query := fmt.Sprintf(selectVoter, selectQuery, joinQuery, whereQuery)
	if sqlTrx != nil {
		err = sqlTrx.SelectContext(ctx, &votersModel, query)
	} else {
		err = v.db.GetMaster().SelectContext(ctx, &votersModel, query)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[VoterRepository.GetAllVoter] Failed to get all voters")
		return nil, err
	}

	var matchedVoters []model.Voter
	for _, voterContract := range voters {
		for _, voterDb := range votersModel {
			if voterContract.VoterAddress.Hex() == voterDb.VoterAddress {
				matchedVoters = append(matchedVoters, voterDb)
				break
			}
		}
	}

	if len(matchedVoters) == 0 {
		log.WithFields(log.Fields{
			"error": "not matching voters found",
		}).ErrorWithCtx(ctx, "[VoterRepository.GetAllVoter] Failed to get all voters")
	}

	return matchedVoters, nil
}

func (v *VoterRepository) GetVoterByNIK(ctx context.Context, nik string) (*model.Voter, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "VoterRepository.GetVoterByNIK")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)
	var (
		voter model.Voter
		args  []any
	)

	voterContract, err := v.contract.GetVoterByNIK(nil, nik)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[VoterRepository.GetVoterByNIK] Failed to get voter by nik")
		return nil, err
	}

	selectQuery := `
			   voters.id,
               voters.nik,
		       voters.voter_address,
			   voters.region,
		       voters.is_registered,
		       voters.has_voted,
		       voters.voted_at,
		       voters.last_login,
		       voters.created_at,
		       voters.updated_at`

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

	if voterContract.VoterAddress.Hex() != voter.VoterAddress {
		log.WithFields(log.Fields{
			"error": "not matching voters found",
		}).ErrorWithCtx(ctx, "[VoterRepository.GetVoterByNIK] Failed to get voter by nik")
		return nil, ErrNoResult
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
	)

	voterContract, err := v.contract.GetVoterByAddress(nil, common.HexToAddress(address))
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[VoterRepository.GetVoterByAddress] Failed to get voter by address")
		return nil, err
	}

	selectQuery := `
			   voters.id,
               voters.nik,
		       voters.voter_address,
			   voters.region,
		       voters.is_registered,
		       voters.has_voted,
		       voters.voted_at,
		       voters.last_login,
		       voters.created_at,
		       voters.updated_at`

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

	if voterContract.VoterAddress.Hex() != voter.VoterAddress {
		log.WithFields(log.Fields{
			"error": "not matching voters found",
		}).ErrorWithCtx(ctx, "[VoterRepository.GetVoterByAddress] Failed to get voter by address")
		return nil, ErrNoResult
	}

	return &voter, nil
}

func (v *VoterRepository) GetVoterByRegion(ctx context.Context, region string) ([]model.Voter, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "VoterRepository.GetVoterByRegion")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)
	var (
		votersModel []model.Voter
		args        []any
	)

	voters, err := v.contract.GetVoterByRegion(nil, region)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[VoterRepository.GetVoterByRegion] Failed to get voter by region")
	}

	selectQuery := `
			   voters.id,
               voters.nik,
		       voters.voter_address,
			   voters.region,
		       voters.is_registered,
		       voters.has_voted,
		       voters.voted_at,
		       voters.last_login,
		       voters.created_at,
		       voters.updated_at`

	whereQuery := " AND voters.is_deleted = false AND voters.region = $1"
	joinQuery := ""
	args = append(args, region)

	query := fmt.Sprintf(selectVoter, selectQuery, joinQuery, whereQuery)
	if sqlTrx != nil {
		err = sqlTrx.SelectContext(ctx, &votersModel, query, args...)
	} else {
		err = v.db.GetMaster().SelectContext(ctx, &votersModel, query, args...)
	}

	fmt.Println(votersModel)
	fmt.Println(voters)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[VoterRepository.GetVoterByRegion] Failed to get voter by region")
		return nil, err
	}

	var matchedVoters []model.Voter
	for _, voterContract := range voters {
		for _, voterDb := range votersModel {
			if voterContract.VoterAddress.Hex() == voterDb.VoterAddress {
				matchedVoters = append(matchedVoters, voterDb)
				break
			}
		}
	}

	if len(matchedVoters) == 0 {
		log.WithFields(log.Fields{
			"error": "not matching voters found",
		}).ErrorWithCtx(ctx, "[VoterRepository.GetVoterByRegion] Failed to get voter by region")
	}

	return matchedVoters, nil
}
