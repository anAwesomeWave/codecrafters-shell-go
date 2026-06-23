package modules

import (
	"errors"
	"fmt"
)

var (
	ErrModuleNotFound = errors.New("module not found")
)

type Module interface {
	ModuleProcessor
	ModuleInfo
}

// Process retrieves and runs input data from user
// GetResult returns result of processed module
type ModuleProcessor interface {
	Process(params string) error
	GetResult() (string, error)

	HandlerName() string
}

type ModuleInfo interface {
	Help() string
}

func GetModule(modName string, allMods map[string]Module) (Module, error) {
	if mod, ok := allMods[modName]; ok {
		return mod, nil
	}

	return nil, fmt.Errorf("module %s not found: %w", modName, ErrModuleNotFound)
}
