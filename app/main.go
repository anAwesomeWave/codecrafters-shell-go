package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$ ")

		s.Scan()

		cmd := s.Text()

		fmt.Printf("%s: command not found\n", cmd)
	}
}
