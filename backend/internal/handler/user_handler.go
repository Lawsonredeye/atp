package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lawson/otterprep/domain"
	"github.com/lawson/otterprep/internal/service"
	"github.com/lawson/otterprep/pkg"
)

var (
	accessTokenExpiry  = 15 * time.Minute
	refreshTokenExpiry = 7 * 24 * time.Hour
)

type UserHandler struct {
	userService       service.UserServiceInterface
	logger            *log.Logger
	accessTokenExpiry time.Duration
	secret            string
}

func NewUserHandler(userService service.UserServiceInterface, logger *log.Logger, secret string) *UserHandler {
	return &UserHandler{
		userService:       userService,
		logger:            logger,
		accessTokenExpiry: accessTokenExpiry,
		secret:            secret,
	}
}

// CreateUser creates a new user
func (h *UserHandler) CreateUser(c echo.Context) error {
	var user domain.RegisterUser
	err := c.Bind(&user)
	if err != nil {
		return err
	}
	now := time.Now()
	newUser := domain.User{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.Password,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	h.logger.Printf("registering new user with email: %s", pkg.ObfuscateDetail(user.Email, "email"))
	createdUser, err := h.userService.CreateUserAccount(c.Request().Context(), newUser, domain.UserUser)
	if err != nil {
		h.logger.Println("error creating user: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	h.logger.Printf("created user with email: %s", pkg.ObfuscateDetail(createdUser.Email, "email"))
	return pkg.SuccessResponse(c, createdUser, http.StatusCreated)
}

func (h *UserHandler) CreateUserAdmin(c echo.Context) error {
	var user domain.RegisterUser
	err := c.Bind(&user)
	if err != nil {
		h.logger.Println("error binding user: ", err)
		return pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}
	now := time.Now()
	newUser := domain.User{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.Password,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	h.logger.Printf("registering new user with email: %s", pkg.ObfuscateDetail(user.Email, "email"))
	createdUser, err := h.userService.CreateUserAccount(c.Request().Context(), newUser, domain.UserAdmin)
	if err != nil {
		h.logger.Println("error creating user: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	h.logger.Printf("created user with email: %s", pkg.ObfuscateDetail(createdUser.Email, "email"))
	return pkg.SuccessResponse(c, createdUser, http.StatusCreated)
}

func (h *UserHandler) Login(c echo.Context) error {
	var user domain.LoginUser
	err := c.Bind(&user)
	if err != nil {
		h.logger.Println("error binding user: ", err)
		return pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}
	loginUser, err := h.userService.Login(c.Request().Context(), user.Email, user.Password)
	if err != nil {
		h.logger.Println("error logging in user: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	h.logger.Printf("user logged in with email: %s", pkg.ObfuscateDetail(loginUser.Email, "email"))
	accessToken, err := pkg.GenerateAccessToken(loginUser.ID, domain.UserUser, h.accessTokenExpiry, h.secret)
	if err != nil {
		h.logger.Println("error generating token: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	refreshToken, err := pkg.GenerateRefreshToken(loginUser.ID, domain.UserUser, refreshTokenExpiry, h.secret)
	if err != nil {
		h.logger.Println("error generating token: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	data := map[string]interface{}{
		"user":          loginUser,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}
	return pkg.SuccessResponse(c, data, http.StatusOK)
}

func (h *UserHandler) UpdateUsername(c echo.Context) error {
	userID := c.Get("user_id").(int64)
	var user domain.UpdateUsername
	err := c.Bind(&user)
	if err != nil {
		h.logger.Println("error binding user: ", err)
		return pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}
	err = h.userService.UpdateUsername(c.Request().Context(), userID, user.NewUsername)
	if err != nil {
		h.logger.Println("error updating username: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	h.logger.Printf("updated username with id: %d", userID)
	return pkg.SuccessResponse(c, nil, http.StatusOK)
}

func (h *UserHandler) UpdateEmail(c echo.Context) error {
	userID := c.Get("user_id").(int64)
	var user domain.UpdateEmail
	err := c.Bind(&user)
	if err != nil {
		h.logger.Println("error binding user: ", err)
		return pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}
	err = h.userService.UpdateEmail(c.Request().Context(), userID, user.NewEmail)
	if err != nil {
		h.logger.Println("error updating email: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	h.logger.Printf("updated email with id: %d", userID)
	return pkg.SuccessResponse(c, nil, http.StatusOK)
}

func (h *UserHandler) UpdatePassword(c echo.Context) error {
	userID := c.Get("user_id").(int64)
	var user domain.UpdatePassword
	err := c.Bind(&user)
	if err != nil {
		h.logger.Println("error binding user: ", err)
		return pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}
	err = h.userService.UpdatePassword(c.Request().Context(), userID, user.NewPassword)
	if err != nil {
		h.logger.Println("error updating password: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	h.logger.Printf("updated password with id: %d", userID)
	return pkg.SuccessResponse(c, nil, http.StatusOK)
}
