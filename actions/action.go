package actions

type Action interface {
	// Identification - need for understanding how to work with action
	// For example: if you need to call [kick] action, we will find the action by [Identification]
	// Meaningful name should be a preferred
	Identification() Identification

	/*
		Handle - call action.
		Unfortunately, we can't provide generics for type hint for those params. It's not a possible now.
		Well, you should to cast (if you want to do this). Describe any action args and response by struct,
		it will help to you which casts.

		Example with casting:

			type KickActionRequest struct {
				UserID uint
			}
			type KickActionResult struct {
				Success bool
				Err error
			}
			type KickAction struct {}
			func (a KickAction) Handle(args any) (any, error) {
				request, ok := args.(KickActionRequest)
				if !ok {
					return nil, errors.New("You should provide [KickActionRequest] as arg")
				}

				// request.UserID

				return KickActionResult{Success: true, err: nil}
			}
			// Implementation...

			result, err := KickAction.Handle(KickActionRequest{UserID: 1})
			if result, ok := result.(KickActionResult); !ok {
				// ???
			}
	*/
	Handle(args any) (any, error)
}
