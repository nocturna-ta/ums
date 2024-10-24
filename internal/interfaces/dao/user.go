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
	insertUser = `INSERT INTO users (id, nik, no_telephone, email, name, password, password_salt, created_at, updated_at)
								VALUES($1, $2, $3, $4, $5, $6, $7,$8,$9)`
	selectUser = `SELECT %s FROM users %s WHERE TRUE %s`
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
		_, err = sqlTrx.ExecContext(ctx, insertUser, user.ID, user.NIK, user.NoTelephone, user.Email, user.Name, user.Password, user.PasswordSalt, user.CreatedAt, user.UpdatedAt)
	} else {
		_, err = repo.db.GetMaster().ExecContext(ctx, insertUser, user.ID, user.NIK, user.NoTelephone, user.Email, user.Name, user.Password, user.PasswordSalt, user.CreatedAt, user.UpdatedAt)
	}

	if err != nil {
		if pqErr, valid := err.(*pq.Error); valid {
			switch pqErr.Code {
			case "23505":
				log.WithFields(log.Fields{
					"error": err,
					"user":  user,
				}).ErrorWithCtx(ctx, "[UserRepository.Insert] Duplicate entry")
				return ErrDuplicate
			}
		}

		log.WithFields(log.Fields{
			"error": err,
			"user":  user,
		}).ErrorWithCtx(ctx, "[UserRepository.Insert] Failed to insert entry")
		return err
	}
	return nil
}
func (repo *UserRepository) GetById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserRepository.GetById")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	selectQuery := `
        users.id as id,
        users.nik as nik,
        users.no_telephone as no_telephone,
        users.email as email,
        users.name as name,
        users.password as password,
        users.password_salt as password_salt,
        users.created_at as created_at,
        users.updated_at as updated_at
    `
	whereQuery := " AND users.id = $1 AND users.is_deleted = false"
	joinQuery := ""
	args := []any{id}

	query := fmt.Sprintf(selectUser, selectQuery, joinQuery, whereQuery)

	var user model.User
	var err error

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
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return &user, nil
}

func (repo *UserRepository) GetByNIK(ctx context.Context, nik string) (*model.User, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserRepository.GetByNIK")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	var (
		user model.User
		args []any
		err  error
	)

	selectQuery := "users.id as id, users.nik as nik, users.no_telephone as no_telephone, users.email as email, users.name as name, users.password as password, users.password_salt as password.salt, users.created_at as created_at, users.updated_at as updated_at"
	whereQuery := " AND users.nik = $1 AND users.id_deleted = false"
	joinQuery := ""
	args = append(args, nik)

	query := fmt.Sprintf(selectUser, selectQuery, joinQuery, whereQuery)

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
			"nik":   nik,
		}).ErrorWithCtx(ctx, "[UserRepository.GetByNIK] Failed to get user by nik")
	}
	return &user, err
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
	whereQuery := " AND users.id = $3 AND users.id_deleted = false"

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
		}).ErrorWithCtx(ctx, "[UserRepository.ChangePassword] Failed to update user")
		return err
	}

	return nil

}
