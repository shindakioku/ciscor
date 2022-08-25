package actions

type RuntimeContainer struct {
	actions map[Identification]Action
}

func (r *RuntimeContainer) Register(action Action) Container {
	r.actions[action.Identification()] = action

	return r
}

func (r *RuntimeContainer) Exists(identification Identification) bool {
	_, exists := r.actions[identification]

	return exists
}

func (r *RuntimeContainer) Get(identification Identification) Action {
	return r.actions[identification]
}

//func (r *RuntimeContainer) Call(identification Identification, args any) (any, error) {
//	action, exists := r.actions[identification]
//	if !exists {
//		return nil, NotRegisteredActionErr
//	}
//
//	return action.Handle(args)
//}

// NewRuntimeContainer init
func NewRuntimeContainer(actions ...Action) Container {
	a := make(map[Identification]Action, len(actions))
	container := RuntimeContainer{
		actions: a,
	}

	// Register passed actions
	for _, action := range actions {
		container.Register(action)
	}

	return &container
}
