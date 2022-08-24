package rbac

import (
	"ciscor/actions"
	"ciscor/config"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSimple(t *testing.T) {
	banAction := actions.Identification("ban")

	adminRole := config.RoleIdentification("admin")
	moderatorRole := config.RoleIdentification("moderator")

	cases := []struct {
		name            string
		argsForInstance func() ([]config.User, map[config.RoleIdentification]config.Role)
		expect          func(err error)
	}{
		{
			name: "Should be error for defined user with doesn't exists role",
			argsForInstance: func() ([]config.User, map[config.RoleIdentification]config.Role) {
				return []config.User{
						{
							ID:    1,
							Roles: []config.RoleIdentification{moderatorRole},
						},
					},
					map[config.RoleIdentification]config.Role{
						adminRole: {
							Name:             "admin",
							Description:      "admin",
							AvailableActions: []actions.Identification{banAction},
						},
					}
			},
			expect: func(err error) {
				assert.Equal(
					t,
					"Role [moderator] is defined for user [1] but doesn't exists in list ([roles] in the config)",
					err.Error(),
				)
			},
		},
		{
			name: "Without error",
			argsForInstance: func() ([]config.User, map[config.RoleIdentification]config.Role) {
				return []config.User{}, map[config.RoleIdentification]config.Role{}
			},
			expect: func(err error) {
				assert.Nil(t, err)
			},
		},
	}

	for _, c := range cases {
		_, err := NewSimple(c.argsForInstance())
		c.expect(err)
	}
}

func TestSimple_CanUse(t *testing.T) {
	kickAction := actions.Identification("kick")
	banAction := actions.Identification("ban")

	adminRole := config.RoleIdentification("admin")
	moderatorRole := config.RoleIdentification("moderator")

	cases := []struct {
		name            string
		argsForInstance func() ([]config.User, map[config.RoleIdentification]config.Role)
		argsForCall     func() (uint, actions.Identification)
		expect          func(result bool)
	}{
		{
			name: "Without any data",
			argsForInstance: func() ([]config.User, map[config.RoleIdentification]config.Role) {
				return []config.User{}, map[config.RoleIdentification]config.Role{}
			},
			argsForCall: func() (uint, actions.Identification) {
				return 1, kickAction
			},
			expect: func(result bool) {
				assert.False(t, result)
			},
		},
		{
			name: "User without roles",
			argsForInstance: func() ([]config.User, map[config.RoleIdentification]config.Role) {
				return []config.User{
						{
							ID:    1,
							Roles: []config.RoleIdentification{adminRole},
						},
					},
					map[config.RoleIdentification]config.Role{
						adminRole: {
							Name:             "admin",
							Description:      "admin",
							AvailableActions: []actions.Identification{banAction},
						},
					}
			},
			argsForCall: func() (uint, actions.Identification) {
				return 1, kickAction
			},
			expect: func(result bool) {
				assert.False(t, result)
			},
		},
		{
			name: "User with role but without access",
			argsForInstance: func() ([]config.User, map[config.RoleIdentification]config.Role) {
				return []config.User{
						{
							ID:    1,
							Roles: []config.RoleIdentification{moderatorRole},
						},
					},
					map[config.RoleIdentification]config.Role{
						moderatorRole: {
							Name:             "moderator",
							Description:      "moderator",
							AvailableActions: []actions.Identification{kickAction},
						},
					}
			},
			argsForCall: func() (uint, actions.Identification) {
				return 1, banAction
			},
			expect: func(result bool) {
				assert.False(t, result)
			},
		},
		{
			name: "User with role and with access",
			argsForInstance: func() ([]config.User, map[config.RoleIdentification]config.Role) {
				return []config.User{
						{
							ID:    1,
							Roles: []config.RoleIdentification{moderatorRole},
						},
					},
					map[config.RoleIdentification]config.Role{
						moderatorRole: {
							Name:             "moderator",
							Description:      "moderator",
							AvailableActions: []actions.Identification{kickAction},
						},
					}
			},
			argsForCall: func() (uint, actions.Identification) {
				return 1, kickAction
			},
			expect: func(result bool) {
				assert.True(t, result)
			},
		},
		{
			name: "User with two roles and with access to action",
			argsForInstance: func() ([]config.User, map[config.RoleIdentification]config.Role) {
				return []config.User{
						{
							ID:    1,
							Roles: []config.RoleIdentification{adminRole, moderatorRole},
						},
					},
					map[config.RoleIdentification]config.Role{
						adminRole: {
							Name:             "admin",
							Description:      "admin",
							AvailableActions: []actions.Identification{banAction},
						},
						moderatorRole: {
							Name:             "moderator",
							Description:      "moderator",
							AvailableActions: []actions.Identification{kickAction},
						},
					}
			},
			argsForCall: func() (uint, actions.Identification) {
				return 1, banAction
			},
			expect: func(result bool) {
				assert.True(t, result)
			},
		},
		{
			name: "User with two roles with same accesses and with access to action",
			argsForInstance: func() ([]config.User, map[config.RoleIdentification]config.Role) {
				return []config.User{
						{
							ID:    1,
							Roles: []config.RoleIdentification{adminRole, moderatorRole},
						},
					},
					map[config.RoleIdentification]config.Role{
						adminRole: {
							Name:             "admin",
							Description:      "admin",
							AvailableActions: []actions.Identification{banAction},
						},
						moderatorRole: {
							Name:             "moderator",
							Description:      "moderator",
							AvailableActions: []actions.Identification{banAction},
						},
					}
			},
			argsForCall: func() (uint, actions.Identification) {
				return 1, banAction
			},
			expect: func(result bool) {
				assert.True(t, result)
			},
		},
		{
			name: "User with two roles with same accesses and without access to action",
			argsForInstance: func() ([]config.User, map[config.RoleIdentification]config.Role) {
				return []config.User{
						{
							ID:    1,
							Roles: []config.RoleIdentification{adminRole, moderatorRole},
						},
					},
					map[config.RoleIdentification]config.Role{
						adminRole: {
							Name:             "admin",
							Description:      "admin",
							AvailableActions: []actions.Identification{kickAction},
						},
						moderatorRole: {
							Name:             "moderator",
							Description:      "moderator",
							AvailableActions: []actions.Identification{kickAction},
						},
					}
			},
			argsForCall: func() (uint, actions.Identification) {
				return 2, banAction
			},
			expect: func(result bool) {
				assert.False(t, result)
			},
		},
	}

	for _, c := range cases {
		instance, err := NewSimple(c.argsForInstance())
		if err != nil {
			t.Error(fmt.Sprintf("Unexpected error on instance creation: (%s) -> ", c.name), err)

			continue
		}

		c.expect(instance.CanUse(c.argsForCall()))
	}
}
