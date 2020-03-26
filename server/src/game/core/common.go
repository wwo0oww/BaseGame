package core

import (
	"fmt"
	"game/config"
)

type Position struct {
	X int32
	Y int32
	Z int32
}

type DIRECTION int32

const (
	DIRECTION_NONE DIRECTION = iota
	DIRECTION_LEFT
	DIRECTION_RIGHT
	DIRECTION_UP
	DIRECTION_DOWN
	DIRECTION_STOP
)
const DirectionNum = 6

var Directions = []DIRECTION{DIRECTION_RIGHT, DIRECTION_LEFT, DIRECTION_UP, DIRECTION_DOWN}

func (self *Position) ToString() string {
	return fmt.Sprintf("Postion{X:%d,Y:%d}", self.X, self.Y)
}

func AbsInt32(a, b int32) int32 {
	a = a - b
	if a > 0 {
		return a
	}
	return -a
}

func GetAreaPosByGPos(pos Position) Position {
	if pos.X < 0 || pos.Y < 0 {
		return Position{X: -1, Y: -1}
	}
	size := config.GetAreaSize()
	return Position{X: pos.X / size, Y: pos.Y / size}
}
