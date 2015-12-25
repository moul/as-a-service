package moul

var ActionsMap map[string]Action

type Action func([]string) (interface{}, error)

func RegisterAction(name string, action Action) {
	if ActionsMap == nil {
		ActionsMap = make(map[string]Action)
	}
	ActionsMap[name] = action
}

func Actions() map[string]Action {
	return ActionsMap
}
