package repository

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/lawson/otterprep/domain"
	"github.com/lawson/otterprep/pkg"
	"github.com/stretchr/testify/assert"
)

func setUpDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	queries := []string{
		"CREATE TABLE users (id integer primary key autoincrement, name text, email text, password_hash text, created_at timestamp, updated_at timestamp)",
		"CREATE TABLE IF NOT EXISTS scores (id INTEGER PRIMARY KEY AUTOINCREMENT, user_id BIGINT, score BIGINT, mode VARCHAR(255), correct_answers BIGINT, incorrect_answers BIGINT, total_questions BIGINT, time_taken_seconds BIGINT, subject_id BIGINT, created_at TIMESTAMP, updated_at TIMESTAMP)",
	}
	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			t.Fatal(err)
		}
	}
	db.SetMaxOpenConns(1)
	return db
}

func TestCreateUser(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()

	userRepo := NewUserRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := userRepo.CreateUser(ctx, domain.User{
		Name:         "John Doe",
		Email:        "john.doe@example.com",
		PasswordHash: "password",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("created user: ", user)
}

func TestGetUserWithID(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()

	userRepo := NewUserRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := userRepo.CreateUser(ctx, domain.User{
		Name:         "John Doe",
		Email:        "john.doe@example.com",
		PasswordHash: "password",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("created user: ", user)

	user, err = userRepo.GetUserWithID(ctx, user.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestUpdateUserPassword(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()

	userRepo := NewUserRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := userRepo.CreateUser(ctx, domain.User{
		Name:         "John Doe",
		Email:        "john.doe@example.com",
		PasswordHash: "password",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("created user: ", user)

	err = userRepo.UpdateUserPassword(ctx, user.ID, "newpassword")
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, err)
	existingUser, err := userRepo.GetUserWithID(ctx, user.ID)
	assert.Nil(t, err)
	assert.Equal(t, true, pkg.CheckPasswordHash("newpassword", existingUser.PasswordHash))
	fmt.Println("updated user: ", existingUser)
}

func TestUpdateUsername(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()

	userRepo := NewUserRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := userRepo.CreateUser(ctx, domain.User{
		Name:         "John Doe",
		Email:        "john.doe@example.com",
		PasswordHash: "password",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("created user: ", user)

	err = userRepo.UpdateUsername(ctx, user.ID, "newusername")
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, err)
	existingUser, err := userRepo.GetUserWithID(ctx, user.ID)
	assert.Nil(t, err)
	assert.Equal(t, "newusername", existingUser.Name)
	fmt.Println("updated user: ", existingUser)
}

func TestDeleteUserByID(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()

	userRepo := NewUserRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := userRepo.CreateUser(ctx, domain.User{
		Name:         "John Doe",
		Email:        "john.doe@example.com",
		PasswordHash: "password",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("created user: ", user)

	err = userRepo.DeleteUserByID(ctx, user.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, err)
	_, err = userRepo.GetUserWithID(ctx, user.ID)
	assert.NotNil(t, err)
}

func TestGetAllUsers(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()

	userRepo := NewUserRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	users := []domain.User{
		{
			Name:         "John Doe",
			Email:        "john.doe@example.com",
			PasswordHash: "password",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "Jane Doe",
			Email:        "jane.doe@example.com",
			PasswordHash: "password",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "Bob Smith",
			Email:        "bob.smith@example.com",
			PasswordHash: "password",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	for _, user := range users {
		_, err := userRepo.CreateUser(ctx, user)
		if err != nil {
			t.Fatal(err)
		}
	}
	fmt.Println("created users: ", users)

	users, err := userRepo.GetAllUsers(ctx)
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, err)
	assert.NotNil(t, users)
	fmt.Println("all users: ", users)
}
