package unit_test

import (
	"path/filepath"
	"testing"

	"tetris-optimizer/internal"
)

func TestReadFileSuccess(t *testing.T) {
	filePath := filepath.Join("..", "testdata", "valid_single_o.txt")

	data, err := internal.ReadFile(filePath)
	if err != nil {
		t.Fatalf("ReadFile returned error: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("ReadFile returned empty data")
	}
}

func TestReadFileMissingFile(t *testing.T) {
	filePath := filepath.Join("..", "testdata", "missing.txt")

	_, err := internal.ReadFile(filePath)
	if err == nil {
		t.Fatal("ReadFile returned nil error for missing file")
	}
}

func TestParseFileSingleTetromino(t *testing.T) {
	data := []byte("##..\n##..\n....\n....\n")

	tetrominoes, err := internal.ParseFile(data)
	if err != nil {
		t.Fatalf("ParseFile returned error: %v", err)
	}

	if len(tetrominoes) != 1 {
		t.Fatalf("ParseFile returned %d tetrominoes, want 1", len(tetrominoes))
	}

	got := tetrominoes[0]
	if got.Letter != 'A' {
		t.Fatalf("tetromino letter = %q, want %q", got.Letter, 'A')
	}

	if len(got.Cells) != 4 {
		t.Fatalf("tetromino has %d cells, want 4", len(got.Cells))
	}
}

func TestParseFileMultipleTetrominoes(t *testing.T) {
	data := []byte(
		"##..\n##..\n....\n....\n\n" +
			"...#\n...#\n...#\n...#\n",
	)

	tetrominoes, err := internal.ParseFile(data)
	if err != nil {
		t.Fatalf("ParseFile returned error: %v", err)
	}

	if len(tetrominoes) != 2 {
		t.Fatalf("ParseFile returned %d tetrominoes, want 2", len(tetrominoes))
	}

	if tetrominoes[0].Letter != 'A' || tetrominoes[1].Letter != 'B' {
		t.Fatal("ParseFile did not assign letters in input order")
	}
}

func TestParseFileEmptyInput(t *testing.T) {
	_, err := internal.ParseFile(nil)
	if err == nil {
		t.Fatal("ParseFile returned nil error for empty input")
	}
}

func TestParseFileRejectsExtraBlankLine(t *testing.T) {
	data := []byte(
		"##..\n##..\n....\n....\n\n\n" +
			"...#\n...#\n...#\n...#\n",
	)

	_, err := internal.ParseFile(data)
	if err == nil {
		t.Fatal("ParseFile returned nil error for extra blank line")
	}
}

func TestParseFileRejectsShortLine(t *testing.T) {
	data := []byte("##.\n##..\n....\n....\n")

	_, err := internal.ParseFile(data)
	if err == nil {
		t.Fatal("ParseFile returned nil error for short line")
	}
}

func TestParseFileRejectsInvalidCharacter(t *testing.T) {
	data := []byte("A#..\n##..\n....\n....\n")

	_, err := internal.ParseFile(data)
	if err == nil {
		t.Fatal("ParseFile returned nil error for invalid character")
	}
}
