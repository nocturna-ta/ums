package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type PendingRegistration struct {
	BaseModel
	ID                uuid.UUID       `db:"id"`
	UserID            uuid.UUID       `db:"user_id"`
	Role              string          `db:"role"`
	EntityData        json.RawMessage `db:"entity_data"`
	SignedTransaction string          `db:"signed_transaction"`
}

type KPUProvinsiData struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Region  string `json:"region"`
}

type KPUKotaData struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Region  string `json:"region"`
}

type VoterData struct {
	NIK          string `json:"nik"`
	VoterAddress string `json:"voter_address"`
	Region       string `json:"region"`
}

func NewPendingRegist(userID uuid.UUID, role string, signedTx string, data interface{}) (*PendingRegistration, error) {
	now := time.Now()

	entityDataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return &PendingRegistration{
		BaseModel: BaseModel{
			CreatedAt: now,
			UpdatedAt: now,
			IsDeleted: false,
		},
		ID:                uuid.New(),
		UserID:            userID,
		Role:              role,
		EntityData:        entityDataJSON,
		SignedTransaction: signedTx,
	}, nil
}

func (pr *PendingRegistration) GetKPUProvinsiData() (*KPUProvinsiData, error) {
	var data KPUProvinsiData
	err := json.Unmarshal(pr.EntityData, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (pr *PendingRegistration) GetKPUKotaData() (*KPUKotaData, error) {
	var data KPUKotaData
	err := json.Unmarshal(pr.EntityData, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (pr *PendingRegistration) GetVoterData() (*VoterData, error) {
	var data VoterData
	err := json.Unmarshal(pr.EntityData, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
