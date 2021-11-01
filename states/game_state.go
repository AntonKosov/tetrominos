package states

import (
	"sync"
	"tetrominos/input"
	"tetrominos/settings"
	"tetrominos/states/gamestate"
	t "tetrominos/tetrominos"
	"time"
)

type gameState struct {
	params    Params
	field     gamestate.Field
	generator gamestate.Generator

	col              int
	row              int
	currentTetromino t.Tetromino
	nextTetromino    t.Tetromino

	score       int
	level       int
	removedRows int

	stopSignal chan struct{}
	wg         sync.WaitGroup
	input      chan input.Input

	gameInputHandlers map[input.Input]func()

	isPaused bool
}

func newGameState(params Params) *gameState {
	gs := &gameState{
		params:     params,
		field:      gamestate.NewField(),
		generator:  gamestate.Generator{},
		stopSignal: make(chan struct{}),
		input:      make(chan input.Input),
	}
	gs.gameInputHandlers = map[input.Input]func(){
		input.EscKey:   gs.pause,
		input.LeftKey:  gs.moveLeft,
		input.RightKey: gs.moveRight,
		input.UpKey:    gs.rotateRight,
		// input.DownKey:  gs.rotateLeft,
		input.DownKey: gs.drop,
	}

	return gs
}

func (s *gameState) Activate() {
	s.outputControlsHint()
	if s.isPaused {
		s.isPaused = false
		return
	}
	s.params.GameView.Activate()
	s.runControl()
}

func (s *gameState) Deactivate() {
	s.params.GameView.Deactivate()
	if s.isPaused {
		return
	}
	close(s.stopSignal)
	close(s.input)
	s.wg.Wait()
}

func (s *gameState) HandleInput(in input.Input) {
	s.input <- in
}

func (s *gameState) outputControlsHint() {
	s.params.GameView.ShowControlHints([]input.KeyDescription{
		{
			Key:         input.EscKey,
			Description: "Pause",
		},
		{
			Key:         input.LeftKey,
			Description: "Move left",
		},
		{
			Key:         input.RightKey,
			Description: "Move right",
		},
		{
			Key:         input.UpKey,
			Description: "Rotate",
		},
		{
			Key:         input.DownKey,
			Description: "Drop",
		},
		ctrlCDescription,
	})
}

func (s *gameState) runControl() {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.nextTetromino = s.generator.GetNextTetromino()
		s.generateNewTetromino()
		ticker := time.NewTicker(gamestate.GetLevel(0).Delay)
		defer ticker.Stop()
		for {
			select {
			case <-s.stopSignal:
				return
			case key := <-s.input:
				if action, ok := s.gameInputHandlers[key]; ok {
					action()
				}
			case <-ticker.C:
				if !s.isPaused && !s.moveDown() {
					s.generateNewTetromino()
					if !s.field.CanBePlaced(s.col, s.row, s.currentTetromino) {
						s.params.ChangeState <- newGameOverState(s.params, s.score)
						return
					}
					if s.level < gamestate.MaxLevel() {
						nextLevel := gamestate.GetLevel(s.level + 1)
						if s.removedRows >= nextLevel.Rows {
							s.level++
							ticker.Reset(nextLevel.Delay)
							s.params.GameView.OutputLevel(s.level)
						}
					}
				}
			}
		}
	}()
}

func (s *gameState) generateNewTetromino() {
	s.currentTetromino = s.nextTetromino
	s.nextTetromino = s.generator.GetNextTetromino()
	s.params.GameView.OutputNextTetromino(s.nextTetromino)
	s.col = (settings.FieldWidth - s.currentTetromino.Size()) / 2
	s.row = 0
	s.params.GameView.Draw(s.col, s.row, s.currentTetromino)
}

func (s *gameState) moveDown() bool {
	if s.field.CanBePlaced(s.col, s.row+1, s.currentTetromino) {
		s.move(0, 1, s.currentTetromino)
		return true
	}
	removedRows, changedRows := s.field.SetTetromino(s.col, s.row, s.currentTetromino)
	if removedRows > 0 {
		s.removedRows += removedRows
		s.score += gamestate.Score(removedRows)
		s.params.GameView.OutputScore(s.score)
		s.params.GameView.DrawRaws(0, changedRows)
	}
	return false
}

func (s *gameState) moveLeft() {
	if s.field.CanBePlaced(s.col-1, s.row, s.currentTetromino) {
		s.move(-1, 0, s.currentTetromino)
	}
}

func (s *gameState) moveRight() {
	if s.field.CanBePlaced(s.col+1, s.row, s.currentTetromino) {
		s.move(1, 0, s.currentTetromino)
	}
}

func (s *gameState) move(dc, dr int, newT t.Tetromino) {
	s.params.GameView.Move(s.col, s.row, s.currentTetromino, s.col+dc, s.row+dr, newT)
	s.col += dc
	s.row += dr
	s.currentTetromino = newT
}

// func (s *gameState) rotateLeft() {
// 	s.rotate(s.currentTetromino.RotateLeft)
// }

func (s *gameState) rotateRight() {
	s.rotate(s.currentTetromino.RotateRight)
}

func (s *gameState) rotate(rotationFunc func() t.Tetromino) {
	rotated := rotationFunc()
	col := s.col
	dc := 0
	if col < 0 {
		dc = -s.col
	} else {
		diff := col + rotated.Size() - settings.FieldWidth
		if diff > 0 {
			dc = -diff
		}
	}
	if !s.field.CanBePlaced(col+dc, s.row, rotated) {
		return
	}
	s.move(dc, 0, rotated)
}

func (s *gameState) drop() {
	for s.field.CanBePlaced(s.col, s.row+1, s.currentTetromino) {
		s.moveDown()
	}
}

func (s *gameState) pause() {
	s.isPaused = true
	s.params.ChangeState <- newPauseState(s.params, s)
}
