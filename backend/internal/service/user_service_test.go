package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/lawson/otterprep/domain"
	"github.com/lawson/otterprep/internal/repository"
	"github.com/lawson/otterprep/pkg"
	"github.com/stretchr/testify/assert"
)

func TestUserServiceCreateUser(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userRepo := repository.NewUserRepository(pool)
	scoreRepo := repository.NewScoreRepository(pool)
	userService := NewUserService(*userRepo, scoreRepo, log.New(os.Stdout, "", 0))

	user := domain.User{
		Name:         "test",
		Email:        "test@example.com",
		PasswordHash: "test1011",
	}

	createdUser, err := userService.CreateUserAccount(ctx, user, "admin")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("created user:", createdUser)

	assert.Equal(t, user.Name, createdUser.Name)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, "", createdUser.PasswordHash)
}

func TestUserServiceGetUserWithID(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	userRepo := repository.NewUserRepository(pool)
	scoreRepo := repository.NewScoreRepository(pool)
	userService := NewUserService(*userRepo, scoreRepo, log.New(os.Stdout, "", 0))
	newUser := domain.User{
		Name:         "test",
		Email:        "test@example.com",
		PasswordHash: "test1001",
	}

	createdUser, err := userService.CreateUserAccount(ctx, newUser, "admin")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("created user:", createdUser)

	assert.Equal(t, newUser.Name, createdUser.Name)
	assert.Equal(t, newUser.Email, createdUser.Email)
	assert.Equal(t, "", createdUser.PasswordHash)

	user, err := userService.GetUserWithID(ctx, createdUser.ID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("user: %+v\n", user)

	assert.Equal(t, user.Name, createdUser.Name)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, "", createdUser.PasswordHash)
}

func TestUserServiceUpdateUsername(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	userRepo := repository.NewUserRepository(pool)
	scoreRepo := repository.NewScoreRepository(pool)
	userService := NewUserService(*userRepo, scoreRepo, log.New(os.Stdout, "", 0))
	newUser := domain.User{
		Name:         "test",
		Email:        "test@email.com",
		PasswordHash: "user101",
	}

	createdUser, err := userService.CreateUserAccount(ctx, newUser, domain.UserUser)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("created user:", createdUser)

	assert.Equal(t, newUser.Name, createdUser.Name)
	assert.Equal(t, newUser.Email, createdUser.Email)
	assert.Equal(t, "", createdUser.PasswordHash)

	err = userService.UpdateUsername(ctx, createdUser.ID, "newtest")
	if err != nil {
		t.Fatal(err)
	}

	updatedUser, err := userService.GetUserWithID(ctx, createdUser.ID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("updated user: %+v\n", updatedUser)

	assert.Equal(t, "newtest", updatedUser.Name)
	assert.Equal(t, newUser.Email, updatedUser.Email)
	assert.Equal(t, "", updatedUser.PasswordHash)
}

func TestUserServiceUpdateUserEmail(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	userRepo := repository.NewUserRepository(pool)
	scoreRepo := repository.NewScoreRepository(pool)
	userService := NewUserService(*userRepo, scoreRepo, log.New(os.Stdout, "", 0))
	newUser := domain.User{
		Name:         "test",
		Email:        "test@email.com",
		PasswordHash: "user101",
	}

	createdUser, err := userService.CreateUserAccount(ctx, newUser, domain.UserUser)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("created user:", createdUser)

	assert.Equal(t, newUser.Name, createdUser.Name)
	assert.Equal(t, newUser.Email, createdUser.Email)
	assert.Equal(t, "", createdUser.PasswordHash)

	err = userService.UpdateEmail(ctx, createdUser.ID, "newtest@email.com")
	if err != nil {
		t.Fatal(err)
	}

	updatedUser, err := userService.GetUserWithID(ctx, createdUser.ID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("updated user: %+v\n", updatedUser)

	assert.Equal(t, "newtest@email.com", updatedUser.Email)
	assert.Equal(t, newUser.Name, updatedUser.Name)
	assert.Equal(t, "", updatedUser.PasswordHash)
}

func TestUserServiceUpdatePassword(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	userRepo := repository.NewUserRepository(pool)
	scoreRepo := repository.NewScoreRepository(pool)
	userService := NewUserService(*userRepo, scoreRepo, log.New(os.Stdout, "", 0))
	newUser := domain.User{
		Name:         "test",
		Email:        "test@email.com",
		PasswordHash: "user101",
	}

	createdUser, err := userService.CreateUserAccount(ctx, newUser, domain.UserUser)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("created user:", createdUser)

	assert.Equal(t, newUser.Name, createdUser.Name)
	assert.Equal(t, newUser.Email, createdUser.Email)
	assert.Equal(t, "", createdUser.PasswordHash)

	err = userService.UpdatePassword(ctx, createdUser.ID, "newuser101")
	if err != nil {
		t.Fatal(err)
	}

	updatedUser, err := userService.userRepo.GetUserWithID(ctx, createdUser.ID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("updated user: %+v\n", updatedUser)

	// This is only possible since I called the actual userRepo to fetch the data without removing the password hash for the sake of testing
	assert.Equal(t, true, pkg.CheckPasswordHash("newuser101", updatedUser.PasswordHash))
	assert.Equal(t, newUser.Email, updatedUser.Email)
}

func TestUserServiceDeleteUserByID(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	userRepo := repository.NewUserRepository(pool)
	scoreRepo := repository.NewScoreRepository(pool)
	userService := NewUserService(*userRepo, scoreRepo, log.New(os.Stdout, "", 0))
	newUser := domain.User{
		Name:         "test",
		Email:        "test@email.com",
		PasswordHash: "user101",
	}

	createdUser, err := userService.CreateUserAccount(ctx, newUser, domain.UserUser)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("created user:", createdUser)

	assert.Equal(t, newUser.Name, createdUser.Name)
	assert.Equal(t, newUser.Email, createdUser.Email)
	assert.Equal(t, "", createdUser.PasswordHash)

	err = userService.DeleteUserByID(ctx, createdUser.ID)
	assert.Nil(t, err)

	_, err = userService.GetUserWithID(ctx, createdUser.ID)
	assert.NotNil(t, err)
}

func TestUserServiceGetAllUsers(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	userRepo := repository.NewUserRepository(pool)
	scoreRepo := repository.NewScoreRepository(pool)
	userService := NewUserService(*userRepo, scoreRepo, log.New(os.Stdout, "", 0))
	newUser := []domain.User{
		{
			Name:         "test",
			Email:        "test@email.com",
			PasswordHash: "user101",
		},
		{
			Name:         "test2",
			Email:        "test2@email.com",
			PasswordHash: "user102",
		},
	}

	for _, user := range newUser {
		createdUser, err := userService.CreateUserAccount(ctx, user, domain.UserUser)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("created user:", createdUser)
	}

	users, err := userService.GetAllUsers(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(users))
}
