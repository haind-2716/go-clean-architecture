package http

import (
	"go-app/internal/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

// AuthHandler represent the httphandler
type AuthHandler struct {
	Usecase domain.AuthUsecase
}

type errorResponse struct {
	Message string `json:"message"`
}

// NewAuthHandler will initialize the Auth endpoint
func NewAuthHandler(g *echo.Group, uc domain.AuthUsecase) {
	handler := &AuthHandler{
		Usecase: uc,
	}

	g.POST("/login", handler.Login)
	g.POST("/register", handler.Register)
}

// Login for user
func (hl *AuthHandler) Login(c echo.Context) error {
	user := &domain.User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &errorResponse{Message: err.Error()})
	}

	ctx := c.Request().Context()
	UserLogin, tokenStr, err := hl.Usecase.Login(ctx, user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, UserLoginResponse{
		UserID:      UserLogin.ID,
		Email:       UserLogin.Email,
		RoleID:      UserLogin.RoleID,
		AccessToken: tokenStr,
	})
}

// Register for user
func (hl *AuthHandler) Register(c echo.Context) error {
	user := &domain.User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &errorResponse{Message: err.Error()})
	}

	ctx := c.Request().Context()
	userRegister, err := hl.Usecase.Register(ctx, user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, userRegister)
}