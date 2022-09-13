package runner

import (
	"github.com/shindakioku/ciscor/actions"
)

type SimpleSyncRunner struct {
}

func (s SimpleSyncRunner) Run(action actions.Action, args any) (any, error) {
	return action.Handle(args)
}

func NewSimpleSyncRunner() SyncRunner {
	return &SimpleSyncRunner{}
}
