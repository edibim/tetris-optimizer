package internal

import "os"

func ReadFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, ErrInvalidInput
	}

	return data, nil
}

func ParseFile(data []byte) ([]Tetromino, error) {
	_ = data

	return nil, ErrNotImplemented
}
