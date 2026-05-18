package internal

// Point represents a single occupied cell in row/column coordinates.
type Point struct {
	Row int
	Col int
}

// Tetromino stores a normalized shape together with its output letter.
type Tetromino struct {
	Cells  []Point
	Letter rune
}

// Width returns the occupied width of the tetromino footprint.
func (t Tetromino) Width() int {
	maxCol := 0
	for _, cell := range t.Cells {
		if cell.Col > maxCol {
			maxCol = cell.Col
		}
	}

	return maxCol + 1
}

// Height returns the occupied height of the tetromino footprint.
func (t Tetromino) Height() int {
	maxRow := 0
	for _, cell := range t.Cells {
		if cell.Row > maxRow {
			maxRow = cell.Row
		}
	}

	return maxRow + 1
}
