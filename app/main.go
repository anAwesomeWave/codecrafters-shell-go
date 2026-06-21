package main

import (
	"fmt"
)

func main() {
	// TODO: Uncomment the code below to pass the first stage
	fmt.Print("$ ")

	var cmd string

	fmt.Scan(&cmd)

	fmt.Printf("%s: command not found", cmd)
}
