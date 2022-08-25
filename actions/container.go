package actions

import "errors"

var (
	NotRegisteredActionErr = errors.New("not registered action")
)

// Container - manager of actions.
// Any actions with actions should be processed via container: registration the action, call of the action, etc
type Container interface {
	/*
		Register - any action should be registered before using
		Action registers with self identification, so you can assign to the action by value.

			type KickAction struct {}
			// Implementation...
			func (a KickAction) Identification() Identification {
				return "kick"
			}

			Container.Register(action)
	*/
	Register(action Action) Container

	/*
		Call action by identification and returns result to client.
		If action is doesn't exist in container it will return error

			isSuccess, err := container.Action("kick", userID)
	*/
	//Call(identification Identification, args any) (any, error)

	/*
		Get returns action or nil

			action := container.Get("kick")
	*/
	Get(identification Identification) Action

	/*
		Exists check if action is defined

			if container.Exists("kick") {
				// ...
			}
	*/
	Exists(identification Identification) bool
}
