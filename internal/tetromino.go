package internal

type Point struct {
	Row int
	Col int
}

type Tetromino struct {
	Cells  []Point
	Letter rune
}
