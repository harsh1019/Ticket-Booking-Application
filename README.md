# 🎟️ Ticket Booking Application - Go Server

This is a Go server for a ticket booking application built using the GoFiber framework. The application supports user registration and login using JWT, event creation, ticket booking, and QR code validation.

## ✨ Features

- 📝 User Registration and Login with JWT
- 🎉 Create New Events
- 🎫 Book Tickets for Events
- 📷 QR Code Scanning and Validations

## 🚀 Getting Started

### 📋 Prerequisites

- 🐹 Go 1.16 or higher
- 🐘 A PostgreSQL database

### ⚙️ Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/TicketBookingApp.git
    cd TicketBookingApp/server
    ```

2. Create a `.env` file with the following environment variables:
    ```env
    SERVER_PORT=3000
    DB_HOST=your_db_host
    DB_NAME=your_db_name
    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_SSLMODE=disable
    JWT_SECRET=your_jwt_secret
    ```

3. Build and run the application:
    ```sh
    go build -o ./bin ./cmd/api/main.go
    ./bin/main
    ```

### 📖 Usage

- The server will start on the port specified in the `.env` file (default is 3000).
- Use the provided endpoints to register, login, create events, and book tickets.

## 🔗 Endpoints

- `POST /register` - Register a new user
- `POST /login` - Login a user and receive a JWT token
- `GET /events` - Retrieve all events
- `GET /events/:eventId` - Retrieve a specific event by ID
- `POST /events` - Create a new event (requires JWT)
- `PUT /events/:eventId` - Update an existing event by ID (requires JWT)
- `DELETE /events/:eventId` - Delete an event by ID (requires JWT)
- `GET /tickets` - Retrieve all tickets
- `GET /tickets/:ticketId` - Retrieve a specific ticket by ID
- `POST /tickets` - Book a ticket for an event (requires JWT)
- `POST /tickets/validate` - Validate a ticket using QR code (requires JWT)

## Dependencies

- [GoFiber](https://gofiber.io/) - Web framework for Go
- [JWT](https://jwt.io/) - JSON Web Tokens for authentication
