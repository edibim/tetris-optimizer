package main

import (
	"fmt"
	"os"

	"tetris-optimizer/internal"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Println("ERROR")
	}
}

func run(args []string) error {
	if len(args) != 1 {
		return internal.ErrInvalidInput
	}

	_, err := internal.ReadFile(args[0])
	if err != nil {
		return err
	}

	return internal.ErrNotImplemented
}
