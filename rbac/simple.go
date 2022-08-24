package rbac

import (
	"github.com/shindakioku/ciscor/actions"
	"github.com/shindakioku/ciscor/config"
	"errors"
	"fmt"
)

type User struct {
	AvailableActions map[actions.Identification]bool
	Roles            map[config.RoleIdentification]config.Role
}

type Simple struct {
	/**
	Users with available actions for quick access
	Key of first map - user id

	Example:
		{1: {
			available_actions: {kick: true, ban: true},
			roles: {admin: {...}, editor: {...}}
		}}
	*/
	users map[uint]User
}

func (s Simple) CanUse(userID uint, action actions.Identification) bool {
	userData, exists := s.users[userID]
	if !exists {
		return false
	}

	_, canUse := userData.AvailableActions[action]

	return canUse
}

// NewSimple - simple implementer of rbac
// Before init we should optimize data for quick access in lifecycle
func NewSimple(users []config.User, roles map[config.RoleIdentification]config.Role) (Rbac, error) {
	u := make(map[uint]User, len(users))
	for _, user := range users {
		newUser := User{
			AvailableActions: make(map[actions.Identification]bool),
			Roles:            make(map[config.RoleIdentification]config.Role),
		}

		// Convert user role to defined role in config
		for _, userRole := range user.Roles {
			role, exists := roles[userRole]
			// Cannot describe role for the user which doesn't define in config
			// Maybe we should skip it with logging only?
			if !exists {
				return Simple{}, errors.New(fmt.Sprintf(
					"Role [%s] is defined for user [%d] but doesn't exists in list ([roles] in the config)",
					userRole, user.ID,
				))
			}

			newUser.Roles[userRole] = role
			for _, availableAction := range role.AvailableActions {
				newUser.AvailableActions[availableAction] = true
			}
		}

		u[user.ID] = newUser
	}

	return Simple{
		users: u,
	}, nil
}
