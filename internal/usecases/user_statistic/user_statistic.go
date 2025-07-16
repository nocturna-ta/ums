package user_statistic

import (
	"context"
	"fmt"
	libCtx "github.com/nocturna-ta/golib/context"
	"github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"github.com/nocturna-ta/ums/internal/usecases/response"
)

func (m *Module) GetApprovedDPTStatistic(ctx context.Context, region string) (*response.ApprovedDPTResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticUseCases.GetApprovedDPTStatistic")
	defer span.End()

	var regionPtr *string
	if region != "" {
		regionPtr = &region
	}

	total, err := m.userStatisticRepo.GetDPTTotal(ctx, regionPtr)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetApprovedDPTStatistic] GetDPTTotal failed")
		return nil, err
	}

	approved, err := m.userStatisticRepo.GetCountDPTByStatus(ctx, "approved", regionPtr)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetApprovedDPTStatistic] GetCountDPTByStatus failed")
		return nil, err
	}

	var percentage float64 = 0
	if total > 0 {
		percentage = float64(approved) / float64(total) * 100
	}

	return &response.ApprovedDPTResponse{
		Percentage:       percentage,
		TotalApprovedDPT: approved,
	}, nil
}

func (m *Module) GetRejectedDPTStatistic(ctx context.Context, region string) (*response.RejectedDPTResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticUseCases.GetRejectedDPTStatistic")
	defer span.End()

	var regionPtr *string
	if region != "" {
		regionPtr = &region
	}

	total, err := m.userStatisticRepo.GetDPTTotal(ctx, regionPtr)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetRejectedDPTStatistic] GetDPTTotal failed")
		return nil, err
	}

	rejected, err := m.userStatisticRepo.GetCountDPTByStatus(ctx, "rejected", regionPtr)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetRejectedDPTStatistic] GetCountDPTByStatus failed")
		return nil, err
	}

	var percentage float64 = 0
	if total > 0 {
		percentage = float64(rejected) / float64(total) * 100
	}

	return &response.RejectedDPTResponse{
		Percentage:       percentage,
		TotalRejectedDPT: rejected,
	}, nil
}

func (m *Module) GetTotalDPTStatistic(ctx context.Context, region string) (*response.TotalDPTResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticUseCases.GetTotalDPTStatistic")
	defer span.End()

	var regionPtr *string
	if region != "" {
		regionPtr = &region
	}

	total, err := m.userStatisticRepo.GetDPTTotal(ctx, regionPtr)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetTotalDPTStatistic] GetDPTTotal failed")
		return nil, err
	}

	approved, err := m.userStatisticRepo.GetCountDPTByStatus(ctx, "approved", regionPtr)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetTotalDPTStatistic] GetCountDPTByStatus for approved failed")
		return nil, err
	}

	pending, err := m.userStatisticRepo.GetCountDPTByStatus(ctx, "pending", regionPtr)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetTotalDPTStatistic] GetCountDPTByStatus for pending failed")
		return nil, err
	}

	rejected, err := m.userStatisticRepo.GetCountDPTByStatus(ctx, "rejected", regionPtr)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetTotalDPTStatistic] GetCountDPTByStatus for rejected failed")
		return nil, err
	}

	return &response.TotalDPTResponse{
		TotalDPT:    total,
		PendingDPT:  pending,
		ApprovedDPT: approved,
		RejectedDPT: rejected,
	}, nil
}

func (m *Module) GetStaffKPUProvinceStatistic(ctx context.Context, region string) (*response.StaffKPUResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticUseCases.GetStaffKPUStatistic")
	defer span.End()

	var regionPtr *string
	if region != "" {
		regionPtr = &region
	}

	kpuStaff, err := m.userStatisticRepo.GetKPUProvinsiStaff(ctx, regionPtr)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetStaffKPUStatistic] GetKPUProvinsiStaff failed")
		return nil, err
	}

	provinceDB, err := m.userStatisticRepo.GetProvinceCount(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetStaffKPUStatistic] GetProvinceCount failed")
		return nil, err
	}

	return &response.StaffKPUResponse{
		Count:       kpuStaff,
		RegionCount: provinceDB,
	}, nil
}

