package config

type User struct {
	ID    uint                 `yaml:"id"`
	Roles []RoleIdentification `yaml:"roles"`
}
