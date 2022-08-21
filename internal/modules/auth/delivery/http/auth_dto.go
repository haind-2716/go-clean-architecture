package http

import (
	"go-app/internal/domain"
	"time"
)

// UserLoginResponse is struct used for log in
type UserLoginResponse struct {
	ID     uint         `json:"id"`
	Name   string       `json:"name"`
	Email  string       `json:"email"`
	RoleID int          `json:"role_id"`
	Auth   AuthResponse `json:"auth"`
}

// AuthResponse is struct used for token
type AuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}

// UserRegisterResponse is struct used for register
type UserRegisterResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	RoleID    int       `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ErrorResponse is struc when error
type ErrorResponse struct {
	Message string `json:"message"`
}

// ConvertUserToUserLoginResponse DTO
func ConvertUserToUserLoginResponse(claims domain.Claims, tokenStr string) UserLoginResponse {
	return UserLoginResponse{
		ID:     claims.ID,
		Name:   claims.Name,
		Email:  claims.Email,
		RoleID: claims.RoleID,
		Auth: AuthResponse{
			AccessToken: tokenStr,
			ExpiresAt:   claims.ExpiresAt,
		},
	}
}

// ConvertUserToUserRegisterResponse DTO
func ConvertUserToUserRegisterResponse(user domain.User) UserRegisterResponse {
	return UserRegisterResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		RoleID:    user.RoleID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
