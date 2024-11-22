package repositories

import (
	"context"
	"fmt"
	"ticketbookingapp/models"

	"gorm.io/gorm"
)

type Ticketrepository struct{
	db *gorm.DB
}


func (r *Ticketrepository) GetMany(ctx context.Context, userId uint) ([]*models.Ticket, error){
    
	tickets := []*models.Ticket{}
	res := r.db.Model(&models.Ticket{}).Where("user_id = ?", userId).Preload("Event").Order("updated_at desc").Find(&tickets)

	if res.Error != nil {
		return nil, res.Error
	}
	fmt.Println(tickets)
	return tickets, nil
}

func (r *Ticketrepository) GetOne(ctx context.Context,userId uint,ticketId uint) (*models.Ticket, error){
   
	tickets := &models.Ticket{}
	res := r.db.Model(tickets).Where("id = ?", ticketId).Where("user_id = ?",userId).Preload("Event").First(tickets)
	
	if res.Error != nil {
		fmt.Println(res.Error)
		return nil, res.Error
	}

	
	return tickets, nil
}

func (r *Ticketrepository) CreateOne(ctx context.Context,userId uint,ticket *models.Ticket) (*models.Ticket, error){
    
	ticket.UserID = userId
	res := r.db.Model(ticket).Create(ticket)
	if res.Error != nil {
		return nil, res.Error
	}

	return r.GetOne(ctx, userId, ticket.ID)
}

func (r *Ticketrepository) UpdateOne(ctx context.Context,userId uint,ticketId uint, updatedData map[string]interface{}) (*models.Ticket, error){
	ticket := &models.Ticket{}
	res := r.db.Model(ticket).Where("id = ?", ticketId).Updates(updatedData)
	if res.Error != nil {
		return nil, res.Error
	}
	
	return r.GetOne(ctx,userId, ticketId)
}


func NewTicketRepository(db *gorm.DB) models.TicketRepository {
	return &Ticketrepository{
		db: db,
	}
}