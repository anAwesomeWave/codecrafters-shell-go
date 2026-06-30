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

func NewModules(mods map[string]Module) *Modules {
	return &Modules{
		mods: mods,
	}
}

type Modules struct {
	mods map[string]Module // handlerName -> handler
}

func (m *Modules) GetModule(modName string) (Module, error) {
	if mod, ok := m.mods[modName]; ok {
		return mod, nil
	}

	return nil, fmt.Errorf("module %s not found: %w", modName, ErrModuleNotFound)
}
