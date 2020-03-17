package obj

import (
	"fmt"
	"game/config"
	"game/core"
	"math/rand"
)

var MAXOBJ int32 = 0

type Obj struct {
	ObjID  int32
	Type   ObjType
	Pos    core.Position
	Size   core.Position
	Status ObjStatus

	// 延迟的帧数
	RelayFrame int32

	areaPr interface{}

	speed     int32
	direction core.DIRECTION
}

type IArea interface {
	CanAddObj(newObj *Obj) (error, *Obj)
	G_GetRightArea() interface{}
	G_GetLeftArea() interface{}
	G_GetUpArea() interface{}
	G_GetDownArea() interface{}
	DoAddObj(objP *Obj)
	DoRemoveObj(objP *Obj)
	GetPos() core.Position
}

func (self *Obj) GetPos() core.Position {
	return self.Pos
}

func (self *Obj) ToString() string {
	return fmt.Sprintf("Obj_%.10d", self.ObjID)
}

func (self *Obj) GetSize() core.Position {
	return self.Size
}

func (self *Obj) GetObjType() ObjType {
	return self.Type
}

func (self *Obj) Init(areaPr interface{}) {
	self.areaPr = areaPr
	size := config.GetAreaSize()
	self.Pos.X = (self.Pos.X + size) % size
	self.Pos.Y = (self.Pos.Y + size) % size
}

func (self *Obj) StatusChange() {
	if self.Status == STATUS_NONE {
		self.RelayFrame = 0
	}
}

func (self *Obj) GetWorldPos() core.Position {
	AreaPos := self.areaPr.(IArea).GetPos()
	return core.Position{X: AreaPos.X*config.GetAreaSize() + self.Pos.X, Y: AreaPos.Y*config.GetAreaSize() + self.Pos.Y}
}

func (self *Obj) ChangePosForCollide(obj *Obj) *Obj {
	offX := rand.Int31n((self.Size.X + obj.Size.X) * 20)
	offY := rand.Int31n((self.Size.Y + obj.Size.Y) * 20)
	Pos := obj.GetWorldPos()
	PosT := Pos
	Pos.X += offX
	Pos.Y += offY
	if obj.Pos.X+self.Size.X > config.MaxMapWorldX() {
		Pos.X = PosT.X + (config.MaxMapWorldX()-PosT.X)/2 - self.Size.X
	}
	if obj.Pos.Y+self.Size.Y > config.MaxMapWorldY() {
		Pos.Y = PosT.Y + (config.MaxMapWorldX()-PosT.Y)/2 - self.Size.Y
	}
	self.Pos = Pos
	return self
}

func (self *Obj) AddRelayFrame(nFrame int32) {
	self.RelayFrame += nFrame
}

func (self *Obj) FrameEnd() {
	self.RelayFrame = 0
}

func (self *Obj) SetDirection(direction core.DIRECTION) {
	self.Status = STATUS_MOVE
	self.direction = direction
}

func (self *Obj) SetSpeed(speed int32) {
	self.speed = speed
}

// 物体占位超出的方向
func GetOverDirection(obj *Obj) core.DIRECTION {
	if IsRightOver(obj) {
		return core.DIRECTION_RIGHT
	}
	if IsLeftOver(obj) {
		return core.DIRECTION_LEFT
	}
	if IsUpOver(obj) {
		return core.DIRECTION_UP
	}
	if IsDownOver(obj) {
		return core.DIRECTION_DOWN
	}
	return core.DIRECTION_NONE
}
func IsRightOver(obj *Obj) bool {
	return obj.GetPos().X+obj.GetSize().X/2 >= config.GetAreaSize()
}
func IsLeftOver(obj *Obj) bool {
	return obj.GetPos().X-obj.GetSize().X/2 < 0
}
func IsUpOver(obj *Obj) bool {
	return obj.GetPos().Y+obj.GetSize().Y/2 >= config.GetAreaSize()
}
func IsDownOver(obj *Obj) bool {
	return obj.GetPos().Y-obj.GetSize().Y/2 < 0
}

func GetMaxObj() int32 {
	MAXOBJ += 1
	return MAXOBJ
}

func init() {
	MAXOBJ = 0
	// todo 持久化 存庫
}
