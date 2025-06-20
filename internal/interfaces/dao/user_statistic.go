package dao

import (
	"context"
	"fmt"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/golib/txmanager/utils"
	"github.com/nocturna-ta/ums/internal/domain/repository"
)

type UserStatisticRepository struct {
	db *sql.Store
}

type OptsUserStatisticRepository struct {
	DB *sql.Store
}

func NewUserStatisticRepository(opts *OptsUserStatisticRepository) repository.UserStatisticRepository {
	return &UserStatisticRepository{
		db: opts.DB,
	}
}

func (u *UserStatisticRepository) GetCountDPTByStatus(ctx context.Context, status string, region *string) (int, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticRepository.GetCountDPTByStatus")
	defer span.End()

	var (
		count       int
		err         error
		args        []any
		whereClause string
	)

	sqlTrx := utils.GetSqlTx(ctx)

	selectQuery := "COUNT(*)"
	joinClause := " LEFT JOIN voters ON users.id = voters.user_id"
	if region != nil {
		whereClause = " AND users.verification_status = $1 AND voters.region = $2 AND users.requested_role = 'voter'"
		args = append(args, status, *region)
	} else {
		whereClause = " AND users.verification_status = $1 AND users.requested_role = 'voter'"
		args = append(args, status)
	}

	query := fmt.Sprintf(selectUser, selectQuery, joinClause, whereClause)

	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &count, query, args...)
	} else {
		err = u.db.GetMaster().GetContext(ctx, &count, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"status": status,
		}).ErrorWithCtx(ctx, "[UserStatisticRepository.GetCountDPTByStatus] failed to get count of DPT by status")
		return 0, err
	}

	return count, nil
}

func (u *UserStatisticRepository) GetDPTTotal(ctx context.Context, region *string) (int, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticRepository.GetDPTTotal")
	defer span.End()

	var (
		count       int
		err         error
		args        []any
		whereClause string
	)

	selectQuery := "COUNT(*)"
	joinClause := "  LEFT JOIN voters ON users.id = voters.user_id"
	if region != nil {
		whereClause = " AND voters.region = $1 AND users.requested_role = 'voter'"
		args = append(args, *region)
	} else {
		whereClause = " AND users.requested_role = 'voter'"
	}

	query := fmt.Sprintf(selectUser, selectQuery, joinClause, whereClause)

	sqlTrx := utils.GetSqlTx(ctx)

	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &count, query, args...)
	} else {
		err = u.db.GetMaster().GetContext(ctx, &count, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticRepository.GetDPTTotal] failed to get DPT total count")
		return 0, err
	}

	return count, nil
}

func (u *UserStatisticRepository) GetDPTVoted(ctx context.Context, region *string) (int, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticRepository.GetDPTVoted")
	defer span.End()

	var (
		count       int
		err         error
		args        []any
		whereClause string
	)

	selectQuery := "COUNT(*)"
	if region != nil {
		whereClause = " AND region = $1 AND has_voted = TRUE"
		args = append(args, *region)
	} else {
		whereClause = " AND has_voted = TRUE"
	}

	query := fmt.Sprintf(selectVoter, selectQuery, "", whereClause)

	sqlTrx := utils.GetSqlTx(ctx)

	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &count, query, args...)
	} else {
		err = u.db.GetMaster().GetContext(ctx, &count, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticRepository.GetDPTVoted] failed to get DPT voted count")
		return 0, err
	}

	return count, nil
}

func (u *UserStatisticRepository) GetDPTNotVoted(ctx context.Context, region *string) (int, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticRepository.GetDPTNotVoted")
	defer span.End()

	var (
		count       int
		err         error
		args        []any
		whereClause string
	)

	selectQuery := "COUNT(*)"
	if region != nil {
		whereClause = " AND region = $1 AND has_voted = FALSE"
		args = append(args, *region)
	} else {
		whereClause = " AND has_voted = FALSE"
	}

	query := fmt.Sprintf(selectVoter, selectQuery, "", whereClause)

	sqlTrx := utils.GetSqlTx(ctx)

	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &count, query, args...)
	} else {
		err = u.db.GetMaster().GetContext(ctx, &count, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticRepository.GetDPTNotVoted] failed to get DPT not voted count")
		return 0, err
	}

	return count, nil
}

