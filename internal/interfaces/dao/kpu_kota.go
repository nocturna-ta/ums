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
	kpuManager2 "github.com/nocturna-ta/votechain-contract/binding/kpuManager"
	"github.com/nocturna-ta/votechain-contract/interfaces"
)

type KPUKotaRepository struct {
	client   ethereum.Client
	contract interfaces.KpuManagerInterface
	db       *sql.Store
}

type OptsKPUKotaRepository struct {
	Client          ethereum.Client
	ContractAddress common.Address
	Contract        interfaces.KpuManagerInterface
	DB              *sql.Store
}

func NewKPUKotaRepository(opts *OptsKPUKotaRepository) repository.KPUKotaRepository {
	var contractInterface interfaces.KpuManagerInterface
	contract, err := kpuManager2.NewKpuManager(opts.ContractAddress, opts.Client.GetEthClient())
	if err != nil {
		return nil
	}
	contractInterface = contract
	return &KPUKotaRepository{
		client:   opts.Client,
		contract: contractInterface,
		db:       opts.DB,
	}
}

const (
	insertKPUKota = `INSERT INTO kpu_kota (id, name, address, region, is_active, created_at, updated_at)
    						VALUES($1,$2,$3,$4,$5,$6,$7)`
	selectKPUKota = `SELECT %s FROM kpu_kota %s WHERE TRUE %s`
)

func (K *KPUKotaRepository) InsertKPUKota(ctx context.Context, kpu *model.KPUKota, signedTransaction string) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUKotaRepository.InsertKPUKota")
	defer span.End()

	tx, err := utils2.StringToTx(signedTransaction)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUKotaRepository.InsertKPUKota] Failed to convert signed transaction to transaction")
		return err
	}
	sqlTrx := utils.GetSqlTx(ctx)

	var ownTransaction bool
	if sqlTrx == nil {
		var err error
		sqlTrx, err = K.db.GetMaster().BeginTxx(ctx, nil)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).ErrorWithCtx(ctx, "[KPUKotaRepository.InsertKPUKota] Failed to begin transaction")
			return err
		}

		ownTransaction = true

		defer func() {
			if err != nil && ownTransaction {
				rollbackErr := sqlTrx.Rollback()
				if rollbackErr != nil {
					log.WithFields(log.Fields{
						"error": rollbackErr,
					}).ErrorWithCtx(ctx, "[KPUKotaRepository.InsertKPUKota] Failed to rollback transaction")
				}
			}
		}()
	}

	_, err = sqlTrx.ExecContext(ctx, insertKPUKota, kpu.ID, kpu.Name, kpu.Address, kpu.Region, kpu.IsActive, kpu.CreatedAt, kpu.UpdatedAt)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code {
			case "23505":
				log.WithFields(log.Fields{
					"error": err,
					"kpu":   kpu,
				}).ErrorWithCtx(ctx, "[KPUKotaRepository.InsertKPUKota] Duplicate entry")
				return ErrDuplicate
			}
		}

		log.WithFields(log.Fields{
			"error": err,
			"kpu":   kpu,
		}).ErrorWithCtx(ctx, "[KPUKotaRepository.InsertKPUKota] Failed to insert kpu kota")
		return err
	}

	err = K.client.SendTransaction(ctx, tx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUKotaRepository.InsertKPUKota] Failed to send transaction")

		if ownTransaction {
			rollbackErr := sqlTrx.Rollback()
			if rollbackErr != nil {
				log.WithFields(log.Fields{
					"error": rollbackErr,
				}).ErrorWithCtx(ctx, "[KPUKotaRepository.InsertKPUKota] Failed to rollback transaction")
			}
		}
		return err
	}

	if ownTransaction {
		if err := sqlTrx.Commit(); err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).ErrorWithCtx(ctx, "[KPUKotaRepository.InsertKPUKota] Failed to commit transaction")
			return err
		}
	}

	return nil
}

func (K *KPUKotaRepository) GetAllKPUKota(ctx context.Context) ([]model.KPUKota, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUKotaRepository.GetAllKPUKota")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)
	var (
		kpuKotaModels []model.KPUKota
		err           error
	)

	kpuKota, err := K.contract.GetAllKPUKota(nil)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUKotaRepository.GetAllKPUKota] Failed to get all kpu kota")
	}

	selectQuery := `kpu_kota.id, kpu_kota.name, kpu_kota.address, kpu_kota.region, kpu_kota.is_active, kpu_kota.created_at, kpu_kota.updated_at`
	whereQuery := " AND kpu_kota.is_deleted = false"
	joinQuery := ""

	query := fmt.Sprintf(selectKPUKota, selectQuery, joinQuery, whereQuery)
	if sqlTrx != nil {
		err = sqlTrx.SelectContext(ctx, &kpuKotaModels, query)
	} else {
		err = K.db.GetMaster().SelectContext(ctx, &kpuKotaModels, query)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUKotaRepository.GetAllKPUKota] Failed to get all kpu kota")
		return nil, err
	}

	var matchedKPUKota []model.KPUKota
	for _, kpuKotas := range kpuKota {
		for _, kpuKotaModel := range kpuKotaModels {
			if kpuKotas.Address.Hex() == kpuKotaModel.Address {
				matchedKPUKota = append(matchedKPUKota, kpuKotaModel)
				break
			}
		}
	}

	if len(matchedKPUKota) == 0 {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUKotaRepository.GetAllKPUKota] Failed to get all kpu kota")
		return nil, err
	}

	return matchedKPUKota, nil
}

func (K *KPUKotaRepository) GetKPUKotaByAddress(ctx context.Context, address string) (*model.KPUKota, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUKotaRepository.GetKPUKotaByAddress")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)
	var (
		kpuKotaModel model.KPUKota
		err          error
		args         []any
	)

	kpuKota, err := K.contract.GetKpuKotaByAddress(nil, common.HexToAddress(address))
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUKotaRepository.GetKPUKotaByAddress] Failed to get kpu kota by address")
		return nil, err
	}

	selectQuery := `kpu_kota.id, kpu_kota.name, kpu_kota.address, kpu_kota.region, kpu_kota.is_active, kpu_kota.created_at, kpu_kota.updated_at`
	whereQuery := " AND kpu_kota.is_deleted = false AND kpu_kota.address = $1"
	joinQuery := ""
	args = append(args, address)

	query := fmt.Sprintf(selectKPUKota, selectQuery, joinQuery, whereQuery)
	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &kpuKotaModel, query, args...)
	} else {
		err = K.db.GetMaster().GetContext(ctx, &kpuKotaModel, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUKotaRepository.GetKPUKotaByAddress] Failed to get kpu kota by address")
		return nil, err
	}

	if kpuKota.Address.Hex() != kpuKotaModel.Address {
		log.WithFields(log.Fields{
			"error": "not matching kpu kota found",
		}).ErrorWithCtx(ctx, "[KPUKotaRepository.GetKPUKotaByAddress] Failed to get kpu kota by address")
		return nil, ErrNoResult
	}

	return &kpuKotaModel, nil
}
