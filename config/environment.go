package config

type Environment int

const (
	Development Environment = iota
	Production
)

func (e Environment) ConfigFileName() (output string) {
	switch e {
	case Development:
		output = "dev"
	case Production:
		output = "prod"
	}

	return
}
