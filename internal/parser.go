package internal

import (
	"fmt"
	"os"
	"strings"
)

// ReadFile loads the raw bytes from the CLI-provided input path.
func ReadFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrReadFile, err)
	}

	return data, nil
}

// ParseFile validates the text format and returns normalized tetrominoes.
func ParseFile(data []byte) ([]Tetromino, error) {
	content := normalizeNewlines(string(data))
	if content == "" {
		return nil, ErrEmptyInput
	}

	lines := strings.Split(content, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	tetrominoes := make([]Tetromino, 0)
	index := 0

	for index < len(lines) {
		if lines[index] == "" {
			return nil, fmt.Errorf("%w: line %d", ErrUnexpectedBlankLine, index+1)
		}

		block, nextIndex, err := parseTetrominoBlock(lines, index, len(tetrominoes))
		if err != nil {
			return nil, err
		}

		if err := ValidateTetromino(block); err != nil {
			return nil, err
		}

		block = NormalizeTetromino(block)

		tetrominoes = append(tetrominoes, block)
		index = nextIndex

		if index == len(lines) {
			break
		}

		if lines[index] != "" {
			return nil, fmt.Errorf("%w: line %d", ErrUnexpectedBlankLine, index+1)
		}

		index++
		if index == len(lines) {
			return nil, fmt.Errorf("%w: after tetromino %c", ErrTrailingSeparator, rune('A'+len(tetrominoes)-1))
		}
	}

	if len(tetrominoes) == 0 {
		return nil, ErrEmptyInput
	}

	return tetrominoes, nil
}

func normalizeNewlines(content string) string {
	return strings.ReplaceAll(content, "\r\n", "\n")
}

func parseTetrominoBlock(lines []string, startIndex int, tetrominoIndex int) (Tetromino, int, error) {
	letter := rune('A' + tetrominoIndex)
	if startIndex+4 > len(lines) {
		return Tetromino{}, 0, fmt.Errorf("%w: tetromino %c starts at line %d", ErrIncompleteBlock, letter, startIndex+1)
	}

	cells := make([]Point, 0, 4)
	for rowOffset := 0; rowOffset < 4; rowOffset++ {
		line := lines[startIndex+rowOffset]
		if len(line) != 4 {
			return Tetromino{}, 0, fmt.Errorf("%w: tetromino %c line %d has width %d", ErrInvalidRowWidth, letter, startIndex+rowOffset+1, len(line))
		}

		for col, char := range line {
			if char != '#' && char != '.' {
				return Tetromino{}, 0, fmt.Errorf("%w: tetromino %c line %d col %d contains %q", ErrInvalidCharacter, letter, startIndex+rowOffset+1, col+1, string(char))
			}

			if char == '#' {
				cells = append(cells, Point{
					Row: rowOffset,
					Col: col,
				})
			}
		}
	}

	tetromino := Tetromino{
		Cells:  cells,
		Letter: letter,
	}

	return tetromino, startIndex + 4, nil
}
