package config

type Config struct {
	// Dynamically users access (rbac)
	Roles map[RoleIdentification]Role `yaml:"roles"`

	// Defined users with roles
	Users []User `yaml:"users"`

	// Full path to binary directory
	BasePath    string      `yaml:"-"`
	Environment Environment `yaml:"-"`
}
