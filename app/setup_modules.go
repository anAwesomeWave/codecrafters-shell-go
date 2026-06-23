package main

import (
	"errors"
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/app/modules"
	"github.com/codecrafters-io/shell-starter-go/app/modules/exit"
)

var (
	ErrModuleNotFound = errors.New("module not found")
)

func NewModules() *Modules {
	exitMod := exit.NewExitModule()

	return &Modules{
		handlers: map[string]modules.Module{
			exitMod.HandlerName(): exitMod,
		},
	}
}

type Modules struct {
	handlers map[string]modules.Module // handlerName -> handler
}

func (m *Modules) GetModule(modName string) (modules.Module, error) {
	if mod, ok := m.handlers[modName]; ok {
		return mod, nil
	}

	return nil, fmt.Errorf("module %s not found: %w", modName, ErrModuleNotFound)
}
