package unit_test

import (
	"testing"

	"tetris-optimizer/internal"
)

func TestNewBoardInitializesWithDots(t *testing.T) {
	board := internal.NewBoard(3)

	if board.Size != 3 {
		t.Fatalf("board size = %d, want 3", board.Size)
	}

	if board.String() != "...\n...\n..." {
		t.Fatalf("board contents = %q, want %q", board.String(), "...\n...\n...")
	}
}

func TestBoardCanPlace(t *testing.T) {
	board := internal.NewBoard(4)
	tetromino := internal.Tetromino{
		Cells: []internal.Point{
			{Row: 0, Col: 0},
			{Row: 0, Col: 1},
			{Row: 1, Col: 0},
			{Row: 1, Col: 1},
		},
		Letter: 'A',
	}

	if !board.CanPlace(tetromino, 1, 1) {
		t.Fatal("CanPlace returned false for a valid placement")
	}

	if board.CanPlace(tetromino, 3, 3) {
		t.Fatal("CanPlace returned true for an out-of-bounds placement")
	}
}

func TestBoardPlaceAndRemoveTetromino(t *testing.T) {
	board := internal.NewBoard(4)
	tetromino := internal.Tetromino{
		Cells: []internal.Point{
			{Row: 0, Col: 0},
			{Row: 0, Col: 1},
			{Row: 1, Col: 0},
			{Row: 1, Col: 1},
		},
		Letter: 'B',
	}

	if err := board.PlaceTetromino(tetromino, 1, 1); err != nil {
		t.Fatalf("PlaceTetromino returned error: %v", err)
	}

	wantPlaced := "....\n.BB.\n.BB.\n...."
	if board.String() != wantPlaced {
		t.Fatalf("board after placement = %q, want %q", board.String(), wantPlaced)
	}

	board.RemoveTetromino(tetromino, 1, 1)

	wantRemoved := "....\n....\n....\n...."
	if board.String() != wantRemoved {
		t.Fatalf("board after removal = %q, want %q", board.String(), wantRemoved)
	}
}

func TestBoardPlaceTetrominoRejectsOccupiedCell(t *testing.T) {
	board := internal.NewBoard(4)
	tetromino := internal.Tetromino{
		Cells: []internal.Point{
			{Row: 0, Col: 0},
			{Row: 0, Col: 1},
			{Row: 1, Col: 0},
			{Row: 1, Col: 1},
		},
		Letter: 'C',
	}

	if err := board.PlaceTetromino(tetromino, 0, 0); err != nil {
		t.Fatalf("first PlaceTetromino returned error: %v", err)
	}

	if err := board.PlaceTetromino(tetromino, 0, 0); err == nil {
		t.Fatal("second PlaceTetromino returned nil error for occupied cells")
	}
}