func (m *Module) GetStaffKPUKotaStatistic(ctx context.Context, region string) (*response.StaffKPUResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticUseCases.GetStaffKPUKotaStatistic")
	defer span.End()

	var regionPtr *string
	if region != "" {
		regionPtr = &region
	}

	kpuStaff, err := m.userStatisticRepo.GetKPUKotaStaff(ctx, regionPtr)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetStaffKPUKotaStatistic] GetKPUKotaStaff failed")
		return nil, err
	}

	districtDB, err := m.userStatisticRepo.GetDistrictCount(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetStaffKPUKotaStatistic] GetDistrictCount failed")
		return nil, err
	}

	return &response.StaffKPUResponse{
		Count:       kpuStaff,
		RegionCount: districtDB,
	}, nil
}

func (m *Module) GetPendingDPTStatistic(ctx context.Context, region string) (*response.PendingDPTResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticUseCases.GetPendingDPTStatistic")
	defer span.End()

	var regionPtr *string
	if region != "" {
		regionPtr = &region
	}

	total, err := m.userStatisticRepo.GetDPTTotal(ctx, regionPtr)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetPendingDPTStatistic] GetDPTTotal failed")
		return nil, err
	}

	pending, err := m.userStatisticRepo.GetCountDPTByStatus(ctx, "pending", regionPtr)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetPendingDPTStatistic] GetCountDPTByStatus failed")
		return nil, err
	}

	var percentage float64 = 0
	if total > 0 {
		percentage = float64(pending) / float64(total) * 100
	}

	return &response.PendingDPTResponse{
		Percentage:      percentage,
		TotalPendingDPT: pending,
	}, nil
}

func (m *Module) GetProvinceInformationDPTStatistic(ctx context.Context) (*[]response.DPTInformationResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticUseCases.GetInformationDPTStatistic")
	defer span.End()

	provinces, err := m.kpuProvinsiRepo.GetAllKPUProvinsi(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetInformationDPTStatistic] GetAllKPUProvinsi failed")
		return nil, err
	}

	provinceMap := make(map[string]model.KPUProvinsi)
	for _, province := range provinces {
		provinceMap[province.Region] = *province
	}

	var dptInfoResponses []response.DPTInformationResponse
	for _, province := range provinceMap {
		cities, err := m.getCitiesInProvince(ctx, province.Region)
		if err != nil {
			log.WithFields(log.Fields{
				"error":    err,
				"province": province.Name,
			}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetInformationDPTStatistic] getCitiesInProvince failed")
			return nil, err
		}

		fmt.Println(cities)

		staffCount, err := m.userStatisticRepo.GetKPUProvinsiStaff(ctx, &province.Region)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetInformationDPTStatistic] GetKPUProvinsiStaff failed")
			return nil, err
		}

		var dpt, votecount int

		for _, city := range cities {
			totalDPT, err := m.userStatisticRepo.GetDPTTotal(ctx, &city)
			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
				}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetInformationDPTStatistic] GetDPTTotal failed")
				return nil, err
			}

			votedCount, err := m.userStatisticRepo.GetDPTVoted(ctx, &city)
			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
				}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetInformationDPTStatistic] GetDPTVoted failed")
				return nil, err
			}

			votecount += votedCount
			dpt += totalDPT
		}

		var votedPercentage float64 = 0
		if dpt > 0 {
			votedPercentage = float64(votecount) / float64(dpt) * 100
		} else {
			votedPercentage = 0
		}

		dptInfoResponses = append(dptInfoResponses, response.DPTInformationResponse{
			KPURegion:          province.Region,
			StaffCount:         staffCount,
			DPTVotedPercentage: votedPercentage,
			TotalDPT:           dpt,
		})
	}

	return &dptInfoResponses, nil
}

