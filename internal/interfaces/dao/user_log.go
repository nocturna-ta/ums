package dao

import (
	"context"
	"fmt"
	"github.com/nocturna-ta/golib/database/sql"
	log2 "github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/golib/txmanager/utils"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"github.com/nocturna-ta/ums/internal/domain/repository"
)

type UserLogRepository struct {
	db *sql.Store
}

type OptsUserLogRepository struct {
	DB *sql.Store
}

func NewUserLogRepository(opts *OptsUserLogRepository) repository.UserLogRepository {
	return &UserLogRepository{
		db: opts.DB,
	}
}

const (
	insertUserLog = `INSERT INTO user_logs (id, user_id, username, name, role, ` + "`time`" + `, activity, activity_type) 
    VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	selectUserLog = `SELECT %s FROM user_logs WHERE TRUE %s`
)

func (u *UserLogRepository) InsertLog(ctx context.Context, log *model.UserLogs) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserLogRepository.InsertLog")
	defer span.End()

	var (
		err error
	)

	sqlTrx := utils.GetSqlTx(ctx)

	if sqlTrx != nil {
		_, err = sqlTrx.ExecContext(ctx, insertUserLog,
			log.ID,
			log.UserID,
			log.Username,
			log.Name,
			log.Role,
			log.Time,
			log.Activity,
			log.ActivityType,
		)
	} else {
		_, err = u.db.GetMaster().ExecContext(ctx, insertUserLog,
			log.ID,
			log.UserID,
			log.Username,
			log.Name,
			log.Role,
			log.Time,
			log.Activity,
			log.ActivityType,
		)
	}

	if err != nil {
		log2.WithFields(log2.Fields{
			"error":    err,
			"user_log": log,
		}).ErrorWithCtx(ctx, "[UserLogRepository.InsertLog] Failed to insert user log")
		return err
	}

	return nil
}

func (u *UserLogRepository) GetUserLogs(ctx context.Context, limit, offset int) ([]model.UserLogs, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserLogRepository.GetUserLogs")
	defer span.End()

	var (
		logs []model.UserLogs
		err  error
		args []any
	)

	sqlTrx := utils.GetSqlTx(ctx)

	selectQuery := `id, user_id, username, name, role, ` + "`time`" + `, activity, activity_type`
	whereQuery := " ORDER BY `time` LIMIT ? OFFSET ?"

	args = append(args, limit, offset)

	query := fmt.Sprintf(selectUserLog, selectQuery, whereQuery)

	fmt.Println("Executing query:", query, "with args:", args)

	if sqlTrx != nil {
		err = sqlTrx.SelectContext(ctx, &logs, query, args...)
	} else {
		err = u.db.GetMaster().SelectContext(ctx, &logs, query, args...)
	}

	if err != nil {
		log2.WithFields(log2.Fields{
			"error": err,
			"logs":  logs,
		}).ErrorWithCtx(ctx, "[UserLogRepository.GetUserLogs] Failed to get user logs")
		return nil, err
	}

	return logs, nil
}
