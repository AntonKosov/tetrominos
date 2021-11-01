package tetrominos

type Tetromino struct {
	Type  TetrominoType
	Shape [][]bool
}

func (t Tetromino) Size() int {
	return len(t.Shape)
}

func (t Tetromino) RotateRight() Tetromino {
	rm := [2][2]int{
		{0, 1},
		{-1, 0},
	}
	return t.rotate(rm, t.Size()-1, 0)
}

func (t Tetromino) RotateLeft() Tetromino {
	rm := [2][2]int{
		{0, -1},
		{1, 0},
	}
	return t.rotate(rm, 0, t.Size()-1)
}

func (t Tetromino) rotate(rm [2][2]int, dc, dr int) Tetromino {
	rotated := createEmpty(t)
	s := t.Size()
	for r := 0; r < s; r++ {
		row := t.Shape[r]
		for c := 0; c < s; c++ {
			if !row[c] {
				continue
			}
			newR := rm[0][0]*r + rm[0][1]*c
			newC := rm[1][0]*r + rm[1][1]*c
			rotated.Shape[dr+newR][dc+newC] = true
		}
	}
	return rotated
}
