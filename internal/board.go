package internal

import (
	"fmt"
	"strings"
)

// Board holds the square grid used by the backtracking solver.
type Board struct {
	Size  int
	Cells [][]rune
}

// NewBoard allocates an empty square board filled with '.' cells.
func NewBoard(size int) *Board {
	cells := make([][]rune, size)
	for row := range cells {
		cells[row] = make([]rune, size)
		for col := range cells[row] {
			cells[row][col] = '.'
		}
	}

	return &Board{
		Size:  size,
		Cells: cells,
	}
}

// CanPlace reports whether a tetromino fits at the given anchor cell.
func (b *Board) CanPlace(tetromino Tetromino, row int, col int) bool {
	for _, cell := range tetromino.Cells {
		targetRow := row + cell.Row
		targetCol := col + cell.Col

		if !b.isInside(targetRow, targetCol) {
			return false
		}

		if b.Cells[targetRow][targetCol] != '.' {
			return false
		}
	}

	return true
}

// PlaceTetromino writes a tetromino to the board after a placement check.
func (b *Board) PlaceTetromino(tetromino Tetromino, row int, col int) error {
	for _, cell := range tetromino.Cells {
		targetRow := row + cell.Row
		targetCol := col + cell.Col

		if !b.isInside(targetRow, targetCol) {
			return fmt.Errorf("%w: row=%d col=%d", ErrBoardTooSmall, targetRow, targetCol)
		}

		if b.Cells[targetRow][targetCol] != '.' {
			return fmt.Errorf("%w: row=%d col=%d", ErrCellOccupied, targetRow, targetCol)
		}
	}

	for _, cell := range tetromino.Cells {
		b.Cells[row+cell.Row][col+cell.Col] = tetromino.Letter
	}

	return nil
}

// RemoveTetromino clears a previously placed tetromino from the board.
func (b *Board) RemoveTetromino(tetromino Tetromino, row int, col int) {
	for _, cell := range tetromino.Cells {
		b.Cells[row+cell.Row][col+cell.Col] = '.'
	}
}

// String renders the board exactly as required by the project output.
func (b *Board) String() string {
	var builder strings.Builder
	for row := range b.Cells {
		builder.WriteString(string(b.Cells[row]))
		if row < len(b.Cells)-1 {
			builder.WriteByte('\n')
		}
	}

	return builder.String()
}

func (b *Board) isInside(row int, col int) bool {
	return row >= 0 && row < b.Size && col >= 0 && col < b.Size
}
