package main

import (
	"github.com/codecrafters-io/shell-starter-go/app/internal/modules"
	"github.com/codecrafters-io/shell-starter-go/app/internal/modules/echo"
	"github.com/codecrafters-io/shell-starter-go/app/internal/modules/exit"
	"github.com/codecrafters-io/shell-starter-go/app/internal/modules/typemod"
)

func NewModules() *Modules {
	exitMod := exit.NewExitModule()
	echoMod := echo.NewEchoModule()
	typeMod := typemod.NewTypeModule()

	handlers := map[string]modules.Module{
		exitMod.HandlerName(): exitMod,
		echoMod.HandlerName(): echoMod,
		typeMod.HandlerName(): typeMod,
	}

	descHandlers := make(map[string]modules.ModuleInfo)
	for k, v := range handlers {
		descHandlers[k] = v
	}

	typeMod.SetModules(descHandlers)

	return &Modules{
		handlers: handlers,
	}
}

type Modules struct {
	handlers map[string]modules.Module // handlerName -> handler
}

func (m *Modules) GetModule(modName string) (modules.Module, error) {
	return modules.GetModule(modName, m.handlers)
}
