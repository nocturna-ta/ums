package repository

import (
	"context"
	"github.com/nocturna-ta/ums/internal/domain/model"
)

type VoterRepository interface {
	InsertVoter(ctx context.Context, voter *model.Voter) error
	GetAllVoter(ctx context.Context) ([]model.Voter, error)
	GetVoterByNIK(ctx context.Context, nik string) (*model.Voter, error)
	GetVoterByAddress(ctx context.Context, address string) (*model.Voter, error)
	GetVoterByRegion(ctx context.Context, region string) ([]model.Voter, error)
}
