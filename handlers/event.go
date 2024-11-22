package handlers

import (
	"context"
	"strconv"
	"ticketbookingapp/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type EventHandler struct {
	repository models.EventRepository
}

func (h *EventHandler) GetMany(c *fiber.Ctx) error {
	
	context,cancel := context.WithTimeout(context.Background(),time.Duration(5*time.Second))
	defer cancel()

	events, err := h.repository.GetMany(context)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"status":  "Failed",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Success",
		"status":  "Success",
		"data":    events,
	})
}

func (h *EventHandler) GetOne(c *fiber.Ctx) error {
	eventId,_ := strconv.Atoi(c.Params("eventId"))

	context,cancel := context.WithTimeout(context.Background(),time.Duration(5*time.Second))
	defer cancel()

	event,err := h.repository.GetOne(context,uint(eventId))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"status":  "Failed",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Success",
		"status":  "Success",
		"data":	event,
	})

}

func (h *EventHandler) CreateOne(c *fiber.Ctx) error {
  event := &models.Event{}

  context,cancel := context.WithTimeout(context.Background(),time.Duration(5*time.Second))
  defer cancel()

  if err := c.BodyParser(event); err != nil {
	return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
	  "message": err.Error(),
	  "status":  "Failed",
	})
  }

  createdEvent,err := h.repository.CreateOne(context,event)
  if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	  "message": err.Error(),
	  "status":  "Failed",
  })}


  return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
	"message": "Event created successfully",
	"status":  "Success",
	"data":    createdEvent,
  })

}

func (h *EventHandler) UpdateOne(c *fiber.Ctx) error {
eventId,_ := strconv.Atoi(c.Params("eventId"))
updatedData := make(map[string]interface{})


context,cancel := context.WithTimeout(context.Background(),time.Duration(5*time.Second))
defer cancel()

if err := c.BodyParser(&updatedData); err != nil {
	return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
		"message": err.Error(),
		"status":  "Failed",
		"data":	nil,
	})
}


event,err := h.repository.UpdateOne(context,uint(eventId),updatedData)
if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": err.Error(),
		"status":  "Failed",
	})
}

return c.Status(fiber.StatusOK).JSON(&fiber.Map{
	"message": "Event updated successfully",
	"status":  "Success",
	"data":    event,
})}


func (h *EventHandler) DeleteOne(c *fiber.Ctx) error {
	eventId,_ := strconv.Atoi(c.Params("eventId"))

	context,cancel := context.WithTimeout(context.Background(),time.Duration(5*time.Second))
	defer cancel()

	err := h.repository.DeleteOne(context,uint(eventId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"status":  "Failed",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Event deleted successfully",
		"status":  "Success",
	})
}

func NewEventHandler(router fiber.Router,repository models.EventRepository) {
	handler := &EventHandler{
		repository: repository,
	}
	
	router.Get("/",handler.GetMany)
	router.Get("/:eventId", handler.GetOne)
	router.Post("/", handler.CreateOne)
	router.Put("/:eventId", handler.UpdateOne)
	router.Delete("/:eventId", handler.DeleteOne)

}