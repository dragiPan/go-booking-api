# Go Booking API

This is a simple REST API built with Go and Gin for managing events and user registrations. It includes:
- CRUD operations for users and events
- JWT-based authentication and authorization
- Middleware for route protection
- Registration and cancellation of events

## Getting Started

### Prerequisites
- Go 1.23.1 or later
- `git` installed on your system

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/dragiPan/go-booking-api.git
   ```
2. Navigate into the project directory:
   ```bash
   cd go-booking-api
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Run app:
   ```bash
   go run main.go
   ```
### API Endpoints
The API will be available at http://localhost:8080
#### Public Endpoints
| Method | Endpoint   | Description            |
|--------|------------|------------------------|
| POST   | /signup    | Register a new user    |
| POST   | /login     | Authenticate a user    |

#### Protected Endpoints (Requires JWT Token)
| Method   | Endpoint                  | Description                   |
|----------|--------------------------|-------------------------------|
| GET      | /events                  | Get all events               |
| GET      | /events/:id              | Get event by ID              |
| POST     | /events                  | Create a new event           |
| PUT      | /events/:id              | Update an event by ID        |
| DELETE   | /events/:id              | Delete an event by ID        |
| POST     | /events/:id/register     | Register for an event        |
| DELETE   | /events/:id/register     | Cancel a registration        |
