package unit_test

import (
	"testing"

	"tetris-optimizer/internal"
)

func TestSolveBoardSingleSquare(t *testing.T) {
	tetrominoes := []internal.Tetromino{
		{
			Cells: []internal.Point{
				{Row: 0, Col: 0},
				{Row: 0, Col: 1},
				{Row: 1, Col: 0},
				{Row: 1, Col: 1},
			},
			Letter: 'A',
		},
	}

	board, err := internal.SolveBoard(tetrominoes)
	if err != nil {
		t.Fatalf("SolveBoard returned error: %v", err)
	}

	if board.String() != "AA\nAA" {
		t.Fatalf("board = %q, want %q", board.String(), "AA\nAA")
	}
}

func TestSolveBoardMultipleTetrominoes(t *testing.T) {
	tetrominoes := []internal.Tetromino{
		{
			Cells: []internal.Point{
				{Row: 0, Col: 0},
				{Row: 1, Col: 0},
				{Row: 2, Col: 0},
				{Row: 3, Col: 0},
			},
			Letter: 'A',
		},
		{
			Cells: []internal.Point{
				{Row: 0, Col: 0},
				{Row: 0, Col: 1},
				{Row: 1, Col: 0},
				{Row: 1, Col: 1},
			},
			Letter: 'B',
		},
	}

	board, err := internal.SolveBoard(tetrominoes)
	if err != nil {
		t.Fatalf("SolveBoard returned error: %v", err)
	}

	if board.Size != 4 {
		t.Fatalf("board size = %d, want 4", board.Size)
	}

	if countRune(board.String(), 'A') != 4 || countRune(board.String(), 'B') != 4 {
		t.Fatalf("board does not contain the expected tetromino letters: %q", board.String())
	}
}

func countRune(text string, target rune) int {
	count := 0
	for _, char := range text {
		if char == target {
			count++
		}
	}

	return count
}
