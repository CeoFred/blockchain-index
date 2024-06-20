package repository

import (
	"fmt"

	"github.com/CeoFred/gin-boilerplate/internal/models"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserInterface interface {
	Find(id uuid.UUID) (*models.User, error)
	Exists(id uuid.UUID) (bool, error)
	Where(condition, value string) ([]*models.User, error)
	Create(user *models.User) (*models.User,error)
	Save(action *models.User) (*models.User, error)
	RawCount(q string, count *int64) error
	QueryWithArgs(q string, args ...interface{}) (*models.User, error)
	QueryRecordsWithArgs(q string, args ...interface{}) ([]*models.User, error)
	RawSmartSelect(q string, res interface{}, args ...interface{}) error
	BatchInsert(events []*models.User) error
}

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserInterface {
	return &UserRepository{
		database: db,
	}
}

func (a *UserRepository) Find(id uuid.UUID) (*models.User, error) {
	var action models.User
	err := a.database.First(&action, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &action, nil
}

func (a *UserRepository) Exists(id uuid.UUID) (bool, error) {
	var count int64
	err := a.database.Model(&models.User{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (a *UserRepository) Where(condition, value string) ([]*models.User, error) {
	var actions []*models.User
	err := a.database.Raw(fmt.Sprintf(`SELECT * FROM users WHERE %s = ?`, condition), value).Scan(&actions).Error
	if err != nil {
		return nil, err
	}
	return actions, nil
}

func (a *UserRepository) Create(user *models.User) (*models.User,error) {
	err := a.database.Table("users").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "address"}},
		DoNothing: true,
	}).Create(user).Error

	if err != nil {
		return nil,err
	}
	var u models.User
	err = a.database.Table("users").Where("address= ?", user.Address).Find(&u).Error
	if err != nil {
		return nil,err
	}

	return &u,nil
}

func (a *UserRepository) Save(action *models.User) (*models.User, error) {
	err := a.database.Model(action).Where("id = ?", action.ID).Updates(action).Error
	if err != nil {
		return nil, err
	}
	return action, nil
}

func (a *UserRepository) RawCount(q string, count *int64) error {
	return a.database.Raw(q).Count(count).Error
}

func (a *UserRepository) QueryWithArgs(q string, args ...interface{}) (*models.User, error) {
	var actions *models.User
	err := a.database.Raw(q, args...).Find(&actions).Error
	if err != nil {
		return nil, err
	}
	return actions, nil
}

func (a *UserRepository) QueryRecordsWithArgs(q string, args ...interface{}) ([]*models.User, error) {
	var actions []*models.User
	err := a.database.Raw(q, args...).Find(&actions).Error
	if err != nil {
		return nil, err
	}
	return actions, nil
}

func (a *UserRepository) RawSmartSelect(q string, res interface{}, args ...interface{}) error {
	return a.database.Raw(q, args...).Scan(res).Error
}

func (a *UserRepository) BatchInsert(events []*models.User) error {
	return a.database.Table("users").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "address"}},
		DoNothing: true,
	}).CreateInBatches(events, 50).Error
}
