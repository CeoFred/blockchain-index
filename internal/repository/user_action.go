package repository

import (
	"fmt"

	"github.com/CeoFred/gin-boilerplate/internal/models"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserActionInterface interface {
	Find(id uuid.UUID) (*models.UserAction, error)
	Exists(id uuid.UUID) (bool, error)
	Where(condition, value string) ([]*models.UserAction, error)
	Create(action *models.UserAction) error
	Save(action *models.UserAction) (*models.UserAction, error)
	RawCount(q string, count *int64) error
	QueryWithArgs(q string, args ...interface{}) (*models.UserAction, error)
	QueryRecordsWithArgs(q string, args ...interface{}) ([]*models.UserAction, error)
	RawSmartSelect(q string, res interface{}, args ...interface{}) error
	BatchInsert(events []*models.UserAction) error
}

type UserActionRepository struct {
	database *gorm.DB
}

func NewUserActionRepository(db *gorm.DB) UserActionInterface {
	return &UserActionRepository{
		database: db,
	}
}

func (a *UserActionRepository) Find(id uuid.UUID) (*models.UserAction, error) {
	var action models.UserAction
	err := a.database.First(&action, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &action, nil
}

func (a *UserActionRepository) Exists(id uuid.UUID) (bool, error) {
	var count int64
	err := a.database.Model(&models.UserAction{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (a *UserActionRepository) Where(condition, value string) ([]*models.UserAction, error) {
	var actions []*models.UserAction
	err := a.database.Raw(fmt.Sprintf(`SELECT * FROM user_actions WHERE %s = ?`, condition), value).Scan(&actions).Error
	if err != nil {
		return nil, err
	}
	return actions, nil
}

func (a *UserActionRepository) Create(action *models.UserAction) error {
	return a.database.Table("user_actions").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "transaction_hash"}},
		DoNothing: true, // Ignore conflicts, do nothing
	}).Create(action).Error
}

func (a *UserActionRepository) Save(action *models.UserAction) (*models.UserAction, error) {
	err := a.database.Model(action).Where("id = ?", action.ID).Updates(action).Error
	if err != nil {
		return nil, err
	}
	return action, nil
}

func (a *UserActionRepository) RawCount(q string, count *int64) error {
	return a.database.Raw(q).Count(count).Error
}

func (a *UserActionRepository) QueryWithArgs(q string, args ...interface{}) (*models.UserAction, error) {
	var actions *models.UserAction
	err := a.database.Raw(q, args...).Find(&actions).Error
	if err != nil {
		return nil, err
	}
	return actions, nil
}
func (a *UserActionRepository) QueryRecordsWithArgs(q string, args ...interface{}) ([]*models.UserAction, error) {
	var actions []*models.UserAction
	err := a.database.Raw(q, args...).Find(&actions).Error
	if err != nil {
		return nil, err
	}
	return actions, nil
}

func (a *UserActionRepository) RawSmartSelect(q string, res interface{}, args ...interface{}) error {
	return a.database.Raw(q, args...).Scan(res).Error
}

func (a *UserActionRepository) BatchInsert(events []*models.UserAction) error {
	return a.database.Table("user_actions").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "transaction_hash"}},
		DoNothing: true, // Ignore conflicts, do nothing
	}).CreateInBatches(events, 50).Error
}
