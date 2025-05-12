package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type PendingRegistration struct {
	BaseModel
	ID         uuid.UUID       `db:"id"`
	UserID     uuid.UUID       `db:"user_id"`
	Role       string          `db:"role"`
	EntityData json.RawMessage `db:"entity_data"`
}

type KPUProvinsiData struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Region    string `json:"region"`
	Telephone string `json:"telephone"`
}

type KPUKotaData struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Region    string `json:"region"`
	Telephone string `json:"telephone"`
}

type VoterData struct {
	NIK          string `json:"nik"`
	Fullname     string `json:"fullname"`
	Gender       string `json:"gender"`
	BirthPlace   string `json:"birth_place"`
	BirthDate    string `json:"birth_date"`
	Residential  string `json:"residential_address"`
	VoterAddress string `json:"voter_address"`
	Region       string `json:"region"`
	KTPPhotoPath string `json:"ktp_photo_path"`
}

func NewPendingRegist(userID uuid.UUID, role string, data interface{}) (*PendingRegistration, error) {
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
		ID:         uuid.New(),
		UserID:     userID,
		Role:       role,
		EntityData: entityDataJSON,
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
