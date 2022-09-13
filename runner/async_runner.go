package runner

import (
	"context"
)

// ASyncRunner - async executor of actions
type ASyncRunner interface {
	Add(job AsyncJob)

	Run(ctx context.Context) error
}
