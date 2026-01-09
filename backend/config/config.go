package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
}

type ServerConfig struct {
	Port      string
	Env       string
	JWTSecret string
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

func Load() (*Config, error) {
	// Load .env file if it exists (ignore error in production)
	_ = godotenv.Load()

	cfg := &Config{
		Server: ServerConfig{
			Port:      getEnv("PORT", "8080"),
			Env:       getEnv("ENV", "development"),
			JWTSecret: getEnv("JWT_SECRET", "your-secret-key"),
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

// getEnv returns the value of the environment variable with the given key
// If the environment variable is not set, it returns the default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
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
