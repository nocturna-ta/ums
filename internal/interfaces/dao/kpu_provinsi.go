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
	"time"
)

type KPUProvinsiRepository struct {
	client ethereum.Client
	db     *sql.Store
}

type OptsKPUProvinsiRepository struct {
	Client ethereum.Client
	DB     *sql.Store
}

func NewKPUProvinsiRepository(opts *OptsKPUProvinsiRepository) repository.KPUProvinsiRepository {
	return &KPUProvinsiRepository{
		client: opts.Client,
		db:     opts.DB,
	}
}

const (
	insertKPUProvinsi = `INSERT INTO kpu_provinsi (id, user_id, username, name, address, region, is_active, photo_path, telephone, registered_at, created_at, updated_at)
    						VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11, $12)`
	selectKPUProvinsi = `SELECT %s FROM kpu_provinsi %s WHERE TRUE %s`
	updateKPUProvinsi = `UPDATE kpu_provinsi SET %s WHERE TRUE %s`
)

func (K *KPUProvinsiRepository) InsertKPUProvinsi(ctx context.Context, kpu *model.KPUProvinsi) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiRepository.InsertKPUProvinsi")
	defer span.End()

	var err error

	sqlTrx := utils.GetSqlTx(ctx)

	if sqlTrx != nil {
		_, err = sqlTrx.ExecContext(ctx, insertKPUProvinsi, kpu.ID, kpu.UserID, kpu.Username, kpu.Name, kpu.Address, kpu.Region, kpu.IsActive,
			kpu.PhotoPath, kpu.Telephone, kpu.RegisteredAt, kpu.CreatedAt, kpu.UpdatedAt)
	} else {
		_, err = K.db.GetMaster().ExecContext(ctx, insertKPUProvinsi, kpu.ID, kpu.UserID, kpu.Username, kpu.Name, kpu.Address, kpu.Region, kpu.IsActive,
			kpu.PhotoPath, kpu.Telephone, kpu.RegisteredAt, kpu.CreatedAt, kpu.UpdatedAt)
	}

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
		}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.InsertKPUProvinsi] Failed to insert kpu provinsi")
		return err
	}

	return nil
}

func (K *KPUProvinsiRepository) SendTxKPUProvinsiBlockchain(ctx context.Context, signedTransaction string) (string, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiRepository.SendTxKPUProvinsiBlockchain")
	defer span.End()

	tx, err := utils2.StringToTx(signedTransaction)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.SendTxKPUProvinsiBlockchain] Failed to convert signed transaction to transaction")
		return "", err
	}

	txHash, err := K.client.SendTransaction(ctx, tx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"tx":    tx,
		}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.SendTxKPUProvinsiBlockchain] Failed to send transaction")
		return "", err
	}

	return txHash, nil
}

func (K *KPUProvinsiRepository) GetAllKPUProvinsi(ctx context.Context) ([]*model.KPUProvinsi, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiRepository.GetAllKPUProvinsi")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)
	var (
		kpuProvinsi []*model.KPUProvinsi
		err         error
	)

	selectQuery := `kpu_provinsi.id, kpu_provinsi.user_id, kpu_provinsi.username, kpu_provinsi.name, kpu_provinsi.address, kpu_provinsi.region,
			kpu_provinsi.is_active, kpu_provinsi.photo_path, kpu_provinsi.telephone, kpu_provinsi.registered_at, kpu_provinsi.created_at, kpu_provinsi.updated_at`
	whereQuery := " AND kpu_provinsi.is_deleted = false"
	joinQuery := ""

	query := fmt.Sprintf(selectKPUProvinsi, selectQuery, joinQuery, whereQuery)
	if sqlTrx != nil {
		err = sqlTrx.SelectContext(ctx, &kpuProvinsi, query)
	} else {
		err = K.db.GetMaster().SelectContext(ctx, &kpuProvinsi, query)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.GetAllKPUProvinsi] Failed to get all kpu provinsi")
		return nil, err
	}

	return kpuProvinsi, nil
}

