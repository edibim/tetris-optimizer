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
