# OtterPrep 🦦

A web application for preparing for JAMB (Joint Admissions and Matriculation Board) exams, powered by JAMB syllabus and past questions.

## Tech Stack

- **Backend:** Go 1.25 with Echo framework
- **Database:** PostgreSQL 16
- **Authentication:** JWT (JSON Web Tokens)
- **Validation:** go-playground/validator
- **Containerization:** Docker & Docker Compose

## Prerequisites

- Go 1.25+
- PostgreSQL 16
- Docker & Docker Compose
- Make (optional)

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/lawson/otterprep.git
cd otterprep
```

### 2. Environment Variables

Create a `.env` file in the `backend/` directory:

```env
# Server
PORT=8080
ENV=development
JWT_SECRET=your-super-secret-key

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=otterprep
DB_PASSWORD=otterprep_password
DB_NAME=otterprep_db
DB_SSL_MODE=disable

# Redis (optional)
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
```

### 3. Run with Docker

```bash
docker compose up
```

This will start:
- PostgreSQL on port `5432`
- PostgreSQL (test) on port `5433`

### 4. Run the Backend

```bash
cd backend
go run cmd/main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### Public Routes

| Method | Endpoint           | Description          |
|--------|-------------------|----------------------|
| POST   | `/user/register`   | Register a new user  |
| POST   | `/user/login`      | User login           |
| POST   | `/admin/register`  | Register admin user  |
| POST   | `/admin/login`     | Admin login          |

### Protected Routes (Requires JWT)

All protected routes require `Authorization: Bearer <token>` header.

#### Admin - Questions

| Method | Endpoint                       | Description                |
|--------|-------------------------------|----------------------------|
| POST   | `/api/v1/admin/questions/bulk`   | Create multiple questions  |
| POST   | `/api/v1/admin/questions/single` | Create single question     |
| GET    | `/api/v1/admin/questions`        | Get all questions          |
| GET    | `/api/v1/admin/questions/:id`    | Get question by ID         |

#### Admin - Subjects

| Method | Endpoint                    | Description          |
|--------|----------------------------|----------------------|
| GET    | `/api/v1/admin/subject`      | Get all subjects     |
| GET    | `/api/v1/admin/subject/:id`  | Get subject by ID    |
| POST   | `/api/v1/admin/subject`      | Create a new subject |

#### Quiz

| Method | Endpoint              | Description      |
|--------|-----------------------|------------------|
| POST   | `/api/v1/quiz/create`  | Create a quiz    |
| GET    | `/api/v1/quiz/submit`  | Submit a quiz    |

## Database Schema

The application uses PostgreSQL with the following tables:

| Table        | Description                          |
|--------------|--------------------------------------|
| `users`      | User accounts                        |
| `user_roles` | User roles (admin, user)             |
| `subjects`   | JAMB subjects (e.g., Mathematics)    |
| `questions`  | Quiz questions                       |
| `options`    | Multiple choice options for questions|
| `answers`    | Explanations for correct answers     |
| `scores`     | User quiz scores and performance     |

Run the schema:

```bash
psql -U otterprep -d otterprep_db -f schema.sql
```

## Project Structure

```
otterprep/
├── docker-compose.yml
├── schema.sql
├── README.md
└── backend/
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── cmd/
    │   └── main.go              # Application entry point
    ├── config/
    │   └── config.go            # Configuration management
    ├── domain/
    │   ├── user.go              # User domain models
    │   ├── quiz.go              # Quiz domain models
    │   ├── question.go          # Question domain models
    │   ├── subject.go           # Subject domain models
    │   └── jwt.go               # JWT models
    ├── internal/
    │   ├── handler/             # HTTP handlers
    │   ├── middleware/          # Auth, validation, error handling
    │   ├── repository/          # Database operations
    │   ├── router/              # Route definitions
    │   └── service/             # Business logic
    └── pkg/
        ├── errors.go            # Custom error definitions
        ├── response.go          # Response helpers
        └── security.go          # Security utilities
```

## Features

- ✅ User authentication (JWT)
- ✅ Role-based access control (Admin/User)
- ✅ Request validation
- ✅ Custom error handling
- ✅ Quiz generation by subject
- ✅ Score tracking
- ✅ Bulk question upload

## License

MIT License
