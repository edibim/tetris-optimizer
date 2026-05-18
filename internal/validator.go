package internal

import "sort"

// ValidateTetromino checks the core shape rules required by the project.
func ValidateTetromino(tetromino Tetromino) error {
	if len(tetromino.Cells) != 4 {
		return ErrInvalidBlockCount
	}

	if !isConnected(tetromino.Cells) {
		return ErrDisconnectedShape
	}

	return nil
}

// NormalizeTetromino shifts the shape to the top-left origin and sorts it.
func NormalizeTetromino(tetromino Tetromino) Tetromino {
	if len(tetromino.Cells) == 0 {
		return tetromino
	}

	minRow := tetromino.Cells[0].Row
	minCol := tetromino.Cells[0].Col
	for _, cell := range tetromino.Cells[1:] {
		if cell.Row < minRow {
			minRow = cell.Row
		}

		if cell.Col < minCol {
			minCol = cell.Col
		}
	}

	normalized := Tetromino{
		Cells:  make([]Point, len(tetromino.Cells)),
		Letter: tetromino.Letter,
	}

	for index, cell := range tetromino.Cells {
		normalized.Cells[index] = Point{
			Row: cell.Row - minRow,
			Col: cell.Col - minCol,
		}
	}

	sort.Slice(normalized.Cells, func(i int, j int) bool {
		if normalized.Cells[i].Row != normalized.Cells[j].Row {
			return normalized.Cells[i].Row < normalized.Cells[j].Row
		}

		return normalized.Cells[i].Col < normalized.Cells[j].Col
	})

	return normalized
}

func isConnected(cells []Point) bool {
	visited := map[Point]bool{
		cells[0]: true,
	}
	queue := []Point{cells[0]}
	cellSet := make(map[Point]bool, len(cells))
	for _, cell := range cells {
		cellSet[cell] = true
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighbor := range orthogonalNeighbors(current) {
			if !cellSet[neighbor] || visited[neighbor] {
				continue
			}

			visited[neighbor] = true
			queue = append(queue, neighbor)
		}
	}

	return len(visited) == len(cells)
}

func orthogonalNeighbors(point Point) []Point {
	return []Point{
		{Row: point.Row - 1, Col: point.Col},
		{Row: point.Row + 1, Col: point.Col},
		{Row: point.Row, Col: point.Col - 1},
		{Row: point.Row, Col: point.Col + 1},
	}
}
