package mMsg

import (
	"game/core"
	"game/map/obj"
	"proto"
)

var start_index int32

var Map interface{} // todo 带优化

type IMap interface {
	GetPlayer2Map() *core.SMap
}

func init() {
	start_index = 4000
	proto.AddMap((*JoinObj)(nil))
	proto.AddMap((*ObjMove)(nil))
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
	return core.GetAreaPosByGPos(self.Obj.Pos)
}

type ObjMove struct {
	Obj       *obj.Obj
	Direction core.DIRECTION
}

func (self *ObjMove) ToString() string {
	return "ObjMove:" + self.Obj.ToString()
}

func (self *ObjMove) GetAreaPos() core.Position {
	// todo 未加容错处理
	return Map.(IMap).GetPlayer2Map().Get(self.Obj.ObjID).(interface{ GetPos() core.Position }).GetPos()
}
