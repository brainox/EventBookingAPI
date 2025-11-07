# Event Booking API

A RESTful API built with Go and Gin framework for managing events and user registrations. Users can create, update, and delete events, as well as register for events created by others.

## Features

- **User Authentication**
  - User signup with password hashing (bcrypt)
  - User login with JWT token generation
  - Protected routes with JWT middleware

- **Event Management**
  - Create, read, update, and delete events
  - Events are tied to the creating user
  - Only event owners can update or delete their events

- **Event Registration**
  - Users can register for events
  - Users can cancel their event registrations
  - Many-to-many relationship between users and events

## Tech Stack

- **Language:** Go 1.24
- **Framework:** Gin Web Framework
- **Database:** SQLite3
- **Authentication:** JWT (golang-jwt/jwt)
- **Password Hashing:** bcrypt

## Prerequisites

- Go 1.24 or higher
- SQLite3

## Installation

1. Clone the repository:
```bash
git clone https://github.com/brainox/EventBookingAPI.git
cd EventBookingAPI
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### Public Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/events` | Get all events |
| GET | `/events/:id` | Get a single event by ID |
| POST | `/signup` | Create a new user account |
| POST | `/login` | Login and receive JWT token |

### Protected Endpoints (Require Authentication)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/events` | Create a new event |
| PUT | `/events/:id` | Update an event (owner only) |
| DELETE | `/events/:id` | Delete an event (owner only) |
| POST | `/events/:id/register` | Register for an event |
| DELETE | `/events/:id/register` | Cancel event registration |

## Authentication

All protected endpoints require a JWT token in the `Authorization` header:

```
Authorization: <your-jwt-token>
```

### Getting a Token

1. **Sign up:**
```bash
POST /signup
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "yourpassword"
}
```

2. **Login:**
```bash
POST /login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "yourpassword"
}
```

Response:
```json
{
  "message": "Login successful.",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

## Usage Examples

### Create an Event

```bash
POST /events
Authorization: <your-token>
Content-Type: application/json

{
  "name": "Tech Conference 2025",
  "description": "Annual technology conference",
  "location": "San Francisco, CA",
  "datetime": "2025-12-15T09:00:00Z"
}
```

### Register for an Event

```bash
POST /events/1/register
Authorization: <your-token>
```

### Update an Event

```bash
PUT /events/1
Authorization: <your-token>
Content-Type: application/json

{
  "name": "Updated Event Name",
  "description": "Updated description",
  "location": "New Location",
  "datetime": "2025-12-20T10:00:00Z"
}
```

### Cancel Registration

```bash
DELETE /events/1/register
Authorization: <your-token>
```

## Database Schema

### Users Table
- `id` - INTEGER (Primary Key, Auto Increment)
- `email` - TEXT (Unique, Not Null)
- `password` - TEXT (Hashed, Not Null)

### Events Table
- `id` - INTEGER (Primary Key, Auto Increment)
- `name` - TEXT (Not Null)
- `description` - TEXT (Not Null)
- `location` - TEXT (Not Null)
- `datetime` - DATETIME (Not Null)
- `user_id` - INTEGER (Foreign Key → users.id)

### Registrations Table
- `id` - INTEGER (Primary Key, Auto Increment)
- `event_id` - INTEGER (Foreign Key → events.id)
- `user_id` - INTEGER (Foreign Key → users.id)

## Project Structure

```
EventBookingAPI/
├── api-test/              # HTTP test files
├── database/              # Database initialization
│   └── database.go
├── middlewares/           # Middleware functions
│   └── auth.go
├── models/                # Data models
│   ├── event.go
│   └── user.go
├── routes/                # Route handlers
│   ├── events.go
│   ├── register.go
│   ├── routes.go
│   └── users.go
├── utils/                 # Utility functions
│   ├── hash.go
│   └── jwt.go
├── main.go                # Application entry point
├── go.mod                 # Go module dependencies
└── README.md
```

## Error Handling

The API returns appropriate HTTP status codes:

- `200` - OK (Successful GET, PUT, DELETE)
- `201` - Created (Successful POST)
- `400` - Bad Request (Invalid input)
- `401` - Unauthorized (Missing/invalid token)
- `403` - Forbidden (Not authorized to perform action)
- `404` - Not Found (Resource doesn't exist)
- `500` - Internal Server Error (Server-side error)

## Security Notes

- Passwords are hashed using bcrypt before storage
- JWT tokens expire after 2 hours
- The JWT secret key should be changed in production (currently hardcoded in `utils/jwt.go`)
- Consider using environment variables for sensitive configuration

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Author

**brainox**