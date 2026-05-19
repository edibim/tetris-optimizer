package internal

import (
	"errors"
	"fmt"
)

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

// FormatCLIError converts internal errors into user-facing CLI messages.
func FormatCLIError(err error) string {
	switch {
	case errors.Is(err, ErrInvalidArgCount):
		return "ERROR: expected exactly one input file path.\nUsage: go run . <input-file>"
	case errors.Is(err, ErrReadFile):
		return fmt.Sprintf("ERROR: could not read the input file.\nDetails: %v", err)
	case errors.Is(err, ErrEmptyInput):
		return "ERROR: the input file is empty."
	case errors.Is(err, ErrUnexpectedBlankLine):
		return fmt.Sprintf("ERROR: invalid file format.\nDetails: %v", err)
	case errors.Is(err, ErrTrailingSeparator):
		return fmt.Sprintf("ERROR: invalid file format.\nDetails: %v", err)
	case errors.Is(err, ErrIncompleteBlock):
		return fmt.Sprintf("ERROR: incomplete tetromino block.\nDetails: %v", err)
	case errors.Is(err, ErrInvalidRowWidth):
		return fmt.Sprintf("ERROR: invalid tetromino row width.\nDetails: %v", err)
	case errors.Is(err, ErrInvalidCharacter):
		return fmt.Sprintf("ERROR: invalid character in tetromino file.\nDetails: %v", err)
	case errors.Is(err, ErrInvalidBlockCount):
		return fmt.Sprintf("ERROR: a tetromino must contain exactly 4 blocks.\nDetails: %v", err)
	case errors.Is(err, ErrDisconnectedShape):
		return fmt.Sprintf("ERROR: a tetromino shape is not orthogonally connected.\nDetails: %v", err)
	case errors.Is(err, ErrInvalidInput):
		return fmt.Sprintf("ERROR: invalid puzzle input.\nDetails: %v", err)
	case errors.Is(err, ErrNoSolution):
		return "ERROR: no valid square arrangement was found for the provided tetrominoes."
	default:
		return fmt.Sprintf("ERROR: %v", err)
	}
}
