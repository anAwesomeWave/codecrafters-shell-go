package echo

import (
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/modules"
)

func NewEchoModule() *EchoModule {
	return &EchoModule{}
}

type EchoModule struct {
	Data string
}

var _ modules.Module = (*EchoModule)(nil)

func (e *EchoModule) Process(params string) error {
	e.Data = strings.Join(strings.Split(params, " ")[1:], " ")
	return nil
}
func (e *EchoModule) GetResult() (string, error) {
	return e.Data, nil
}

func (e *EchoModule) HandlerName() string {
	return "echo"
}
