package mMsg

import (
	"game/core"
	"game/map/obj"
	"proto"
)

var start_index int32

func init() {
	start_index = 4000
	proto.AddMap((*JoinObj)(nil))
}

type IMapMsg interface {
	ToString() string
	GetAreaPos() core.Position
}

type JoinObj struct {
	Obj   *obj.Obj
	Times int32
}

func (self *JoinObj) ToString() string {
	return "JoinObj:" + self.Obj.ToString()
}

func (self *JoinObj) GetAreaPos() core.Position {
	return core.GetObjAreaPos(self.Obj.Pos)
}
