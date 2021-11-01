package tetrominos

type TetrominoType rune

const (
	TetrominoI TetrominoType = 'I'
	TetrominoJ TetrominoType = 'J'
	TetrominoL TetrominoType = 'L'
	TetrominoO TetrominoType = 'O'
	TetrominoS TetrominoType = 'S'
	TetrominoT TetrominoType = 'T'
	TetrominoZ TetrominoType = 'Z'
)

type FieldRow []*TetrominoType
