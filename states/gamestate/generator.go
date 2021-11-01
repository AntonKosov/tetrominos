package gamestate

import (
	"math/rand"
	t "tetrominos/tetrominos"
)

type Generator struct {
	queue []t.Tetromino
}

func (g *Generator) GetNextTetromino() t.Tetromino {
	// if g.queue == nil {
	// 	g.queue = shuffle(t.TetrominoI, t.TetrominoI, t.TetrominoI, t.TetrominoI, t.TetrominoI)
	// }
	if g.queue == nil {
		g.queue = shuffle(t.TetrominoI, t.TetrominoJ, t.TetrominoL, t.TetrominoT)
	}
	nextTetromino := g.queue[0]
	if len(g.queue) == 1 {
		// g.queue = shuffle(t.TetrominoI, t.TetrominoI, t.TetrominoI, t.TetrominoI, t.TetrominoI)
		g.queue = shuffle(
			t.TetrominoI,
			t.TetrominoJ,
			t.TetrominoL,
			t.TetrominoO,
			t.TetrominoS,
			t.TetrominoT,
			t.TetrominoZ,
		)
	} else {
		g.queue = g.queue[1:]
	}
	return nextTetromino
}

func shuffle(ts ...t.TetrominoType) []t.Tetromino {
	for i := 0; i < len(ts); i++ {
		p := rand.Intn(len(ts))
		ts[i], ts[p] = ts[p], ts[i]
	}

	r := make([]t.Tetromino, 0, len(ts))
	for _, tt := range ts {
		r = append(r, t.Generate(tt))
	}

	return r
}
