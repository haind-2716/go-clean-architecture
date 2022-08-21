package http

import (
	"go-app/internal/domain"
	"time"
)

type RolesResponse []RoleResponse

// RoleResponse is struct used for role
type RoleResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
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

// ConvertRoleToRoleResponse DTO
func ConvertRoleToRoleResponse(role *domain.Role) RoleResponse {
	return RoleResponse{
		ID:        role.ID,
		Name:      role.Name,
		Slug:      role.Slug,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

// ConvertRolesToRolesResponse DTO
func ConvertRolesToRolesResponse(roles []domain.Role) RolesResponse {
	rolesRes := make(RolesResponse, 0)

	for _, r := range roles {
		roleRes := RoleResponse{
			ID:        r.ID,
			Name:      r.Name,
			Slug:      r.Slug,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		}

		rolesRes = append(rolesRes, roleRes)
	}

	return rolesRes
}
