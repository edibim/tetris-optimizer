package internal

import (
	"os"
	"strings"
)

func ReadFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, ErrInvalidInput
	}

	return data, nil
}

func ParseFile(data []byte) ([]Tetromino, error) {
	content := normalizeNewlines(string(data))
	if content == "" {
		return nil, ErrInvalidInput
	}

	lines := strings.Split(content, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	tetrominoes := make([]Tetromino, 0)
	index := 0

	for index < len(lines) {
		if lines[index] == "" {
			return nil, ErrInvalidInput
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
			return nil, ErrInvalidInput
		}

		index++
		if index == len(lines) {
			return nil, ErrInvalidInput
		}
	}

	if len(tetrominoes) == 0 {
		return nil, ErrInvalidInput
	}

	return tetrominoes, nil
}

func normalizeNewlines(content string) string {
	return strings.ReplaceAll(content, "\r\n", "\n")
}

func parseTetrominoBlock(lines []string, startIndex int, tetrominoIndex int) (Tetromino, int, error) {
	if startIndex+4 > len(lines) {
		return Tetromino{}, 0, ErrInvalidInput
	}

	cells := make([]Point, 0, 4)
	for rowOffset := 0; rowOffset < 4; rowOffset++ {
		line := lines[startIndex+rowOffset]
		if len(line) != 4 {
			return Tetromino{}, 0, ErrInvalidInput
		}

		for col, char := range line {
			if char != '#' && char != '.' {
				return Tetromino{}, 0, ErrInvalidInput
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
		Letter: rune('A' + tetrominoIndex),
	}

	return tetromino, startIndex + 4, nil
}
