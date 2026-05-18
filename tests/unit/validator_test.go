package unit_test

import (
	"testing"

	"tetris-optimizer/internal"
)

func TestValidateTetrominoValidSquare(t *testing.T) {
	tetromino := internal.Tetromino{
		Cells: []internal.Point{
			{Row: 0, Col: 0},
			{Row: 0, Col: 1},
			{Row: 1, Col: 0},
			{Row: 1, Col: 1},
		},
		Letter: 'A',
	}

	err := internal.ValidateTetromino(tetromino)
	if err != nil {
		t.Fatalf("ValidateTetromino returned error: %v", err)
	}
}

func TestValidateTetrominoRejectsWrongBlockCount(t *testing.T) {
	tetromino := internal.Tetromino{
		Cells: []internal.Point{
			{Row: 0, Col: 0},
			{Row: 0, Col: 1},
			{Row: 1, Col: 0},
		},
		Letter: 'A',
	}

	err := internal.ValidateTetromino(tetromino)
	if err == nil {
		t.Fatal("ValidateTetromino returned nil error for wrong block count")
	}
}

func TestValidateTetrominoRejectsDisconnectedShape(t *testing.T) {
	tetromino := internal.Tetromino{
		Cells: []internal.Point{
			{Row: 0, Col: 0},
			{Row: 0, Col: 1},
			{Row: 3, Col: 2},
			{Row: 3, Col: 3},
		},
		Letter: 'A',
	}

	err := internal.ValidateTetromino(tetromino)
	if err == nil {
		t.Fatal("ValidateTetromino returned nil error for disconnected shape")
	}
}

func TestNormalizeTetrominoMovesShapeToTopLeft(t *testing.T) {
	tetromino := internal.Tetromino{
		Cells: []internal.Point{
			{Row: 2, Col: 2},
			{Row: 2, Col: 3},
			{Row: 3, Col: 2},
			{Row: 3, Col: 3},
		},
		Letter: 'B',
	}

	got := internal.NormalizeTetromino(tetromino)

	want := []internal.Point{
		{Row: 0, Col: 0},
		{Row: 0, Col: 1},
		{Row: 1, Col: 0},
		{Row: 1, Col: 1},
	}

	if got.Letter != tetromino.Letter {
		t.Fatalf("NormalizeTetromino changed letter from %q to %q", tetromino.Letter, got.Letter)
	}

	for index := range want {
		if got.Cells[index] != want[index] {
			t.Fatalf("cell %d = %+v, want %+v", index, got.Cells[index], want[index])
		}
	}
}
