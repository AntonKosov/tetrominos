package states

import (
	"sync"
	"tetrominos/input"
	"tetrominos/settings"
	"tetrominos/states/gamestate"
	t "tetrominos/tetrominos"
	"tetrominos/ticker"
)

type gameState struct {
	params      Params
	field       gamestate.Field
	generator   gamestate.Generator
	tickerGroup *ticker.Group

	col              int
	row              int
	currentTetromino *t.Tetromino
	nextTetromino    t.Tetromino

	generateNewTetrominoSignal chan struct{}

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
		params:    params,
		field:     gamestate.NewField(),
		generator: gamestate.Generator{},

		// buffered, to avoid issues if the state is disactivating
		generateNewTetrominoSignal: make(chan struct{}, 1),

		stopSignal:  make(chan struct{}),
		input:       make(chan input.Input),
		tickerGroup: ticker.NewGroup(),
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
		s.tickerGroup.Resume()
		s.isPaused = false
		return
	}
	s.params.GameView.Activate(s.tickerGroup)
	s.runControl()
}

func (s *gameState) Deactivate() {
	s.params.GameView.Deactivate()
	if s.isPaused {
		s.tickerGroup.Pause()
		return
	}
	close(s.stopSignal)
	close(s.input)
	s.wg.Wait()
	close(s.generateNewTetrominoSignal)
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
		tickID, ticker := s.tickerGroup.NewTicker(gamestate.GetLevel(0).Delay)
		defer s.tickerGroup.DeleteTicker(tickID)
		for {
			select {
			case <-s.stopSignal:
				return
			case key := <-s.input:
				if action, ok := s.gameInputHandlers[key]; ok {
					action()
				}
			case <-s.generateNewTetrominoSignal:
				if !s.generateNewTetromino() {
					s.params.ChangeState <- newGameOverState(s.params, s.score)
					return
				}
				if s.level < gamestate.MaxLevel() {
					nextLevel := gamestate.GetLevel(s.level + 1)
					if s.removedRows >= nextLevel.Rows {
						s.level++
						s.params.GameView.OutputLevel(s.level)
					}
				}
				s.tickerGroup.Reset(tickID, gamestate.GetLevel(s.level).Delay)
			case <-ticker:
				if !s.isPaused && s.currentTetromino != nil {
					s.moveDown()
				}
			}
		}
	}()
}

func (s *gameState) generateNewTetromino() bool {
	ct := s.nextTetromino
	s.currentTetromino = &ct
	s.col = (settings.FieldWidth - s.currentTetromino.Size()) / 2
	s.row = 0
	if !s.field.CanBePlaced(s.col, s.row, *s.currentTetromino) {
		return false
	}
	s.nextTetromino = s.generator.GetNextTetromino()
	s.params.GameView.OutputNextTetromino(s.nextTetromino)
	s.params.GameView.Draw(s.col, s.row, ct)
	return true
}

func (s *gameState) moveDown() bool {
	if s.field.CanBePlaced(s.col, s.row+1, *s.currentTetromino) {
		s.move(0, 1, *s.currentTetromino)
		return true
	}
	removedRows, changedRows := s.field.SetTetromino(s.col, s.row, *s.currentTetromino)
	s.currentTetromino = nil
	s.wg.Add(1)
	go func() {
		s.wg.Done()
		rowsCount := len(removedRows)
		if rowsCount > 0 {
			s.removedRows += rowsCount
			score := gamestate.Score(rowsCount)
			s.score += score
			s.params.GameView.RemoveRows(removedRows, changedRows, score)
			s.params.GameView.OutputScore(s.score)
		}
		s.generateNewTetrominoSignal <- struct{}{}
	}()
	return false
}

func (s *gameState) moveLeft() {
	if s.field.CanBePlaced(s.col-1, s.row, *s.currentTetromino) {
		s.move(-1, 0, *s.currentTetromino)
	}
}

func (s *gameState) moveRight() {
	if s.field.CanBePlaced(s.col+1, s.row, *s.currentTetromino) {
		s.move(1, 0, *s.currentTetromino)
	}
}

func (s *gameState) move(dc, dr int, newT t.Tetromino) {
	s.params.GameView.Move(s.col, s.row, *s.currentTetromino, s.col+dc, s.row+dr, newT)
	s.col += dc
	s.row += dr
	s.currentTetromino = &newT
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
	for s.field.CanBePlaced(s.col, s.row+1, *s.currentTetromino) {
		s.moveDown()
	}
}

func (s *gameState) pause() {
	s.isPaused = true
	s.params.ChangeState <- newPauseState(s.params, s)
}
