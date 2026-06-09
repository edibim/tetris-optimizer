package main

import (
	"fmt"
	"io"
	"os"

	"tetris-optimizer/internal"
)

// main runs the CLI program and prints any user-facing error.
// Keeping error printing here keeps the rest of the code testable.
func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, internal.FormatCLIError(err))
	}
}

// run coordinates CLI input, parsing, solving, and final rendering.
func run(args []string, stdout io.Writer) error {
	if len(args) != 1 {
		return internal.ErrInvalidArgCount
	}

	data, err := internal.ReadFile(args[0])
	if err != nil {
		return err
	}

	tetrominoes, err := internal.ParseFile(data)
	if err != nil {
		return err
	}

	board, err := internal.SolveBoard(tetrominoes)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(stdout, board.String())
	return err
}
