package repository

import (
	"fmt"

	"github.com/CeoFred/gin-boilerplate/internal/models"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FaucetInterface interface {
	Find(id uuid.UUID) (*models.Faucet, error)
	Exists(id uuid.UUID) (bool, error)
	Where(condition, value string) ([]*models.Faucet, error)
	Create(action *models.Faucet) error
	Save(action *models.Faucet) (*models.Faucet, error)
	RawCount(q string, count *int64) error
	QueryWithArgs(q string, args ...interface{}) (*models.Faucet, error)
	QueryRecordsWithArgs(q string, args ...interface{}) ([]*models.Faucet, error)
	RawSmartSelect(q string, res interface{}, args ...interface{}) error
	BatchInsert(events []*models.Faucet) error
}

type FaucetRepository struct {
	database *gorm.DB
}

func NewFaucetRepository(db *gorm.DB) FaucetInterface {
	return &FaucetRepository{
		database: db,
	}
}

func (a *FaucetRepository) Find(id uuid.UUID) (*models.Faucet, error) {
	var action models.Faucet
	err := a.database.First(&action, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &action, nil
}

func (a *FaucetRepository) Exists(id uuid.UUID) (bool, error) {
	var count int64
	err := a.database.Model(&models.Faucet{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (a *FaucetRepository) Where(condition, value string) ([]*models.Faucet, error) {
	var faucet []*models.Faucet
	err := a.database.Raw(fmt.Sprintf(`SELECT * FROM faucets WHERE %s = ?`, condition), value).Scan(&faucet).Error
	if err != nil {
		return nil, err
	}
	return faucet, nil
}

func (a *FaucetRepository) Create(action *models.Faucet) error {
	return a.database.Table("faucets").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "address"}},
		DoNothing: true, // Ignore conflicts, do nothing
	}).Create(action).Error
}

func (a *FaucetRepository) Save(action *models.Faucet) (*models.Faucet, error) {
	err := a.database.Model(action).Where("id = ?", action.ID).Updates(action).Error
	if err != nil {
		return nil, err
	}
	return action, nil
}

func (a *FaucetRepository) RawCount(q string, count *int64) error {
	return a.database.Raw(q).Count(count).Error
}

func (a *FaucetRepository) QueryWithArgs(q string, args ...interface{}) (*models.Faucet, error) {
	var faucet *models.Faucet
	err := a.database.Raw(q, args...).Find(&faucet).Error
	if err != nil {
		return nil, err
	}
	if faucet.Asset == "" {
		return nil, nil
	}
	return faucet, nil
}
func (a *FaucetRepository) QueryRecordsWithArgs(q string, args ...interface{}) ([]*models.Faucet, error) {
	var faucet []*models.Faucet
	err := a.database.Raw(q, args...).Find(&faucet).Error
	if err != nil {
		return nil, err
	}
	return faucet, nil
}

func (a *FaucetRepository) RawSmartSelect(q string, res interface{}, args ...interface{}) error {
	return a.database.Raw(q, args...).Scan(res).Error
}

func (a *FaucetRepository) BatchInsert(events []*models.Faucet) error {
	return a.database.Table("user_faucet").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "transaction_hash"}},
		DoNothing: true, // Ignore conflicts, do nothing
	}).CreateInBatches(events, 50).Error
}
