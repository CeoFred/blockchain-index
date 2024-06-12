package repository

import (
	"fmt"

	"github.com/CeoFred/gin-boilerplate/internal/models"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type EventInterface interface {
	Find(id uuid.UUID) (*models.Event, error)
	Exists(id uuid.UUID) (bool, error)
	Where(condition, value string) ([]*models.Event, error)
	Create(contract *models.Event) error
	Save(contract *models.Event) (*models.Event, error)
	RawCount(q string, count *int64) error
	QueryWithArgs(q string, args ...interface{}) (*models.Event, error)
	QueryRecordsWithArgs(q string, args ...interface{}) ([]*models.Event, error)
	RawSmartSelect(q string, res interface{}, args ...interface{}) error
}

type EventRepository struct {
	database *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventInterface {
	return &EventRepository{
		database: db,
	}
}

func (a *EventRepository) Find(id uuid.UUID) (*models.Event, error) {
	var contract models.Event
	err := a.database.First(&contract, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &contract, nil
}

func (a *EventRepository) Exists(id uuid.UUID) (bool, error) {
	var count int64
	err := a.database.Model(&models.Event{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (a *EventRepository) Where(condition, value string) ([]*models.Event, error) {
	var contracts []*models.Event
	err := a.database.Raw(fmt.Sprintf(`SELECT * FROM events WHERE %s = ?`, condition), value).Scan(&contracts).Error
	if err != nil {
		return nil, err
	}
	return contracts, nil
}

func (a *EventRepository) Create(contract *models.Event) error {
	return a.database.Create(contract).Error
}

func (a *EventRepository) Save(contract *models.Event) (*models.Event, error) {
	err := a.database.Model(contract).Where("id = ?", contract.ID).Updates(contract).Error
	if err != nil {
		return nil, err
	}
	return contract, nil
}

func (a *EventRepository) RawCount(q string, count *int64) error {
	return a.database.Raw(q).Count(count).Error
}

func (a *EventRepository) QueryWithArgs(q string, args ...interface{}) (*models.Event, error) {
	var contracts *models.Event
	err := a.database.Raw(q, args...).Find(&contracts).Error
	if err != nil {
		return nil, err
	}
	return contracts, nil
}

func (a *EventRepository) QueryRecordsWithArgs(q string, args ...interface{}) ([]*models.Event, error) {
	var contracts []*models.Event
	err := a.database.Raw(q, args...).Find(&contracts).Error
	if err != nil {
		return nil, err
	}
	return contracts, nil
}

func (a *EventRepository) RawSmartSelect(q string, res interface{}, args ...interface{}) error {
	return a.database.Raw(q, args...).Scan(res).Error
}