func (m *Module) GetKotaInformationDPTStatistic(ctx context.Context) (*[]response.DPTInformationResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticUseCases.GetKotaInformationDPTStatistic")
	defer span.End()

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetKotaInformationDPTStatistic] GetRequestContext failed")
		return nil, err
	}

	kpuProvinsi, err := m.kpuProvinsiRepo.GetKPUProvinsiByUserID(ctx, reqCtx.GetUserId())
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetKotaInformationDPTStatistic] GetKPUProvinsiByUserID failed")
		return nil, err
	}

	cities, err := m.getCitiesInProvince(ctx, kpuProvinsi.Region)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetKotaInformationDPTStatistic] getCitiesInProvince failed")
		return nil, err
	}

	voters, err := m.voterRepo.GetAllVoter(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetKotaInformationDPTStatistic] Get All Voter Failed")
		return nil, err
	}

	var filteredCities []string
	for _, city := range cities {
		for _, voter := range voters {
			if voter.Region == city {
				filteredCities = append(filteredCities, city)
				break // No need to check other voters for this city
			}
		}

	}

	var dptInfoResponses []response.DPTInformationResponse
	for _, kota := range filteredCities {
		staffCount, err := m.userStatisticRepo.GetKPUKotaStaff(ctx, &kota)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetKotaInformationDPTStatistic] GetKPUKotaStaff failed")
			return nil, err
		}

		totalDPT, err := m.userStatisticRepo.GetDPTTotal(ctx, &kota)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetKotaInformationDPTStatistic] GetDPTTotal failed")
			return nil, err
		}

		votedCount, err := m.userStatisticRepo.GetDPTVoted(ctx, &kota)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetKotaInformationDPTStatistic] GetDPTVoted failed")
			return nil, err
		}

		var votedPercentage float64 = 0
		if totalDPT > 0 {
			votedPercentage = float64(votedCount) / float64(totalDPT) * 100
		}

		dptInfoResponses = append(dptInfoResponses, response.DPTInformationResponse{
			KPURegion:          kota,
			StaffCount:         staffCount,
			DPTVotedPercentage: votedPercentage,
			TotalDPT:           totalDPT,
		})
	}

	return &dptInfoResponses, nil
}

func (m *Module) GetVotedStatistic(ctx context.Context, region string) (*response.VotedStatisticResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticUseCases.GetVotedStatistic")
	defer span.End()

	var regionPtr *string
	if region != "" {
		regionPtr = &region
	}

	votedCount, err := m.userStatisticRepo.GetDPTVoted(ctx, regionPtr)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetVotedStatistic] GetDPTVoted failed")
		return nil, err
	}

	totalDPT, err := m.userStatisticRepo.GetDPTTotal(ctx, regionPtr)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.GetVotedStatistic] GetDPTTotal failed")
		return nil, err
	}

	var percentage float64 = 0
	if totalDPT > 0 {
		percentage = float64(votedCount) / float64(totalDPT) * 100
	}

	return &response.VotedStatisticResponse{
		Percentage: percentage,
		TotalDPT:   totalDPT,
	}, nil
}

func (m *Module) getCitiesInProvince(ctx context.Context, province string) ([]string, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticUseCases.getCitiesInProvince")
	defer span.End()

	cities, err := m.wilayahAPIClient.GetRegenciesByProvinceName(ctx, province)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err,
			"province": province,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.getCitiesInProvince] GetRegenciesByProvinceName failed")
	}

	allKpuKota, err := m.kpuKotaRepo.GetAllKPUKota(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[UserStatisticUseCases.getCitiesInProvince] GetAllKPUKota failed")
		return nil, err
	}

	kpuKotaMap := make(map[string]model.KPUKota)
	for _, kota := range allKpuKota {
		kpuKotaMap[kota.Region] = *kota
	}

	var citiesToAdd []string

	for _, city := range cities {
		citiesToAdd = append(citiesToAdd, city.Name)
	}

	return citiesToAdd, nil
}
