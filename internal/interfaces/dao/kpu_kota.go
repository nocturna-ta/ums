package dao

import (
	"context"
	sql2 "database/sql"
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
	insertKPUKota = `INSERT INTO kpu_kota (id, user_id, username, name, address, region, is_active, photo_path, telephone, registered_at, created_at, updated_at)
    						VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`
	selectKPUKota = `SELECT %s FROM kpu_kota %s WHERE TRUE %s`
	updateKPUKota = `UPDATE kpu_kota SET %s WHERE TRUE %s`
)

func (K *KPUKotaRepository) InsertKPUKota(ctx context.Context, kpu *model.KPUKota) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUKotaRepository.InsertKPUKota")
	defer span.End()

	var err error

	sqlTrx := utils.GetSqlTx(ctx)

	if sqlTrx != nil {
		_, err = sqlTrx.ExecContext(ctx, insertKPUKota, kpu.ID, kpu.UserID, kpu.Username, kpu.Name,
			kpu.Address, kpu.Region, kpu.IsActive, kpu.PhotoPath, kpu.Telephone, kpu.RegisteredAt, kpu.CreatedAt, kpu.UpdatedAt)
	} else {
		_, err = K.db.GetMaster().ExecContext(ctx, insertKPUKota, kpu.ID, kpu.UserID, kpu.Username, kpu.Name,
			kpu.Address, kpu.Region, kpu.IsActive, kpu.PhotoPath, kpu.Telephone, kpu.RegisteredAt, kpu.CreatedAt, kpu.UpdatedAt)
	}

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

	return nil
}

func (K *KPUKotaRepository) SendTxKPUKotaBlockchain(ctx context.Context, signedTransaction string) (string, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUKotaRepository.SendTxKPUKotaBlockchain")
	defer span.End()

	tx, err := utils2.StringToTx(signedTransaction)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUKotaRepository.InsertKPUKotaBlockchain] Failed to convert signed transaction to transaction")
		return "", err
	}

	txHash, err := K.client.SendTransaction(ctx, tx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"tx":    tx,
		}).ErrorWithCtx(ctx, "[KPUKotaRepository.InsertKPUKotaBlockchain] Failed to send transaction")
		return "", err
	}

	return txHash, nil
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

	selectQuery := `kpu_kota.id, kpu_kota.user_id, kpu_kota.username, kpu_kota.name, kpu_kota.address, kpu_kota.region, kpu_kota.is_active,
kpu_kota.photo_path,kpu_kota.telephone, kpu_kota.registered_at, kpu_kota.created_at, kpu_kota.updated_at`
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

	selectQuery := `kpu_kota.id, kpu_kota.user_id, kpu_kota.username, kpu_kota.name, kpu_kota.address, kpu_kota.region, kpu_kota.is_active,
kpu_kota.photo_path,kpu_kota.telephone, kpu_kota.registered_at, kpu_kota.created_at, kpu_kota.updated_at`
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

func (K *KPUKotaRepository) UpdateKPUKotaPhoto(ctx context.Context, id uuid.UUID, photoPath string) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUKotaRepository.UpdateKPUKotaPhoto")
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

	query := fmt.Sprintf(updateKPUKota, setQuery, whereQuery)

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
		}).ErrorWithCtx(ctx, "[KPUKotaRepository.UpdateKPUKotaPhoto] Failed to update photo path")
		return err
	}

	return nil
}

func (K *KPUKotaRepository) GetKPUKotaByID(ctx context.Context, id uuid.UUID) (*model.KPUKota, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUKotaRepository.GetKPUKotaByID")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)
	var (
		kpuKotaModel model.KPUKota
		err          error
		args         []any
	)

	selectQuery := `kpu_kota.id, kpu_kota.user_id, kpu_kota.username, kpu_kota.name, kpu_kota.address, kpu_kota.region, kpu_kota.is_active,
kpu_kota.photo_path,kpu_kota.telephone, kpu_kota.registered_at, kpu_kota.created_at, kpu_kota.updated_at`
	whereQuery := " AND kpu_kota.is_deleted = false AND kpu_kota.id = $1"
	joinQuery := ""
	args = append(args, id)

	query := fmt.Sprintf(selectKPUKota, selectQuery, joinQuery, whereQuery)
	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &kpuKotaModel, query, args...)
	} else {
		err = K.db.GetMaster().GetContext(ctx, &kpuKotaModel, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    id,
		}).ErrorWithCtx(ctx, "[KPUKotaRepository.GetKPUKotaByID] Failed to get kpu kota by ID")
		return nil, err
	}

	return &kpuKotaModel, nil
}

func (K *KPUKotaRepository) UpdateKPUKota(ctx context.Context, kpu *model.KPUKota) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUKotaRepository.UpdateKPUKota")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	var (
		err    error
		result sql2.Result
		args   []any
	)

	setQuery := "name = $1, region = $2, telephone = $3, username = $4, updated_at = $5"
	whereQuery := " AND id = $6 AND is_deleted = false"

	args = append(args, kpu.Name, kpu.Region, kpu.Telephone, kpu.Username, time.Now(), kpu.ID)
	query := fmt.Sprintf(updateKPUKota, setQuery, whereQuery)

	if sqlTrx != nil {
		result, err = sqlTrx.ExecContext(ctx, query, args...)
	} else {
		result, err = K.db.GetMaster().ExecContext(ctx, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error":    err,
			"kpu-kota": kpu,
		}).ErrorWithCtx(ctx, "[KPUKotaRepository.UpdateKPUKota] Failed to update kpu kota")
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err,
			"kpu-kota": kpu,
		}).ErrorWithCtx(ctx, "[KPUKotaRepository.UpdateKPUKota] Failed to get rows affected")
		return err
	}

	if rowAffected == 0 {
		return ErrNoUpdateHappened
	}

	return nil
}
