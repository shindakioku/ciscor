package rbac

import "github.com/shindakioku/ciscor/actions"

type Rbac interface {
	// CanUse - can a user use action
	CanUse(userID uint, action actions.Identification) bool
}
