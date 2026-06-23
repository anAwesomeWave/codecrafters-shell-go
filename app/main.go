package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"errors"
)

func run(s *bufio.Scanner, handlers *Modules) {
	for {
		fmt.Print("$ ")

		s.Scan()

		cmd := s.Text()

		ans, err := process(cmd, handlers)
		if err != nil {
			panic(err)
		}

		fmt.Println(ans)

	}

}

func process(cmd string, handlers *Modules) (string, error) {
	commandName := strings.Split(cmd, " ")[0]

	mod, err := handlers.GetModule(commandName)
	if err != nil {
		if errors.Is(err, ErrModuleNotFound) {
			return fmt.Sprintf("%s: command not found", cmd), nil
		}

		return "", nil
	}

	if err := mod.Process(cmd); err != nil {
		return "", fmt.Errorf("couldn't process mod=%s cmd=%s: %w", mod.HandlerName(), cmd, err)
	}

	ans, err := mod.GetResult()
	if err != nil {
		return "", fmt.Errorf("couldn't get result mod=%s cmd=%s: %w", mod.HandlerName(), cmd, err)
	}

	return ans, nil
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	handlers := NewModules()

	run(s, handlers)
}
