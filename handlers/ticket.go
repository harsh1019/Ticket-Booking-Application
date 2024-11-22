package handlers

import (
	"context"
	"strconv"
	"ticketbookingapp/models"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
	"fmt"
)

type TicketHandler struct {
	repository models.TicketRepository
}

func (h *TicketHandler) GetMany(c *fiber.Ctx) error {
	context,cancel := context.WithTimeout(context.Background(),time.Duration(5*time.Second))
	defer cancel()

	userId := uint(c.Locals("userId").(float64))
	tickets,err := h.repository.GetMany(context,userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  "Failed",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Success",
		"status":  "Success",
		"data":    tickets,
	})
}

func (h *TicketHandler) GetOne(c *fiber.Ctx) error {

	ticketId,_ := strconv.Atoi(c.Params("ticketId"))
	userId := uint(c.Locals("userId").(float64))
	context,cancel := context.WithTimeout(context.Background(),time.Duration(5*time.Second))
	defer cancel()

	ticket,err := h.repository.GetOne(context,userId,uint(ticketId))

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"message": err.Error(),
			"status":  "Failed",
		})
	}

	var QRCode []byte
	QRCode,err = qrcode.Encode(
		fmt.Sprintf("ticketId:%v,ownerId:%v", ticketId, userId),
		qrcode.Medium,
		256,
	)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  "Failed",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Success",
		"status":  "Success",
		"data":	&fiber.Map{
			"ticket": ticket,
			"qrcode": QRCode,
		},
	})
}

func (h *TicketHandler) CreateOne(c *fiber.Ctx) error {
	ticket := &models.Ticket{}
	userId := uint(c.Locals("userId").(float64))
	context,cancel := context.WithTimeout(context.Background(),time.Duration(5*time.Second))
	defer cancel()

	if err := c.BodyParser(ticket); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
			"status":  "Failed",
			"data":nil,
		})
	}

	createdTicket,err := h.repository.CreateOne(context,userId,ticket)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  "Failed",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"message": "Ticket Created Successfully",
		"status":  "Success",
		"data":	createdTicket,
	})

}

func (h *TicketHandler) ValidateOne(c *fiber.Ctx) error {
	ValidateBody := &models.ValidateTicketRequest{}

	context,cancel := context.WithTimeout(context.Background(),time.Duration(5*time.Second))
	defer cancel()

	if err := c.BodyParser(ValidateBody); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": err.Error(),
			"status":  "Failed!!",
			"data":nil,
		})
	}

	
	validateData := make(map[string]interface{})
	validateData["entered"] = true

	Ticket,err := h.repository.UpdateOne(context,ValidateBody.OwnerID,ValidateBody.TicketID,validateData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  "Failed!!!",
			"data":nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Welcome to the show!!!",
		"status":  "Success",
		"data":	Ticket,
	})
}


func NewTicketHandler(router fiber.Router,repository models.TicketRepository){
	handler := &TicketHandler{
		repository: repository,
	}


	router.Get("/",handler.GetMany)
	router.Get("/:ticketId",handler.GetOne)
	router.Post("/",handler.CreateOne)
	router.Post("/validate",handler.ValidateOne)
}

