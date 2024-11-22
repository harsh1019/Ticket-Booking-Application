package models

import (
	"time"
	"context"
)

type Ticket struct {
	ID	 uint `json:"id" gorm:"primaryKey"`
	EventID uint `json:"eventId"`
	UserID uint `json:"userId" gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Event Event `json:"event" gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Entered bool `json:"entered" default:"false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TicketRepository interface {
	GetMany(ctx context.Context,userId uint) ([]*Ticket, error)
	GetOne(ctx context.Context, userId uint,ticketId uint) (*Ticket, error)
	CreateOne(ctx context.Context, userId uint,ticket *Ticket) (*Ticket, error)
	UpdateOne(ctx context.Context, userId uint,ticketId uint, updatedData map[string]interface{}) (*Ticket, error)
}

type ValidateTicketRequest struct {	
	TicketID uint `json:"ticketId"`
	OwnerID uint `json:"ownerId"`
}