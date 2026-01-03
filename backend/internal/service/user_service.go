package service

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/lawson/otterprep/domain"
	"github.com/lawson/otterprep/internal/repository"
	"github.com/lawson/otterprep/pkg"
)

type UserServiceInterface interface {
	CreateUserAccount(ctx context.Context, user domain.User) (*domain.User, error)
	GetUserWithID(ctx context.Context, userId int64) (*domain.User, error)
	UpdateUsername(ctx context.Context, userId int64, newUsername string) error
	UpdatePassword(ctx context.Context, userId int64, newPassword string) error
	GetAllUsers(ctx context.Context) ([]domain.User, error)
	DeleteUserByID(ctx context.Context, userId int64) error
	Login(ctx context.Context, email string, password string) (*domain.User, error)
}

type userService struct {
	userRepo repository.UserRepository
	logger   *log.Logger
}

func NewUserService(userRepo repository.UserRepository, logger *log.Logger) *userService {
	return &userService{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (s *userService) CreateUserAccount(ctx context.Context, user domain.User, role string) (*domain.User, error) {
	if user.Name == "" {
		s.logger.Println("error creating user: ", pkg.ErrInvalidName)
		return nil, pkg.ErrInvalidName
	}
	if user.Email == "" {
		s.logger.Println("error creating user: ", pkg.ErrInvalidEmail)
		return nil, pkg.ErrInvalidEmail
	}
	if user.PasswordHash == "" { // the len of the user password should be >= 6
		s.logger.Println("error creating user: ", pkg.ErrInvalidPasswordHash)
		return nil, pkg.ErrInvalidPasswordHash
	}
	createdUser, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		s.logger.Println("error creating user: ", err)
		return nil, err
	}
	if role == "" {
		s.logger.Println("error creating user: ", pkg.ErrInvalidRole, "assigning a user role")
		role = domain.UserUser
	}
	if err := s.userRepo.CreateUserRoles(ctx, createdUser.ID, strings.ToLower(role)); err != nil {
		s.logger.Println("error creating user roles: ", err)
		return nil, err
	}
	createdUser.PasswordHash = ""
	s.logger.Println("created user: ", createdUser)
	return createdUser, nil
}

func (s *userService) GetUserWithID(ctx context.Context, userId int64) (*domain.User, error) {
	if userId == 0 {
		return nil, pkg.ErrInvalidUserID
	}
	s.logger.Println("getting user with id: ", userId)
	user, err := s.userRepo.GetUserWithID(ctx, userId)
	if err != nil {
		s.logger.Println("error getting user: ", err)
		return nil, err
	}
	user.PasswordHash = ""
	s.logger.Println("getting user with id: ", user)
	return user, nil
}

func (s *userService) UpdateUsername(ctx context.Context, userId int64, newUsername string) error {
	if userId == 0 {
		s.logger.Println("error updating username: ", pkg.ErrInvalidUserID)
		return pkg.ErrInvalidUserID
	}
	if newUsername == "" {
		s.logger.Println("error updating username: ", pkg.ErrInvalidName)
		return pkg.ErrInvalidName
	}
	err := s.userRepo.UpdateUsername(ctx, userId, newUsername)
	if err != nil {
		s.logger.Println("error updating username: ", err)
		return err
	}
	s.logger.Println("updated username: ", newUsername)
	return nil
}

func (s *userService) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	s.logger.Println("getting all users")
	users, err := s.userRepo.GetAllUsers(ctx)
	if err != nil {
		s.logger.Println("error getting all users: ", err)
		return nil, err
	}
	for i := range users {
		users[i].PasswordHash = ""
	}
	s.logger.Println("getting all users: ", users)
	return users, nil
}

func (s *userService) DeleteUserByID(ctx context.Context, userId int64) error {
	if userId == 0 {
		s.logger.Println("error deleting user: ", pkg.ErrInvalidUserID)
		return pkg.ErrInvalidUserID
	}
	err := s.userRepo.DeleteUserByID(ctx, userId)
	if err != nil {
		s.logger.Println("error deleting user: ", err)
		return err
	}
	s.logger.Println("deleted user: ", userId)
	return nil
}

func (s *userService) UpdatePassword(ctx context.Context, userId int64, newPassword string) error {
	if userId == 0 {
		s.logger.Println("error updating password: ", pkg.ErrInvalidUserID)
		return pkg.ErrInvalidUserID
	}
	if newPassword == "" || len(newPassword) < 6 {
		s.logger.Println("error updating password: ", pkg.ErrInvalidPasswordHash)
		return pkg.ErrInvalidPasswordHash
	}
	err := s.userRepo.UpdateUserPassword(ctx, userId, newPassword)
	if err != nil {
		s.logger.Println("error updating password: ", err)
		if errors.Is(err, pkg.ErrInvalidPasswordHash) {
			return pkg.ErrInvalidPasswordHash
		}
		return pkg.ErrInternalServerError
	}
	s.logger.Println("updated password: ", newPassword)
	return nil
}

func (s *userService) Login(ctx context.Context, email string, password string) (*domain.User, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		s.logger.Println("error getting user: ", err)
		return nil, err
	}
	if user == nil {
		s.logger.Println("user not found")
		return nil, pkg.ErrUserNotFound
	}
	if !pkg.CheckPasswordHash(password, user.PasswordHash) {
		s.logger.Println("password does not match")
		return nil, pkg.ErrInvalidPasswordHash
	}
	s.logger.Println("user logged in: ", user)
	return user, nil
}
