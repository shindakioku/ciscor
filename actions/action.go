package actions

type Identification string

type Action interface {
	// Key for identification
	Key() Identification

	Handle() error
}
