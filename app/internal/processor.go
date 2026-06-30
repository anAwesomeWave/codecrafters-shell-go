package internal

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	dirscanner "github.com/codecrafters-io/shell-starter-go/app/internal/dir_scanner"
	"github.com/codecrafters-io/shell-starter-go/app/internal/envs"
	"github.com/codecrafters-io/shell-starter-go/app/internal/modules"
)

func NewProcessor(modProcessor *modules.Modules) *Processor {
	return &Processor{
		m: modProcessor,
	}
}

type Processor struct {
	m *modules.Modules
}

func (p *Processor) processBuiltin(commandName string, cmd string) (string, error) {
	mod, err := p.m.GetModule(commandName)
	if err != nil {
		return "", err
	}

	if err := mod.Process(cmd); err != nil {
		return "", fmt.Errorf("couldn't process mod=%s cmd=%s: %w", mod.HandlerName(), cmd, err)
	}

	ans, err := mod.GetResult()
	if err != nil {
		return "", fmt.Errorf("couldn't get result mod=%s cmd=%s: %w", mod.HandlerName(), cmd, err)
	}

	return ans + "\n", nil
}

func (p *Processor) processExec(commandName string, cmd string) (string, error) {
	_, err := dirscanner.ExecLookup(commandName, envs.PathEnvProcess(envs.PATH_ENV))
	if err != nil {
		return "", fmt.Errorf("couldn't lookup exec file: %w", err)
	}

	execCmd := exec.Command(commandName, strings.Split(cmd, " ")[1:]...)

	out, err := execCmd.Output()
	if err != nil {
		return err.Error(), nil // , fmt.Errorf("couldn't run exec: %w", err)
	}

	return string(out), nil
}

func (p *Processor) Process(cmd string) (string, error) {
	commandName := strings.Split(cmd, " ")[0]

	ans, err := p.processBuiltin(commandName, cmd)
	if err != nil && !errors.Is(err, modules.ErrModuleNotFound) {
		return "", fmt.Errorf("couldn't run shell builtin: %w", err)
	}

	if err == nil {
		return ans, nil
	}

	ans, err = p.processExec(commandName, cmd)
	if err != nil {
		if errors.Is(err, dirscanner.ErrExecNotFound) {
			return fmt.Sprintf("%s: command not found\n", cmd), nil
		}

		return "", fmt.Errorf("couldn't process exec file: %w", err)
	}

	return ans, nil
}
