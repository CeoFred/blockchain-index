package repository

import (
	"fmt"

	"github.com/CeoFred/gin-boilerplate/internal/models"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ContractEventInterface interface {
	Find(id uuid.UUID) (*models.ContractEvent, error)
	Exists(id uuid.UUID) (bool, error)
	Where(condition, value string) ([]*models.ContractEvent, error)
	Create(contract *models.ContractEvent) error
	Save(contract *models.ContractEvent) (*models.ContractEvent, error)
	RawCount(q string, count *int64) error
	QueryWithArgs(q string, args ...interface{}) (*models.ContractEvent, error)
	QueryRecordsWithArgs(q string, args ...interface{}) ([]*models.ContractEvent, error)
	RawSmartSelect(q string, res interface{}, args ...interface{}) error
	BatchInsert(events []*models.ContractEvent) error
}

type ContractEventRepository struct {
	database *gorm.DB
}

func NewContractEventRepository(db *gorm.DB) ContractEventInterface {
	return &ContractEventRepository{
		database: db,
	}
}

func (a *ContractEventRepository) Find(id uuid.UUID) (*models.ContractEvent, error) {
	var contract models.ContractEvent
	err := a.database.First(&contract, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &contract, nil
}

func (a *ContractEventRepository) Exists(id uuid.UUID) (bool, error) {
	var count int64
	err := a.database.Model(&models.ContractEvent{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (a *ContractEventRepository) Where(condition, value string) ([]*models.ContractEvent, error) {
	var contracts []*models.ContractEvent
	err := a.database.Raw(fmt.Sprintf(`SELECT * FROM contracts WHERE %s = ?`, condition), value).Scan(&contracts).Error
	if err != nil {
		return nil, err
	}
	return contracts, nil
}

func (a *ContractEventRepository) Create(contract *models.ContractEvent) error {
	return a.database.Create(contract).Error
}

func (a *ContractEventRepository) Save(contract *models.ContractEvent) (*models.ContractEvent, error) {
	err := a.database.Model(contract).Where("id = ?", contract.ID).Updates(contract).Error
	if err != nil {
		return nil, err
	}
	return contract, nil
}

func (a *ContractEventRepository) RawCount(q string, count *int64) error {
	return a.database.Raw(q).Count(count).Error
}

func (a *ContractEventRepository) QueryWithArgs(q string, args ...interface{}) (*models.ContractEvent, error) {
	var contracts *models.ContractEvent
	err := a.database.Raw(q, args...).Find(&contracts).Error
	if err != nil {
		return nil, err
	}
	return contracts, nil
}

func (a *ContractEventRepository) QueryRecordsWithArgs(q string, args ...interface{}) ([]*models.ContractEvent, error) {
	var contracts []*models.ContractEvent
	err := a.database.Raw(q, args...).Find(&contracts).Error
	if err != nil {
		return nil, err
	}
	return contracts, nil
}

func (a *ContractEventRepository) RawSmartSelect(q string, res interface{}, args ...interface{}) error {
	return a.database.Raw(q, args...).Scan(res).Error
}

func (a *ContractEventRepository) BatchInsert(events []*models.ContractEvent) error {
	return a.database.Table("contract_events").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "contract_id"}, {Name: "event_id"}},
		DoNothing: true, // Ignore conflicts, do nothing
	}).CreateInBatches(events, 10).Error
}
