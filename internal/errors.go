package internal

import "errors"

var (
	ErrInvalidInput        = errors.New("invalid input")
	ErrInvalidArgCount     = errors.New("expected exactly one CLI argument")
	ErrReadFile            = errors.New("failed to read input file")
	ErrEmptyInput          = errors.New("input file is empty")
	ErrUnexpectedBlankLine = errors.New("unexpected blank line")
	ErrTrailingSeparator   = errors.New("trailing separator after last tetromino")
	ErrIncompleteBlock     = errors.New("incomplete tetromino block")
	ErrInvalidRowWidth     = errors.New("tetromino row must be exactly 4 characters wide")
	ErrInvalidCharacter    = errors.New("tetromino contains an invalid character")
	ErrInvalidBlockCount   = errors.New("tetromino must contain exactly 4 blocks")
	ErrDisconnectedShape   = errors.New("tetromino blocks must be orthogonally connected")
	ErrBoardTooSmall       = errors.New("board is too small for the requested placement")
	ErrCellOccupied        = errors.New("board cell is already occupied")
	ErrNoSolution          = errors.New("no valid arrangement found")
)
