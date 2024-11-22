package repositories

import (
	"context"
	"gorm.io/gorm"
	"ticketbookingapp/models"
)

type EventRepository struct{
	db *gorm.DB
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error){
	// get all events
	//get all events in postgress db
	events := []*models.Event{}
	res := r.db.Model(&models.Event{}).Order("updated_at desc").Find(&events)

	if res.Error != nil {
		return nil, res.Error
	}

	return events, nil
}

func (r *EventRepository) GetOne(ctx context.Context, eventId uint) (*models.Event, error){
	// get one event
	event := &models.Event{}

	res := r.db.Model(event).Where("id = ?", eventId).First(&event)
	if res.Error != nil {
		return nil, res.Error
	}

	return event, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event *models.Event) (*models.Event, error){
	// create one event
	res := r.db.Model(event).Create(&event)
	if res.Error != nil {
		return nil, res.Error
	}

	return event, nil
}


func (r *EventRepository) UpdateOne(ctx context.Context, eventId uint, updatedData map[string]interface{}) (*models.Event, error){
	// update one event
	event := &models.Event{}
	res := r.db.Model(event).Where("id = ?", eventId).Updates(updatedData)
	if res.Error != nil {
		return nil, res.Error
	}


	getres := r.db.Where("id = ?", eventId).First(&event)
	if getres.Error != nil {
		return nil, getres.Error
	}

	return event, nil
}

func (r *EventRepository) DeleteOne(ctx context.Context, eventId uint) error{
	// delete one event
	event := &models.Event{}
	res := r.db.Model(event).Where("id = ?", eventId).Delete(&event)
	return res.Error
}


func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{
		db: db,
	}
}