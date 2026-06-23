package typemod

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/internal/modules"
)

func NewTypeModule() *TypeModule {
	return &TypeModule{}
}

func (t *TypeModule) SetModules(modules map[string]modules.ModuleInfo) {
	t.modules = modules
}

type TypeModule struct {
	modules map[string]modules.ModuleInfo
	Data    string
}

var _ modules.Module = (*TypeModule)(nil)

func (t *TypeModule) Process(params string) error {
	args := strings.Split(params, " ")[1:]

	anses := make([]string, 0, len(args))
	for _, a := range args {
		anses = append(anses, t.processSingleArg(a))
	}

	t.Data = strings.Join(anses, "\n")

	return nil
}

func (t *TypeModule) processSingleArg(arg string) string {
	if mod, ok := t.modules[arg]; ok {
		return mod.Help()
	}

	return fmt.Sprintf("%s: not found", arg)
}

func (t *TypeModule) GetResult() (string, error) {
	return t.Data, nil
}

func (t *TypeModule) HandlerName() string {
	return "type"
}

func (t *TypeModule) Help() string {
	return "type is a shell builtin"
}
