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

| Method | Endpoint           | Description          | Rate Limit |
|--------|-------------------|----------------------|------------|
| POST   | `/user/register`   | Register a new user  | 3/min      |
| POST   | `/user/login`      | User login           | 5/min      |
| POST   | `/admin/register`  | Register admin user  | 3/min      |
| POST   | `/admin/login`     | Admin login          | 5/min      |
| POST   | `/auth/refresh`    | Refresh access token | 10/min     |

### Token Refresh

When your access token expires, use the refresh token to get a new pair:

```bash
POST /auth/refresh
Content-Type: application/json

{
  "refresh_token": "your-refresh-token-here"
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "access_token": "new-access-token",
    "refresh_token": "new-refresh-token",
    "expires_in": 900
  }
}
```

### Protected Routes (Requires JWT)

All protected routes require `Authorization: Bearer <token>` header.

**Rate Limit:** 100 requests per minute per IP for all protected routes.

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

#### User Profile

| Method | Endpoint                  | Description            |
|--------|--------------------------|------------------------|
| GET    | `/api/v1/dashboard`        | Get user dashboard     |
| PUT    | `/api/v1/user/username`    | Update username        |
| PUT    | `/api/v1/user/email`       | Update email           |
| PUT    | `/api/v1/user/password`    | Update password        |
| DELETE | `/api/v1/user/account`     | Delete user account    |

#### Quiz

| Method | Endpoint              | Description      |
|--------|-----------------------|------------------|
| POST   | `/api/v1/quiz/create`  | Create a quiz    |
| POST   | `/api/v1/quiz/submit`  | Submit a quiz    |

#### Leaderboard

| Method | Endpoint                           | Description                    |
|--------|------------------------------------|--------------------------------|
| GET    | `/api/v1/leaderboard`              | Get global leaderboard         |
| GET    | `/api/v1/leaderboard/me`           | Get authenticated user's rank  |
| GET    | `/api/v1/leaderboard/user/:user_id`| Get specific user's rank       |

**Query Parameters:**

| Parameter    | Type   | Default    | Description                                      |
|--------------|--------|------------|--------------------------------------------------|
| `subject_id` | int    | -          | Filter leaderboard by subject                    |
| `period`     | string | `all_time` | Time period: `all_time`, `weekly`, `monthly`     |
| `limit`      | int    | 10         | Number of entries to return (max: 100)           |
| `offset`     | int    | 0          | Pagination offset                                |

**Example Requests:**

```bash
# Get global leaderboard (top 10, all time)
GET /api/v1/leaderboard

# Get weekly leaderboard with top 20 users
GET /api/v1/leaderboard?period=weekly&limit=20

# Get leaderboard for Mathematics subject (subject_id=1)
GET /api/v1/leaderboard?subject_id=1

# Get monthly leaderboard for a subject with pagination
GET /api/v1/leaderboard?subject_id=1&period=monthly&limit=10&offset=10

# Get my rank on the global leaderboard
GET /api/v1/leaderboard/me

# Get my rank for a specific subject
GET /api/v1/leaderboard/me?subject_id=1

# Get another user's rank
GET /api/v1/leaderboard/user/5

# Get another user's rank for a specific subject
GET /api/v1/leaderboard/user/5?subject_id=1
```

**Example Response - Leaderboard:**

```json
{
  "success": true,
  "data": {
    "subject_id": 1,
    "subject_name": "Mathematics",
    "period": "all_time",
    "total_users": 150,
    "entries": [
      {
        "rank": 1,
        "user_id": 42,
        "user_name": "John Doe",
        "total_score": 850,
        "total_quizzes": 25,
        "correct_answers": 170,
        "total_questions": 200,
        "accuracy_percent": 85.0
      }
    ]
  }
}
```

**Example Response - User Rank:**

```json
{
  "success": true,
  "data": {
    "user_id": 5,
    "user_name": "Jane Smith",
    "rank": 12,
    "total_score": 420,
    "total_quizzes": 15,
    "correct_answers": 84,
    "total_questions": 100,
    "accuracy_percent": 84.0,
    "total_users": 150
  }
}
```

**Error Response (User has no scores):**

```json
{
  "success": false,
  "error": "user has no quiz scores yet",
  "status": 404
}
```
- `limit` (optional): Number of entries (default: 10, max: 100)
- `offset` (optional): Pagination offset

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

- ✅ User authentication (JWT with refresh tokens)
- ✅ Role-based access control (Admin/User)
- ✅ Request validation
- ✅ Custom error handling
- ✅ Rate limiting (login, register, API endpoints)
- ✅ Quiz generation by subject
- ✅ Score tracking
- ✅ Bulk question upload
- ✅ User profile management
- ✅ Leaderboard system (global, subject-specific, weekly, monthly)
- ✅ User dashboard with stats

## License

MIT License
