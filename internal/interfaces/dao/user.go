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
	insertUser = `INSERT INTO users (id, username, email, password, password_salt, role, is_active, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
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
		_, err = sqlTrx.ExecContext(ctx, insertUser, user.ID, user.Username, user.Email, user.Password, user.PasswordSalt, user.Role, user.IsActive, user.CreatedAt, user.UpdatedAt)
	} else {
		_, err = repo.db.GetMaster().ExecContext(ctx, insertUser, user.ID, user.Username, user.Email, user.Password, user.PasswordSalt, user.Role, user.IsActive, user.CreatedAt, user.UpdatedAt)
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

	selectQuery := "users.id, users.username, users.email, users.password, users.password_salt, users.role, users.is_active, users.created_at, users.updated_at"
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

func (repo *UserRepository) Update(ctx context.Context, id uuid.UUID, update *model.UserUpdate) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserRepository.Update")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	var (
		err  error
		args []any
	)

	if update == nil {
		return ErrNilParam
	}

	setQuery := "username = $1, updated_at = $2"
	whereQuery := " AND id = $3 AND is_deleted = FALSE"
	args = append(args, update.Username, time.Now(), id)

	query := fmt.Sprintf(updateUser, setQuery, whereQuery)
	if sqlTrx != nil {
		_, err = sqlTrx.ExecContext(ctx, query, args...)
	} else {
		_, err = repo.db.GetMaster().ExecContext(ctx, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"user":  update,
		}).ErrorWithCtx(ctx, "[UserRepository.Update] Failed to update user")
		return err
	}

	return nil
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

	selectQuery := "users.id, users.username, users.email, users.password, users.password_salt, users.role, users.is_active, users.created_at, users.updated_at"
	whereQuery := " AND users.email = $1 AND users.is_deleted = FALSE"
	args = append(args, email)

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
	}

	return &user, nil
}
