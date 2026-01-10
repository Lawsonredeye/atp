package handler

import (
	"errors"
	"log"
	"net/http"
	"slices"
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
	userService service.UserServiceInterface
	logger      *log.Logger
	secret      string
}

func NewUserHandler(userService service.UserServiceInterface, logger *log.Logger, secret string) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      logger,
		secret:      secret,
	}
}

// CreateUser creates a new user
// @Summary Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body domain.RegisterUser true "User"
// @Success 201 {object} domain.User
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users [post]
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
		if errors.Is(err, pkg.ErrUserAlreadyExists) {
			return pkg.ErrorResponse(c, err, http.StatusConflict)
		}
		if errors.Is(err, pkg.ErrInvalidPasswordLength) {
			return pkg.ErrorResponse(c, err, http.StatusBadRequest)
		}
		if errors.Is(err, pkg.ErrInvalidName) {
			return pkg.ErrorResponse(c, err, http.StatusBadRequest)
		}
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	h.logger.Printf("created user with email: %s", pkg.ObfuscateDetail(createdUser.Email, "email"))
	return pkg.SuccessResponse(c, createdUser, http.StatusCreated)
}

// CreateUserAdmin creates a new admin user
// @Summary Create a new admin user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body domain.RegisterUser true "User"
// @Success 201 {object} domain.User
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/admin [post]
func (h *UserHandler) CreateUserAdmin(c echo.Context) error {
	var user domain.RegisterUser
	if err := c.Bind(&user); err != nil {
		h.logger.Println("error binding user: ", err)
		return pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}
	if err := c.Validate(&user); err != nil {
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
	createdUser, err := h.userService.CreateUserAccount(c.Request().Context(), newUser, domain.UserAdmin)
	if err != nil {
		h.logger.Println("error creating user: ", err)
		if errors.Is(err, pkg.ErrUserAlreadyExists) {
			return pkg.ErrorResponse(c, err, http.StatusConflict)
		}
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	h.logger.Printf("created user with email: %s", pkg.ObfuscateDetail(createdUser.Email, "email"))
	return pkg.SuccessResponse(c, createdUser, http.StatusCreated)
}

// Login logs in a user
// @Summary Login a user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body domain.LoginUser true "User"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/login [post]
func (h *UserHandler) Login(c echo.Context) error {
	var user domain.LoginUser
	if err := c.Bind(&user); err != nil {
		h.logger.Println("error binding user: ", err)
		return pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}
	if err := c.Validate(&user); err != nil {
		return err
	}
	loginUser, err := h.userService.Login(c.Request().Context(), user.Email, user.Password)
	if err != nil {
		h.logger.Println("error logging in user: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	h.logger.Printf("user logged in with email: %s", pkg.ObfuscateDetail(loginUser.Email, "email"))
	accessToken, err := pkg.GenerateAccessToken(loginUser.ID, domain.UserUser, accessTokenExpiry, h.secret)
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

// AdminLogin logs in a user
// @Summary Login a user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body domain.LoginUser true "User"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/login [post]
func (h *UserHandler) AdminLogin(c echo.Context) error {
	var user domain.LoginUser
	if err := c.Bind(&user); err != nil {
		h.logger.Println("error binding user: ", err)
		return pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}
	if err := c.Validate(&user); err != nil {
		return err
	}
	loginUser, err := h.userService.Login(c.Request().Context(), user.Email, user.Password)
	if err != nil {
		h.logger.Println("error logging in user: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	roles, err := h.userService.GetUserRoles(c.Request().Context(), loginUser.ID)
	if err != nil {
		return err
	}
	if slices.Contains(roles, domain.UserAdmin) == false {
		h.logger.Println("Alert user does not have role admin")
		return pkg.ErrorResponse(c, errors.New("forbidden access"), http.StatusForbidden)
	}
	h.logger.Printf("admin logged in with email: %s", pkg.ObfuscateDetail(loginUser.Email, "email"))
	accessToken, err := pkg.GenerateAccessToken(loginUser.ID, domain.UserAdmin, accessTokenExpiry, h.secret)
	if err != nil {
		h.logger.Println("error generating token: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	refreshToken, err := pkg.GenerateRefreshToken(loginUser.ID, domain.UserAdmin, refreshTokenExpiry, h.secret)
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

// UpdateUsername updates a user's username
// @Summary Update a user's username
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param user body domain.UpdateUsername true "User"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{user_id}/username [put]
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

// UpdateEmail updates a user's email
// @Summary Update a user's email
// @Tags Users
// @Accept JSON
// @Produce JSON
// @Param user_id path int true "User ID"
// @Param user body domain.UpdateEmail true "User"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{user_id}/email [put]
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

// UpdatePassword updates a user's password
// @Summary Update a user's password
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param user body domain.UpdatePassword true "User"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{user_id}/password [put]
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

// DeleteUserAccount deletes a user's account
// @Summary Delete a user's account
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{user_id}/account [delete]
func (h *UserHandler) DeleteUserAccount(c echo.Context) error {
	userId := c.Get("user_id").(int64)
	err := h.userService.DeleteUserByID(c.Request().Context(), userId)
	if err != nil {
		h.logger.Println("error deleting user account: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	h.logger.Printf("deleted user account with id: %d", userId)
	return pkg.SuccessResponse(c, nil, http.StatusOK)
}

// UserDashboard gets a user's dashboard, the total accumulated score, the number of quizzes, questions, and others
// @Summary Get a user's dashboard
// @Tags Users
// @Accept JSON
// @Produce JSON
// @Param user_id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{user_id}/dashboard [get]
func (h *UserHandler) UserDashboard(c echo.Context) error {
	userId := c.Get("user_id").(int64)
	userDashboard, err := h.userService.UserDashboard(c.Request().Context(), userId)
	if err != nil {
		h.logger.Println("error getting user dashboard: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	h.logger.Printf("got user dashboard with id: %d", userId)
	return pkg.SuccessResponse(c, userDashboard, http.StatusOK)
}
