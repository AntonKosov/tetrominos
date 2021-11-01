package gamestate

import (
	"tetrominos/settings"
	t "tetrominos/tetrominos"
)

type Field struct {
	data []t.FieldRow
}

func NewField() Field {
	f := Field{
		data: make([]t.FieldRow, 0, settings.FieldHeight),
	}

	for i := 0; i < settings.FieldHeight; i++ {
		f.data = append(f.data, createNewRow())
	}

	return f
}

func (f *Field) CanBePlaced(col, row int, t t.Tetromino) bool {
	for r := 0; r < t.Size(); r++ {
		tRow := t.Shape[r]
		for c := 0; c < t.Size(); c++ {
			if !tRow[c] {
				continue
			}
			testCol := col + c
			if testCol < 0 || testCol >= settings.FieldWidth {
				return false
			}
			testRow := row + r
			if testRow < 0 || testRow >= settings.FieldHeight {
				return false
			}
			if !f.isFree(testCol, testRow) {
				return false
			}
		}
	}

	return true
}

func (f *Field) isFree(column, row int) bool {
	return f.data[row][column] == nil
}

// SetTetromino returns the number of removed rows and changed rows
func (f *Field) SetTetromino(col, row int, tr t.Tetromino) (int, []t.FieldRow) {
	f.bakeTetromino(col, row, tr)
	bottomRow := row + tr.Size() - 1
	if bottomRow >= settings.FieldHeight {
		bottomRow = settings.FieldHeight - 1
	}
	removedRows, lowestRemovedRow := f.removeFilledRows(row, bottomRow)
	var changedRows []t.FieldRow
	if removedRows > 0 {
		changedRows = f.data[:lowestRemovedRow+1]
	}
	return removedRows, changedRows
}

func (f *Field) bakeTetromino(col, row int, t t.Tetromino) {
	for c := 0; c < t.Size(); c++ {
		for r := 0; r < t.Size(); r++ {
			if t.Shape[r][c] {
				f.data[row+r][col+c] = &t.Type
			}
		}
	}
}

// removeFilledRows returns the number of removed rows and the index of the
// lowest removed row
func (f *Field) removeFilledRows(topRow, bottomRow int) (int, int) {
	removedRows := 0
	lowestRemovedRow := 0
	for r := topRow; r <= bottomRow; r++ {
		isFull := true
		for c := 0; c < settings.FieldWidth; c++ {
			if f.isFree(c, r) {
				isFull = false
				break
			}
		}
		if !isFull {
			continue
		}
		removedRows++
		lowestRemovedRow = r
		f.data = append(
			[]t.FieldRow{createNewRow()},
			append(f.data[:r], f.data[r+1:]...)...,
		)
	}

	return removedRows, lowestRemovedRow
}

func createNewRow() t.FieldRow {
	return make(t.FieldRow, settings.FieldWidth)
}
