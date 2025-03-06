package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lib/pq"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/golib/txmanager/utils"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"github.com/nocturna-ta/ums/pkg/binding"
	utils2 "github.com/nocturna-ta/ums/pkg/utils"
)

type KPUBranchRepository struct {
	client   *ethclient.Client
	contract *binding.Votechain
	db       *sql.Store
}

type OptsKPUBranchRepository struct {
	Client          *ethclient.Client
	ContractAddress common.Address
	DB              *sql.Store
}

func NewKPUBranchRepository(opts *OptsKPUBranchRepository) repository.KPUBranchRepository {
	contract, err := binding.NewVotechain(opts.ContractAddress, opts.Client)
	if err != nil {
		return nil
	}
	return &KPUBranchRepository{
		client:   opts.Client,
		contract: contract,
		db:       opts.DB,
	}
}

const (
	insertKPUBranch = `INSERT INTO kpu_branches (id, name, branch_address, region, is_active, password, password_salt, created_at, updated_at)
    						VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	selectKPUBranch = `SELECT %s FROM kpu_branches %s WHERE TRUE %s`
)

func (K *KPUBranchRepository) InsertKPUBranch(ctx context.Context, kpuBranch *model.KPUBranch, signedTransaction string) error {
	spam, ctx := tracing.StartSpanFromContext(ctx, "KPUBranchRepository.InsertKPUBranch")
	defer spam.End()

	var (
		err error
	)

	sqlTrx := utils.GetSqlTx(ctx)

	if sqlTrx != nil {
		_, err = sqlTrx.ExecContext(ctx, insertKPUBranch, kpuBranch.ID, kpuBranch.Name, kpuBranch.BranchAddress, kpuBranch.Region, kpuBranch.IsActive, kpuBranch.Password, kpuBranch.PasswordSalt, kpuBranch.CreatedAt, kpuBranch.UpdatedAt)
	} else {
		_, err = K.db.GetMaster().ExecContext(ctx, insertKPUBranch, kpuBranch.ID, kpuBranch.Name, kpuBranch.BranchAddress, kpuBranch.Region, kpuBranch.IsActive, kpuBranch.Password, kpuBranch.PasswordSalt, kpuBranch.CreatedAt, kpuBranch.UpdatedAt)
	}

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code {
			case "23505":
				log.WithFields(log.Fields{
					"error":     err,
					"kpuBranch": kpuBranch,
				}).ErrorWithCtx(ctx, "[VoterRepository.InsertVoter] Duplicate entry")
				return ErrDuplicate
			}
		}

		log.WithFields(log.Fields{
			"error":     err,
			"kpuBranch": kpuBranch,
		}).ErrorWithCtx(ctx, "[KPUBranchRepository.InsertKPUBranch] Failed to insert kpu branch")
		return err
	}

	tx, err := utils2.StringToTx(signedTransaction)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUBranchRepository.InsertKPUBranch] Failed to convert signed transaction to transaction")
		return err
	}

	err = K.client.SendTransaction(ctx, tx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUBranchRepository.InsertKPUBranch] Failed to send transaction")
		return err
	}

	return nil
}

func (K *KPUBranchRepository) GetAllKPUBranch(ctx context.Context) ([]model.KPUBranch, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUBranchRepository.GetAllKPUBranch")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)
	var (
		kpuBranchModels []model.KPUBranch
		err             error
	)

	kpuBranches, err := K.contract.GetAllKPUBranches(nil)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUBranchRepository.GetAllKPUBranch] Failed to get all kpu branches")
	}

	selectQuery := `kpu_branches.id, kpu_branches.name, kpu_branches.branch_address, kpu_branches.region, kpu_branches.is_active, kpu_branches.password, kpu_branches.password_salt, kpu_branches.created_at, kpu_branches.updated_at`
	whereQuery := " AND kpu_branches.is_deleted = false"
	joinQuery := ""

	query := fmt.Sprintf(selectKPUBranch, selectQuery, joinQuery, whereQuery)
	if sqlTrx != nil {
		err = sqlTrx.SelectContext(ctx, &kpuBranchModels, query)
	} else {
		err = K.db.GetMaster().SelectContext(ctx, &kpuBranchModels, query)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUBranchRepository.GetAllKPUBranch] Failed to get all kpu branches")
		return nil, err
	}

	var matchedKPUBranches []model.KPUBranch
	for _, kpuBranch := range kpuBranches {
		for _, kpuBranchModel := range kpuBranchModels {
			if kpuBranch.BranchAddress.Hex() == kpuBranchModel.BranchAddress {
				matchedKPUBranches = append(matchedKPUBranches, kpuBranchModel)
				break
			}
		}
	}

	if len(matchedKPUBranches) == 0 {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUBranchRepository.GetAllKPUBranch] Failed to get all kpu branches")
		return nil, err
	}

	return matchedKPUBranches, nil

}

func (K *KPUBranchRepository) GetKPUBranchByAddress(ctx context.Context, address string) (*model.KPUBranch, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUBranchRepository.GetKPUBranchByAddress")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)
	var (
		kpuBranchModel model.KPUBranch
		err            error
		args           []any
	)

	kpuBranches, err := K.contract.GetBranchByAddress(nil, common.HexToAddress(address))
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUBranchRepository.GetKPUBranchByAddress] Failed to get kpu branch by address")
		return nil, err
	}

	selectQuery := `kpu_branches.id, kpu_branches.name, kpu_branches.branch_address, kpu_branches.region, kpu_branches.is_active, kpu_branches.password, kpu_branches.password_salt, kpu_branches.created_at, kpu_branches.updated_at`
	whereQuery := " AND kpu_branches.is_deleted = false AND kpu_branches.branch_address = $1"
	joinQuery := ""
	args = append(args, address)

	query := fmt.Sprintf(selectKPUBranch, selectQuery, joinQuery, whereQuery)
	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &kpuBranchModel, query, args...)
	} else {
		err = K.db.GetMaster().GetContext(ctx, &kpuBranchModel, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUBranchRepository.GetKPUBranchByAddress] Failed to get kpu branch by address")
		return nil, err
	}

	if kpuBranches.BranchAddress.Hex() != kpuBranchModel.BranchAddress {
		log.WithFields(log.Fields{
			"error": "not matching kpu branches found",
		}).ErrorWithCtx(ctx, "[KPUBranchRepository.GetKPUBranchByAddress] Failed to get kpu branch by address")
		return nil, ErrNoResult
	}

	return &kpuBranchModel, nil
}
