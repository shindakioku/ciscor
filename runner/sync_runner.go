package runner

import "github.com/shindakioku/ciscor/actions"

// SyncRunner - synchronous executor of actions
type SyncRunner interface {
	Run(action actions.Action, args any) (any, error)
}
