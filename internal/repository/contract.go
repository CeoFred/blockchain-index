package repository

import (
	"fmt"

	"github.com/CeoFred/gin-boilerplate/internal/models"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type ContractInterface interface {
	Find(id uuid.UUID) (*models.Contract, error)
	Exists(id uuid.UUID) (bool, error)
	Where(condition, value string) ([]*models.Contract, error)
	Create(contract *models.Contract) error
	Save(contract *models.Contract) (*models.Contract, error)
	RawCount(q string, count *int64) error
	QueryWithArgs(q string, args ...interface{}) (*models.Contract, error)
	QueryRecordsWithArgs(q string, args ...interface{}) ([]*models.Contract, error)
	RawSmartSelect(q string, res interface{}, args ...interface{}) error
}

type ContractRepository struct {
	database *gorm.DB
}

func NewContractRepository(db *gorm.DB) ContractInterface {
	return &ContractRepository{
		database: db,
	}
}

func (a *ContractRepository) Find(id uuid.UUID) (*models.Contract, error) {
	var contract models.Contract
	err := a.database.First(&contract, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &contract, nil
}

func (a *ContractRepository) Exists(id uuid.UUID) (bool, error) {
	var count int64
	err := a.database.Model(&models.Contract{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (a *ContractRepository) Where(condition, value string) ([]*models.Contract, error) {
	var contracts []*models.Contract
	err := a.database.Raw(fmt.Sprintf(`SELECT * FROM contracts WHERE %s = ?`, condition), value).Scan(&contracts).Error
	if err != nil {
		return nil, err
	}
	return contracts, nil
}

func (a *ContractRepository) Create(contract *models.Contract) error {
	return a.database.Create(contract).Error
}

func (a *ContractRepository) Save(contract *models.Contract) (*models.Contract, error) {
	err := a.database.Model(contract).Where("id = ?", contract.ID).Updates(contract).Error
	if err != nil {
		return nil, err
	}
	return contract, nil
}

func (a *ContractRepository) RawCount(q string, count *int64) error {
	return a.database.Raw(q).Count(count).Error
}

func (a *ContractRepository) QueryWithArgs(q string, args ...interface{}) (*models.Contract, error) {
	var contracts *models.Contract
	err := a.database.Raw(q, args...).Find(&contracts).Error
	if err != nil {
		return nil, err
	}

	if contracts.ID.IsNil() {
		return nil, nil
	}

	return contracts, nil
}

func (a *ContractRepository) QueryRecordsWithArgs(q string, args ...interface{}) ([]*models.Contract, error) {
	var contracts []*models.Contract
	err := a.database.Raw(q, args...).Find(&contracts).Error
	if err != nil {
		return nil, err
	}
	return contracts, nil
}

func (a *ContractRepository) RawSmartSelect(q string, res interface{}, args ...interface{}) error {
	return a.database.Raw(q, args...).Scan(res).Error
}
