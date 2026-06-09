package internal

// SolveBoard finds the smallest square board that fits all tetrominoes.
func SolveBoard(tetrominoes []Tetromino) (*Board, error) {
	if len(tetrominoes) == 0 {
		return nil, ErrInvalidInput
	}

	size := minBoardSize(len(tetrominoes))
	for {
		board := NewBoard(size)
		if solveFromIndex(board, tetrominoes, 0) {
			return board, nil
		}

		size++
	}
}

// solveFromIndex tries to place tetrominoes one by one.
// This is the visible recursive backtracking step of the solver.
func solveFromIndex(board *Board, tetrominoes []Tetromino, index int) bool {
	if index == len(tetrominoes) {
		return true
	}

	current := tetrominoes[index]
	maxRow := board.Size - current.Height()
	maxCol := board.Size - current.Width()

	for row := 0; row <= maxRow; row++ {
		for col := 0; col <= maxCol; col++ {
			if err := board.PlaceTetromino(current, row, col); err != nil {
				continue
			}

			if solveFromIndex(board, tetrominoes, index+1) {
				return true
			}

			board.RemoveTetromino(current, row, col)
		}
	}

	return false
}

// minBoardSize finds the smallest square that can hold all blocks by area.
// The solver starts here and grows the board only when needed.
func minBoardSize(pieceCount int) int {
	size := 2
	for size*size < pieceCount*4 {
		size++
	}

	return size
}
