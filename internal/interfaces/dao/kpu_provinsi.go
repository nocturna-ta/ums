package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
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
	kpuManager2 "github.com/nocturna-ta/votechain-contract/binding/kpuManager"
	"github.com/nocturna-ta/votechain-contract/interfaces"
	"time"
)

type KPUProvinsiRepository struct {
	client   ethereum.Client
	contract interfaces.KpuManagerInterface
	db       *sql.Store
}

type OptsKPUProvinsiRepository struct {
	Client          ethereum.Client
	ContractAddress common.Address
	Contract        interfaces.KpuManagerInterface
	DB              *sql.Store
}

func NewKPUProvinsiRepository(opts *OptsKPUProvinsiRepository) repository.KPUProvinsiRepository {
	var contractInterface interfaces.KpuManagerInterface
	contract, err := kpuManager2.NewKpuManager(opts.ContractAddress, opts.Client.GetEthClient())
	if err != nil {
		return nil
	}
	contractInterface = contract
	return &KPUProvinsiRepository{
		client:   opts.Client,
		contract: contractInterface,
		db:       opts.DB,
	}
}

const (
	insertKPUProvinsi = `INSERT INTO kpu_provinsi (id, user_id, name, address, region, is_active, photo_path, created_at, updated_at)
    						VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	selectKPUProvinsi = `SELECT %s FROM kpu_provinsi %s WHERE TRUE %s`
	updateKPUProvinsi = `UPDATE kpu_provinsi SET %s WHERE TRUE %s`
)

func (K *KPUProvinsiRepository) InsertKPUProvinsi(ctx context.Context, kpu *model.KPUProvinsi, signedTransaction string) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiRepository.InsertKPUProvinsi")
	defer span.End()

	tx, err := utils2.StringToTx(signedTransaction)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.InsertKPUProvinsi] Failed to convert signed transaction to transaction")
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
			}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.InsertKPUProvinsi] Failed to begin transaction")
			return err
		}

		ownTransaction = true

		defer func() {
			if err != nil && ownTransaction {
				rollbackErr := sqlTrx.Rollback()
				if rollbackErr != nil {
					log.WithFields(log.Fields{
						"error": rollbackErr,
					}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.InsertKPUProvinsi] Failed to rollback transaction")
				}
			}
		}()
	}

	_, err = sqlTrx.ExecContext(ctx, insertKPUProvinsi, kpu.ID, kpu.UserID, kpu.Name, kpu.Address, kpu.Region, kpu.IsActive, kpu.PhotoPath, kpu.CreatedAt, kpu.UpdatedAt)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code {
			case "23505":
				log.WithFields(log.Fields{
					"error": err,
					"kpu":   kpu,
				}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.InsertKPUProvinsi] Duplicate entry")
				return ErrDuplicate
			}
		}

		log.WithFields(log.Fields{
			"error": err,
			"kpu":   kpu,
		}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.InsertKPUProvinsi] Failed to insert kpu kota")
		return err
	}

	err = K.client.SendTransaction(ctx, tx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.InsertKPUProvinsi] Failed to send transaction")

		if ownTransaction {
			rollbackErr := sqlTrx.Rollback()
			if rollbackErr != nil {
				log.WithFields(log.Fields{
					"error": rollbackErr,
				}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.InsertKPUProvinsi] Failed to rollback transaction")
			}
		}
		return err
	}

	if ownTransaction {
		if err := sqlTrx.Commit(); err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.InsertKPUProvinsi] Failed to commit transaction")
			return err
		}
	}

	return nil
}

func (K *KPUProvinsiRepository) GetAllKPUProvinsi(ctx context.Context) ([]model.KPUProvinsi, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiRepository.GetAllKPUProvinsi")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)
	var (
		kpuProvinsiModels []model.KPUProvinsi
		err               error
	)

	kpuProvinsi, err := K.contract.GetAllKPUProvinsi(nil)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.GetAllKPUProvinsi] Failed to get all kpu provinsi")
	}

	selectQuery := `kpu_provinsi.id, kpu_provinsi.name, kpu_provinsi.address, kpu_provinsi.region, kpu_provinsi.is_active, kpu_provinsi.photo_path, kpu_provinsi.created_at, kpu_provinsi.updated_at`
	whereQuery := " AND kpu_provinsi.is_deleted = false"
	joinQuery := ""

	query := fmt.Sprintf(selectKPUProvinsi, selectQuery, joinQuery, whereQuery)
	if sqlTrx != nil {
		err = sqlTrx.SelectContext(ctx, &kpuProvinsiModels, query)
	} else {
		err = K.db.GetMaster().SelectContext(ctx, &kpuProvinsiModels, query)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.GetAllKPUProvinsi] Failed to get all kpu provinsi")
		return nil, err
	}

	var matchedKPUProvinsi []model.KPUProvinsi
	for _, kpuProvins := range kpuProvinsi {
		for _, kpuProvinsiModel := range kpuProvinsiModels {
			if kpuProvins.Address.Hex() == kpuProvinsiModel.Address {
				matchedKPUProvinsi = append(matchedKPUProvinsi, kpuProvinsiModel)
				break
			}
		}
	}

	if len(matchedKPUProvinsi) == 0 {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.GetAllKPUProvinsi] Failed to get all kpu provinsi")
		return nil, err
	}

	return matchedKPUProvinsi, nil
}

func (K *KPUProvinsiRepository) GetKPUProvinsiByAddress(ctx context.Context, address string) (*model.KPUProvinsi, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiRepository.GetKPUProvinsiByAddress")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)
	var (
		kpuProvinsiModel model.KPUProvinsi
		err              error
		args             []any
	)

	kpuProvinsi, err := K.contract.GetKpuProvinsiByAddress(nil, common.HexToAddress(address))
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.GetKPUProvinsiByAddress] Failed to get kpu provinsi by address")
		return nil, err
	}

	selectQuery := `kpu_provinsi.id, kpu_provinsi.name, kpu_provinsi.address, kpu_provinsi.region, kpu_provinsi.is_active, kpu_provinsi.photo_path, kpu_provinsi.created_at, kpu_provinsi.updated_at`
	whereQuery := " AND kpu_provinsi.is_deleted = false AND kpu_provinsi.address = $1"
	joinQuery := ""
	args = append(args, address)

	query := fmt.Sprintf(selectKPUProvinsi, selectQuery, joinQuery, whereQuery)
	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &kpuProvinsiModel, query, args...)
	} else {
		err = K.db.GetMaster().GetContext(ctx, &kpuProvinsiModel, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.GetKPUProvinsiByAddress] Failed to get kpu provinsi by address")
		return nil, err
	}

	if kpuProvinsi.Address.Hex() != kpuProvinsiModel.Address {
		log.WithFields(log.Fields{
			"error": "not matching kpu provinsi] found",
		}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.GetKPUProvinsiByAddress] Failed to get kpu provinsi by address")
		return nil, ErrNoResult
	}

	return &kpuProvinsiModel, nil
}

func (K *KPUProvinsiRepository) GetKPUProvinsiByID(ctx context.Context, id uuid.UUID) (*model.KPUProvinsi, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiRepository.GetKPUProvinsiByID")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)
	var (
		kpuProvinsiModel model.KPUProvinsi
		err              error
		args             []any
	)

	selectQuery := `kpu_provinsi.id, kpu_provinsi.name, kpu_provinsi.address, kpu_provinsi.region, kpu_provinsi.is_active, kpu_provinsi.photo_path, kpu_provinsi.created_at, kpu_provinsi.updated_at`
	whereQuery := " AND kpu_provinsi.is_deleted = false AND kpu_provinsi.id = $1"
	joinQuery := ""
	args = append(args, id)

	query := fmt.Sprintf(selectKPUProvinsi, selectQuery, joinQuery, whereQuery)
	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &kpuProvinsiModel, query, args...)
	} else {
		err = K.db.GetMaster().GetContext(ctx, &kpuProvinsiModel, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    id,
		}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.GetKPUProvinsiByID] Failed to get kpu provinsi by ID")
		return nil, err
	}

	return &kpuProvinsiModel, nil
}

func (K *KPUProvinsiRepository) UpdateKPUProvinsiPhoto(ctx context.Context, id uuid.UUID, photoPath string) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiRepository.UpdateKPUProvinsiPhoto")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)
	var (
		err  error
		args []any
	)

	now := time.Now()

	setQuery := "photo_path = $1, updated_at = $2"
	whereQuery := " AND id = $3 AND is_deleted = false"
	args = append(args, photoPath, now, id)

	query := fmt.Sprintf(updateKPUProvinsi, setQuery, whereQuery)

	if sqlTrx != nil {
		_, err = sqlTrx.ExecContext(ctx, query, photoPath, now, id)
	} else {
		_, err = K.db.GetMaster().ExecContext(ctx, query, photoPath, now, id)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error":     err,
			"id":        id,
			"photoPath": photoPath,
		}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.UpdateKPUProvinsiPhoto] Failed to update photo path")
		return err
	}

	return nil
}
