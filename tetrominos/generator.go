package tetrominos

const f = false
const t = true

var tetrimonoFactory map[TetrominoType]func() Tetromino

func init() {
	tetrimonoFactory = map[TetrominoType]func() Tetromino{
		TetrominoI: newTetrominoI,
		TetrominoJ: newTetrominoJ,
		TetrominoL: newTetrominoL,
		TetrominoO: newTetrominoO,
		TetrominoS: newTetrominoS,
		TetrominoT: newTetrominoT,
		TetrominoZ: newTetrominoZ,
	}
}

func Generate(t TetrominoType) Tetromino {
	return tetrimonoFactory[t]()
}

func createEmpty(t Tetromino) Tetromino {
	var shape [][]bool
	switch t.Size() {
	case 2:
		shape = [][]bool{
			{f, f},
			{f, f},
		}
	case 3:
		shape = [][]bool{
			{f, f, f},
			{f, f, f},
			{f, f, f},
		}
	case 4:
		shape = [][]bool{
			{f, f, f, f},
			{f, f, f, f},
			{f, f, f, f},
			{f, f, f, f},
		}
	default:
		panic("unexpected shape size")
	}

	return Tetromino{
		Type:  t.Type,
		Shape: shape,
	}
}

func newTetrominoI() Tetromino {
	return Tetromino{
		Type: TetrominoI,
		Shape: [][]bool{
			{f, f, f, f},
			{t, t, t, t},
			{f, f, f, f},
			{f, f, f, f},
		},
	}
}

func newTetrominoJ() Tetromino {
	return Tetromino{
		Type: TetrominoJ,
		Shape: [][]bool{
			{t, f, f},
			{t, t, t},
			{f, f, f},
		},
	}
}

func newTetrominoL() Tetromino {
	return Tetromino{
		Type: TetrominoL,
		Shape: [][]bool{
			{f, f, t},
			{t, t, t},
			{f, f, f},
		},
	}
}

func newTetrominoO() Tetromino {
	return Tetromino{
		Type: TetrominoO,
		Shape: [][]bool{
			{t, t},
			{t, t},
		},
	}
}

func newTetrominoS() Tetromino {
	return Tetromino{
		Type: TetrominoS,
		Shape: [][]bool{
			{f, t, t},
			{t, t, f},
			{f, f, f},
		},
	}
}

func newTetrominoT() Tetromino {
	return Tetromino{
		Type: TetrominoT,
		Shape: [][]bool{
			{f, t, f},
			{t, t, t},
			{f, f, f},
		},
	}
}

func newTetrominoZ() Tetromino {
	return Tetromino{
		Type: TetrominoZ,
		Shape: [][]bool{
			{t, t, f},
			{f, t, t},
			{f, f, f},
		},
	}
}
