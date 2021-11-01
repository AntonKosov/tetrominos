package gamestate

import "fmt"

func Score(removedRows int) int {
	switch removedRows {
	case 1:
		return 1
	case 2:
		return 3
	case 3:
		return 5
	case 4:
		return 8
	}
	panic(fmt.Sprintf("Cannot calculate score for %v rows", removedRows))
}
