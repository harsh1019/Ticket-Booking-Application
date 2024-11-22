package main

import (
	"fmt"
	"ticketbookingapp/config"
	"ticketbookingapp/handlers"
	"ticketbookingapp/repositories"
	"ticketbookingapp/db"
	"github.com/gofiber/fiber/v2"
	"ticketbookingapp/services"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"ticketbookingapp/middlewares"
)

func main() {
   
	envConfig := config.NewEnvConfig()
	db := db.GetDBConnectionString(envConfig,db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName: "Ticket Booking App",
		ServerHeader: "Fiber",
	})

	app.Use(logger.New())
	// Repositories
	EventRepository := repositories.NewEventRepository(db)
	TicketRepository := repositories.NewTicketRepository(db)
	AuthRepository := repositories.NewAuthRepository(db)

	//service
	AuthService := services.NewAuthService(AuthRepository)

	//Routing
	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), AuthService)


	//privateRoutes:
	privateRoutes := server.Use(middlewares.AuthProtected(db))
	
	handlers.NewEventHandler(privateRoutes.Group("/events"), EventRepository)
	handlers.NewTicketHandler(privateRoutes.Group("/tickets"), TicketRepository)
	app.Listen(fmt.Sprintf(":%s", envConfig.ServerPort))

}