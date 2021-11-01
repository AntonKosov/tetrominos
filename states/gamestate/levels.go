package gamestate

import "time"

type Level struct {
	Rows  int
	Delay time.Duration
}

var levels []Level

func init() {
	levels = []Level{}
	rows := 0
	for d := 400; d >= 100; d -= 20 {
		levels = append(levels, Level{
			rows,
			time.Millisecond * time.Duration(d),
		})
		rows += 5
	}
}

func GetLevel(l int) Level {
	return levels[l]
}

func MaxLevel() int {
	return len(levels) - 1
}
