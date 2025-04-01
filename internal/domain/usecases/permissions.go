package usecases

import "github.com/huuloc2026/go-social/internal/domain/entities"

type Permission string

const (
	PermissionUserRead   Permission = "user:read"
	PermissionUserWrite  Permission = "user:write"
	PermissionUserDelete Permission = "user:delete"
	PermissionAdmin      Permission = "admin"
)

var RolePermissions = map[entities.Role][]Permission{
	entities.RoleUser: {
		PermissionUserRead,
	},
	entities.RoleAdmin: {
		PermissionUserRead,
		PermissionUserWrite,
		PermissionUserDelete,
		PermissionAdmin,
	},
}

type PermissionChecker interface {
	HasPermission(user *entities.User, permission Permission) bool
}

type permissionChecker struct{}

func NewPermissionChecker() PermissionChecker {
	return &permissionChecker{}
}

func (c *permissionChecker) HasPermission(user *entities.User, permission Permission) bool {
	if user == nil {
		return false
	}

	permissions, ok := RolePermissions[user.Role]
	if !ok {
		return false
	}

	for _, p := range permissions {
		if p == permission || p == PermissionAdmin {
			return true
		}
	}

	return false
}
