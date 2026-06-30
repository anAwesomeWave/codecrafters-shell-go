package typemod

import (
	"errors"
	"fmt"
	"strings"

	dirscanner "github.com/codecrafters-io/shell-starter-go/app/internal/dir_scanner"
	"github.com/codecrafters-io/shell-starter-go/app/internal/envs"
	"github.com/codecrafters-io/shell-starter-go/app/internal/modules"
	my_slices "github.com/codecrafters-io/shell-starter-go/app/internal/utils/slices"
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

	anses, err := my_slices.MapE(args, func(arg string) (string, error) {
		return t.processSingleArg(arg)
	})
	if err != nil {
		return err // todo: desc
	}

	t.Data = strings.Join(anses, "\n")

	return nil
}

func (t *TypeModule) processSingleArg(arg string) (string, error) {
	if mod, ok := t.modules[arg]; ok {
		return mod.Help(), nil
	}

	path, err := dirscanner.ExecLookup(arg, envs.PathEnvProcess(envs.PATH_ENV))
	if err != nil {
		if errors.Is(err, dirscanner.ErrExecNotFound) {
			return fmt.Sprintf("%s: not found", arg), nil
		}

		return "", err
	}

	return arg + " is " + path, nil
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
