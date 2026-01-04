package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/lawson/otterprep/domain"
	"github.com/lawson/otterprep/pkg"
)

type UserRepository struct {
	db *sql.DB
}

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user domain.User) (*domain.User, error)
	UpdateUserPassword(ctx context.Context, userId int64, newPassword string) error
	GetUserWithID(ctx context.Context, userId int64) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	DeleteUserByID(ctx context.Context, userId int64) error
	GetAllUsers(ctx context.Context) ([]domain.User, error)
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser creates a new user in the database.
func (ur *UserRepository) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	query := "INSERT INTO users (name, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	passwordHash, err := pkg.HashPassword(user.PasswordHash)
	if err != nil {
		return nil, err
	}
	user.PasswordHash = passwordHash
	result, err := ur.db.ExecContext(ctx, query, user.Name, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	user.ID, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUserPassword updates the password hash of a user in the database.
func (ur *UserRepository) UpdateUserPassword(ctx context.Context, userId int64, newPassword string) error {
	passwordHash, err := pkg.HashPassword(newPassword)
	if err != nil {
		return err
	}
	updatedAt := time.Now()
	query := fmt.Sprintf("UPDATE users SET password_hash = '%s', updated_at = '%s' WHERE id = %d", passwordHash, updatedAt, userId)
	_, err = ur.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) UpdateUserEmail(ctx context.Context, userId int64, newEmail string) error {
	updatedAt := time.Now()
	query := fmt.Sprintf("UPDATE users SET email = '%s', updated_at = '%s' WHERE id = %d", newEmail, updatedAt, userId)
	_, err := ur.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUsername updates the username of a user in the database.
func (ur *UserRepository) UpdateUsername(ctx context.Context, userId int64, newUsername string) error {
	updatedAt := time.Now()
	query := fmt.Sprintf("UPDATE users SET name = '%s', updated_at = '%s' WHERE id = %d", newUsername, updatedAt, userId)
	_, err := ur.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

// GetUserWithID gets a user from the database by ID.
func (ur *UserRepository) GetUserWithID(ctx context.Context, userId int64) (*domain.User, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE id = %d", userId)
	row := ur.db.QueryRowContext(ctx, query)
	user := domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE email = '%s'", email)
	row := ur.db.QueryRowContext(ctx, query)
	user := domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteUserByID deletes a user from the database by ID.
func (ur *UserRepository) DeleteUserByID(ctx context.Context, userId int64) error {
	query := fmt.Sprintf("DELETE FROM users WHERE id = %d", userId)
	_, err := ur.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

// GetAllUsers gets all users from the database.
func (ur *UserRepository) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	query := "SELECT id, name, email, created_at, updated_at FROM users"
	rows, err := ur.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []domain.User{}
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// CreateUserRoles creates a new user role in the database.
func (ur *UserRepository) CreateUserRoles(ctx context.Context, userId int64, role string) error {
	query := fmt.Sprintf("INSERT INTO user_roles (user_id, role) VALUES (%d, '%s')", userId, role)
	_, err := ur.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

// GetUserRoles gets all user roles from the database by user ID.
func (ur *UserRepository) GetUserRoles(ctx context.Context, userId int64) ([]string, error) {
	query := fmt.Sprintf("SELECT role FROM user_roles WHERE user_id = %d", userId)
	rows, err := ur.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	roles := []string{}
	for rows.Next() {
		role := ""
		err := rows.Scan(&role)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}
