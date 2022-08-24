package config

import "github.com/shindakioku/ciscor/actions"

type RoleIdentification string

type Role struct {
	Name             string                   `yaml:"name"`
	Description      string                   `yaml:"description"`
	AvailableActions []actions.Identification `yaml:"available_actions"`
}
