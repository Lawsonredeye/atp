package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	Email    EmailConfig
}

type ServerConfig struct {
	Port         string
	Env          string
	JWTSecret    string
	AllowOrigins []string
	FrontendURL  string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type EmailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
	FromName string
}

func Load() (*Config, error) {
	// Load .env file if it exists (ignore error in production)
	if err := godotenv.Load(); err != nil {
		log.Printf("Info: .env file not loaded: %v", err)
	} else {
		log.Println("Info: .env file loaded successfully")
	}

	cfg := &Config{
		Server: ServerConfig{
			Port:         getEnv("PORT", "8080"),
			Env:          getEnv("ENV", "development"),
			JWTSecret:    getEnv("JWT_SECRET", "your-secret-key"),
			AllowOrigins: getEnvSlice("CORS_ALLOWED_ORIGINS", []string{"*"}),
			FrontendURL:  getEnv("FRONTEND_URL", "http://localhost:5173"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "otterprep"),
			Password: getEnv("DB_PASSWORD", "otterprep_password"),
			DBName:   getEnv("DB_NAME", "otterprep_db"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       0,
		},
		Email: EmailConfig{
			Host:     getEnv("SMTP_HOST", "smtp-relay.brevo.com"),
			Port:     getEnvInt("SMTP_PORT", 587),
			Username: getEnv("SMTP_USERNAME", ""),
			Password: getEnv("SMTP_PASSWORD", ""),
			From:     getEnv("SMTP_FROM", "noreply@acethatpaper.com"),
			FromName: getEnv("SMTP_FROM_NAME", "AceThatPaper"),
		},
	}

	return cfg, nil
}

func (c *DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
}

func (c *DatabaseConfig) PostgresInit() *sql.DB {
	connStr := c.Connect()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open database connection: ", err)
	}
	// Validate the connection by pinging
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	return db
}

func (c *RedisConfig) Address() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func (c *DatabaseConfig) Connect() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
}

// Connect returns the connection string for Redis
// Format: host:port
func (c *RedisConfig) Connect() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

// RedisInit creates and returns a new Redis client
func (c *RedisConfig) RedisInit() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Address(),
		Password: c.Password,
		DB:       c.DB,
	})

	// Validate the connection by pinging
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		log.Printf("Warning: Failed to connect to Redis: %v", err)
	}

	return client
}

// getEnv returns the value of the environment variable with the given key
// If the environment variable is not set, it returns the default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvInt returns the integer value of the environment variable with the given key
// If the environment variable is not set or invalid, it returns the default value
func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

// getEnvSlice returns a slice of strings from a comma-separated environment variable
func getEnvSlice(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		return splitAndTrim(value, ",")
	}
	return defaultValue
}

// splitAndTrim splits a string by separator and trims whitespace from each part
func splitAndTrim(s string, sep string) []string {
	var result []string
	start := 0
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			part := trim(s[start:i])
			if part != "" {
				result = append(result, part)
			}
			start = i + len(sep)
		}
	}
	// Add the last part
	part := trim(s[start:])
	if part != "" {
		result = append(result, part)
	}
	return result
}

// trim removes leading and trailing whitespace
func trim(s string) string {
	start := 0
	end := len(s)
	for start < end && (s[start] == ' ' || s[start] == '\t') {
		start++
	}
	for end > start && (s[end-1] == ' ' || s[end-1] == '\t') {
		end--
	}
	return s[start:end]
}

// parseDuration returns the duration value of the given string
// If the string is not a valid duration, it returns the default duration
func parseDuration(value string, defaultDuration time.Duration) time.Duration {
	d, err := time.ParseDuration(value)
	if err != nil {
		return defaultDuration
	}
	return d
}