func (K *KPUProvinsiRepository) GetKPUProvinsiByAddress(ctx context.Context, address string) (*model.KPUProvinsi, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiRepository.GetKPUProvinsiByAddress")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)
	var (
		kpuProvinsi model.KPUProvinsi
		err         error
		args        []any
	)

	selectQuery := `kpu_provinsi.id, kpu_provinsi.user_id, kpu_provinsi.username, kpu_provinsi.name, kpu_provinsi.address, kpu_provinsi.region,
	kpu_provinsi.is_active, kpu_provinsi.photo_path, kpu_provinsi.telephone, kpu_provinsi.registered_at, kpu_provinsi.created_at, kpu_provinsi.updated_at`
	whereQuery := " AND kpu_provinsi.is_deleted = false AND kpu_provinsi.address = $1"
	joinQuery := ""
	args = append(args, address)

	query := fmt.Sprintf(selectKPUProvinsi, selectQuery, joinQuery, whereQuery)
	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &kpuProvinsi, query, args...)
	} else {
		err = K.db.GetMaster().GetContext(ctx, &kpuProvinsi, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.GetKPUProvinsiByAddress] Failed to get kpu provinsi by address")
		return nil, err
	}
	return &kpuProvinsi, nil
}

func (K *KPUProvinsiRepository) GetKPUProvinsiByID(ctx context.Context, id uuid.UUID) (*model.KPUProvinsi, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiRepository.GetKPUProvinsiByID")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)
	var (
		kpuProvinsi model.KPUProvinsi
		err         error
		args        []any
	)

	selectQuery := `kpu_provinsi.id, kpu_provinsi.user_id, kpu_provinsi.username, kpu_provinsi.name, kpu_provinsi.address, kpu_provinsi.region,
	kpu_provinsi.is_active, kpu_provinsi.photo_path, kpu_provinsi.telephone, kpu_provinsi.registered_at, kpu_provinsi.created_at, kpu_provinsi.updated_at`
	whereQuery := " AND kpu_provinsi.is_deleted = false AND kpu_provinsi.id = $1"
	joinQuery := ""
	args = append(args, id)

	query := fmt.Sprintf(selectKPUProvinsi, selectQuery, joinQuery, whereQuery)
	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &kpuProvinsi, query, args...)
	} else {
		err = K.db.GetMaster().GetContext(ctx, &kpuProvinsi, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    id,
		}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.GetKPUProvinsiByID] Failed to get kpu provinsi by ID")
		return nil, err
	}

	return &kpuProvinsi, nil
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

func (K *KPUProvinsiRepository) UpdateKPUProvinsi(ctx context.Context, kpu *model.KPUProvinsi) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiRepository.UpdateKPUProvinsi")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	var (
		err    error
		args   []any
		result sql2.Result
	)

	setQuery := "name = $1, region = $2, telephone = $3, username = $4, updated_at = $5"
	whereQuery := " AND id = $6 AND is_deleted = false"
	args = append(args, kpu.Name, kpu.Region, kpu.Telephone, kpu.Username, time.Now(), kpu.ID)
	query := fmt.Sprintf(updateKPUProvinsi, setQuery, whereQuery)

	if sqlTrx != nil {
		result, err = sqlTrx.ExecContext(ctx, query, args...)
	} else {
		result, err = K.db.GetMaster().ExecContext(ctx, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"kpu":   kpu,
		}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.UpdateKPUProvinsi] Failed to update kpu provinsi")
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"kpu":   kpu,
		}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.UpdateKPUProvinsi] Failed to get rows affected")
		return err
	}

	if rowAffected == 0 {
		return ErrNoUpdateHappened
	}

	return nil
}

func (K *KPUProvinsiRepository) GetKPUProvinsiByUserID(ctx context.Context, userID uuid.UUID) (*model.KPUProvinsi, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiRepository.GetKPUProvinsiByID")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)
	var (
		kpuProvinsi model.KPUProvinsi
		err         error
		args        []any
	)

	selectQuery := `kpu_provinsi.id, kpu_provinsi.user_id, kpu_provinsi.username, kpu_provinsi.name, kpu_provinsi.address, kpu_provinsi.region,
	kpu_provinsi.is_active, kpu_provinsi.photo_path, kpu_provinsi.telephone, kpu_provinsi.registered_at, kpu_provinsi.created_at, kpu_provinsi.updated_at`
	whereQuery := " AND kpu_provinsi.is_deleted = false AND kpu_provinsi.user_id = $1"
	joinQuery := ""
	args = append(args, userID)

	query := fmt.Sprintf(selectKPUProvinsi, selectQuery, joinQuery, whereQuery)
	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &kpuProvinsi, query, args...)
	} else {
		err = K.db.GetMaster().GetContext(ctx, &kpuProvinsi, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"user_id": userID,
		}).ErrorWithCtx(ctx, "[KPUProvinsiRepository.GetKPUProvinsiByID] Failed to get kpu provinsi by User ID")
		return nil, err
	}

	return &kpuProvinsi, nil
}
