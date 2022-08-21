package http

import (
	"go-app/internal/domain"
	"time"
)

type UsersResponse []UserResponse

// UserResponse is struct used for user
type UserResponse struct {
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

// StatusResponse is struc when success
type StatusResponse struct {
	Status bool `json:"status"`
}

// ConvertUserToUserResponse DTO
func ConvertUserToUserResponse(user *domain.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		RoleID:    user.RoleID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// ConvertUsersToUsersResponse DTO
func ConvertUsersToUsersResponse(users []domain.User) UsersResponse {
	usersRes := make(UsersResponse, 0)

	for _, u := range users {
		userRes := UserResponse{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			RoleID:    u.RoleID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		}

		usersRes = append(usersRes, userRes)
	}

	return usersRes
}
