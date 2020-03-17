package core

import (
	"fmt"
	"game/config"
)

type Position struct {
	X int32
	Y int32
}

type DIRECTION int32

const (
	DIRECTION_NONE DIRECTION = iota
	DIRECTION_RIGHT
	DIRECTION_LEFT
	DIRECTION_UP
	DIRECTION_DOWN
)

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

func GetObjAreaPos(pos Position) Position {
	size := config.GetAreaSize()
	return Position{X: pos.X / size, Y: pos.X / size}
}
