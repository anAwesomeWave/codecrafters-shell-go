package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/internal"
)

func run(processor *internal.Processor, s *bufio.Scanner) {
	for {
		fmt.Print("$ ")

		s.Scan()

		cmd := s.Text()

		ans, err := processor.Process(cmd)
		if err != nil {
			panic(err)
		}

		fmt.Print(ans)
	}

}
func main() {
	s := bufio.NewScanner(os.Stdin)
	handlers := SetupModules()

	processor := internal.NewProcessor(handlers)

	run(processor, s)
}
