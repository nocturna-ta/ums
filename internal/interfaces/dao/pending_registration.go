package dao

import (
	"context"
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
	"time"
)

type PendingRegistrationRepository struct {
	db *sql.Store
}

type OptsPendingRegistrationRepository struct {
	DB *sql.Store
}

func NewPendingRegistrationRepository(opts *OptsPendingRegistrationRepository) repository.PendingRegistrationRepository {
	return &PendingRegistrationRepository{
		db: opts.DB,
	}
}

const (
	insertPendingRegistration = `INSERT INTO pending_registrations (id, user_id, role, entity_data, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6)`
	selectPendingRegistration = `SELECT %s FROM pending_registrations WHERE TRUE %s`
	deletePendingRegistration = `UPDATE pending_registrations SET is_deleted = TRUE, updated_at = $1 WHERE id = $2`
)

func (repo *PendingRegistrationRepository) Insert(ctx context.Context, registration *model.PendingRegistration) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "PendingRegistrationRepository.Insert")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	var err error

	if sqlTrx != nil {
		_, err = sqlTrx.ExecContext(ctx, insertPendingRegistration,
			registration.ID,
			registration.UserID,
			registration.Role,
			registration.EntityData,
			registration.CreatedAt,
			registration.UpdatedAt)

	} else {
		_, err = repo.db.GetMaster().ExecContext(ctx, insertPendingRegistration,
			registration.ID,
			registration.UserID,
			registration.Role,
			registration.EntityData,
			registration.CreatedAt,
			registration.UpdatedAt)
	}

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code {
			case "23505":
				log.WithFields(log.Fields{
					"error":        err,
					"registration": registration,
				}).ErrorWithCtx(ctx, "[PendingRegistrationRepository.Insert] Registration already exists")
				return ErrDuplicate
			}
		}
		log.WithFields(log.Fields{
			"error":        err,
			"registration": registration,
		}).ErrorWithCtx(ctx, "[PendingRegistrationRepository.Insert] Failed to insert registration")
		return err
	}

	return nil

}

func (repo *PendingRegistrationRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*model.PendingRegistration, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "PendingRegistrationRepository.GetByUserID")
	defer span.End()

	var (
		err          error
		registration model.PendingRegistration
		args         []any
	)
	sqlTrx := utils.GetSqlTx(ctx)

	selectQuery := "id, user_id, role, entity_data, created_at, updated_at"
	whereClause := " AND user_id = $1"
	args = append(args, userID)
	query := fmt.Sprintf(selectPendingRegistration, selectQuery, whereClause)

	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &registration, query, args...)

	} else {
		err = repo.db.GetMaster().GetContext(ctx, &registration, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userID": userID,
		}).ErrorWithCtx(ctx, "[PendingRegistrationRepository.GetByUserID] Failed to get registration")
		return nil, err
	}

	return &registration, nil
}

func (repo *PendingRegistrationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "PendingRegistrationRepository.Delete")
	defer span.End()

	var err error
	sqlTrx := utils.GetSqlTx(ctx)
	now := time.Now()

	if sqlTrx != nil {
		_, err = sqlTrx.ExecContext(ctx, deletePendingRegistration, now, id)
	} else {
		_, err = repo.db.GetMaster().ExecContext(ctx, deletePendingRegistration, now, id)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    id,
		}).ErrorWithCtx(ctx, "[PendingRegistrationRepository.Delete] Failed to delete registration")
		return err
	}

	return nil
}