func (u *UserStatisticRepository) GetKPUProvinsiApprovedUsers(ctx context.Context) (int, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticRepository.GetKPUProvinsiApprovedUsers")
	defer span.End()

	var (
		count int
		err   error
	)

	selectQuery := "COUNT(*)"
	whereClause := " AND status = 'approved' AND role = 'kpu_provinsi'"

	query := fmt.Sprintf(selectUser, selectQuery, "", whereClause)

	sqlTrx := utils.GetSqlTx(ctx)

	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &count, query)
	} else {
		err = u.db.GetMaster().GetContext(ctx, &count, query)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticRepository.GetKPUProvinsiApprovedUsers] failed to get KPU Provinsi approved users count")
		return 0, err
	}

	return count, nil
}

func (u *UserStatisticRepository) GetKPUProvinsiStaff(ctx context.Context, region *string) (int, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticRepository.GetKPUProvinsiStaff")
	defer span.End()

	var (
		count       int
		err         error
		args        []any
		whereClause string
	)

	selectQuery := "COUNT(*)"
	if region != nil {
		whereClause = " AND region = $1"
		args = append(args, *region)
	} else {
		whereClause = ""
	}

	query := fmt.Sprintf(selectKPUProvinsi, selectQuery, "", whereClause)

	sqlTrx := utils.GetSqlTx(ctx)

	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &count, query, args...)
	} else {
		err = u.db.GetMaster().GetContext(ctx, &count, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticRepository.GetKPUProvinsiStaff] failed to get KPU Provinsi staff count")
		return 0, err
	}

	return count, nil
}

func (u *UserStatisticRepository) GetKPUKotaApprovedUsers(ctx context.Context) (int, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticRepository.GetKPUKotaApprovedUsers")
	defer span.End()

	var (
		count int
		err   error
	)

	selectQuery := "COUNT(*)"
	whereClause := " AND status = 'approved' AND role = 'kpu_kota'"

	query := fmt.Sprintf(selectUser, selectQuery, "", whereClause)

	sqlTrx := utils.GetSqlTx(ctx)

	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &count, query)
	} else {
		err = u.db.GetMaster().GetContext(ctx, &count, query)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticRepository.GetKPUKotaApprovedUsers] failed to get KPU Kota approved users count")
		return 0, err
	}

	return count, nil
}

func (u *UserStatisticRepository) GetKPUKotaStaff(ctx context.Context, region *string) (int, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticRepository.GetKPUKotaStaff")
	defer span.End()

	var (
		count       int
		err         error
		args        []any
		whereClause string
	)

	selectQuery := "COUNT(*)"
	if region != nil {
		whereClause = " AND region = $1"
		args = append(args, *region)
	} else {
		whereClause = ""
	}

	query := fmt.Sprintf(selectKPUKota, selectQuery, "", whereClause)

	sqlTrx := utils.GetSqlTx(ctx)

	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &count, query, args...)
	} else {
		err = u.db.GetMaster().GetContext(ctx, &count, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticRepository.GetKPUKotaStaff] failed to get KPU Kota staff count")
		return 0, err
	}

	return count, nil
}

func (u *UserStatisticRepository) GetProvinceCount(ctx context.Context) (int, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticRepository.GetProvinceCount")
	defer span.End()

	var (
		count int
		err   error
	)

	sqlTrx := utils.GetSqlTx(ctx)

	selectQuery := "COUNT(DISTINCT region)"

	query := fmt.Sprintf(selectKPUProvinsi, selectQuery, "", "")

	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &count, query)
	} else {
		err = u.db.GetMaster().GetContext(ctx, &count, query)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticRepository.GetProvinceCount] failed to get province count")
		return 0, err
	}
	return count, nil
}

func (u *UserStatisticRepository) GetDistrictCount(ctx context.Context) (int, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticRepository.GetDistrictCount")
	defer span.End()

	var (
		count int
		err   error
	)

	sqlTrx := utils.GetSqlTx(ctx)

	selectQuery := "COUNT(DISTINCT region)"

	query := fmt.Sprintf(selectKPUKota, selectQuery, "", "")

	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &count, query)
	} else {
		err = u.db.GetMaster().GetContext(ctx, &count, query)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticRepository.GetDistrictCount] failed to get district count")
		return 0, err
	}
	return count, nil
}
