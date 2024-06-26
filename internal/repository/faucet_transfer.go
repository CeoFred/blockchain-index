package repository

import (
	"fmt"

	"github.com/CeoFred/gin-boilerplate/internal/models"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FaucetTransferInterface interface {
	Find(id uuid.UUID) (*models.FaucetTransfer, error)
	Exists(id uuid.UUID) (bool, error)
	Where(condition, value string) ([]*models.FaucetTransfer, error)
	Create(action *models.FaucetTransfer) error
	Save(action *models.FaucetTransfer) (*models.FaucetTransfer, error)
	RawCount(q string, count *int64) error
	QueryWithArgs(q string, args ...interface{}) (*models.FaucetTransfer, error)
	QueryRecordsWithArgs(q string, args ...interface{}) ([]*models.FaucetTransfer, error)
	RawSmartSelect(q string, res interface{}, args ...interface{}) error
	BatchInsert(events []*models.FaucetTransfer) error
}

type FaucetTransferRepository struct {
	database *gorm.DB
}

func NewFaucetTransferRepository(db *gorm.DB) FaucetTransferInterface {
	return &FaucetTransferRepository{
		database: db,
	}
}

func (a *FaucetTransferRepository) Find(id uuid.UUID) (*models.FaucetTransfer, error) {
	var action models.FaucetTransfer
	err := a.database.First(&action, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &action, nil
}

func (a *FaucetTransferRepository) Exists(id uuid.UUID) (bool, error) {
	var count int64
	err := a.database.Model(&models.FaucetTransfer{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (a *FaucetTransferRepository) Where(condition, value string) ([]*models.FaucetTransfer, error) {
	var actions []*models.FaucetTransfer
	err := a.database.Raw(fmt.Sprintf(`SELECT * FROM faucet_transfers WHERE %s = ?`, condition), value).Scan(&actions).Error
	if err != nil {
		return nil, err
	}
	return actions, nil
}

func (a *FaucetTransferRepository) Create(action *models.FaucetTransfer) error {
	return a.database.Table("faucet_transfers").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "transaction_hash"}, {Name: "block_number"}},
		DoNothing: true, // Ignore conflicts, do nothing
	}).Create(action).Error
}

func (a *FaucetTransferRepository) Save(action *models.FaucetTransfer) (*models.FaucetTransfer, error) {
	err := a.database.Model(action).Where("id = ?", action.ID).Updates(action).Error
	if err != nil {
		return nil, err
	}
	return action, nil
}

func (a *FaucetTransferRepository) RawCount(q string, count *int64) error {
	return a.database.Raw(q).Count(count).Error
}

func (a *FaucetTransferRepository) QueryWithArgs(q string, args ...interface{}) (*models.FaucetTransfer, error) {
	var actions *models.FaucetTransfer
	err := a.database.Raw(q, args...).Find(&actions).Error
	if err != nil {
		return nil, err
	}
	return actions, nil
}
func (a *FaucetTransferRepository) QueryRecordsWithArgs(q string, args ...interface{}) ([]*models.FaucetTransfer, error) {
	var actions []*models.FaucetTransfer
	err := a.database.Raw(q, args...).Find(&actions).Error
	if err != nil {
		return nil, err
	}
	return actions, nil
}

func (a *FaucetTransferRepository) RawSmartSelect(q string, res interface{}, args ...interface{}) error {
	return a.database.Raw(q, args...).Scan(res).Error
}

func (a *FaucetTransferRepository) BatchInsert(events []*models.FaucetTransfer) error {
	return a.database.Table("user_actions").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "transaction_hash"}},
		DoNothing: true, // Ignore conflicts, do nothing
	}).CreateInBatches(events, 50).Error
}
