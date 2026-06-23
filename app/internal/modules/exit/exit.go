package exit

import (
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/internal/modules"
)

func NewExitModule() *ExitModule {
	return &ExitModule{}
}

type ExitModule struct {
}

var _ modules.Module = (*ExitModule)(nil)

func (e *ExitModule) Process(params string) error {
	os.Exit(0)

	return nil
}
func (e *ExitModule) GetResult() (string, error) {
	return "", nil
}

func (e *ExitModule) HandlerName() string {
	return "exit"
}

func (e *ExitModule) Help() string {
	return "exit is a shell builtin"
}
