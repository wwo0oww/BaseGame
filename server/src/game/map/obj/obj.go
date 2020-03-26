package obj

import (
	"fmt"
	"game/config"
	"game/core"
	xerror "lib/error"
	"lib/log"
	"math/rand"
	"proto/net"
)

var MAXOBJ int32 = 0

type Obj struct {
	ObjID      int32
	Type       int32
	Pos        core.Position
	Size       core.Position
	Status     ObjStatus
	PlayerFlag int
	Name       string
	// 延迟的帧数
	RelayFrame int32

	areaPr interface{}

	speed     int32
	direction core.DIRECTION
}

type IArea interface {
	CanAddObj(newObj *Obj) (*xerror.XError, *Obj)
	CanMoveObj(newObj *Obj) (*xerror.XError, *Obj)
	G_GetRightAreas() []interface{}
	G_GetLeftAreas() []interface{}
	G_GetUpAreas() []interface{}
	G_GetDownAreas() []interface{}
	DoAddObj(*Obj, core.DIRECTION)
	DoRemoveObj(*Obj, core.DIRECTION)
	GetPos() core.Position
	TrySendMsgToRegPlayersByPos(interface{}, *net.PPosition)
	GetMap() interface{}
	GetArea(x, y int32) interface{}
	GetMapFrame() uint32
}

func (self *Obj) SetArea(area interface{}) {
	self.areaPr = area
}

func (self *Obj) GetArea() interface{} {
	return self.areaPr
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

func (self *Obj) GetObjType() int32 {
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

func (self *Obj) MakeSureObjPos(area interface{}) *xerror.XError {
	self.SetArea(area)
	var distance, xdistance, ydistance int32
	distance = 0
	count := 0
	PPos := core.Pos2PPos(self.areaPr, self.Pos)
	W, H := config.GetMapWH()
	pos := core.Position{X: PPos.X, Y: PPos.Y}
	for true {
		distance++
		size := self.Size.X
		xdistance = -distance
		failIndex := 0
		MaxIndex := 0
		findFun := func() bool {
			MaxIndex++
			self.Pos.X += (size * xdistance)
			self.Pos.Y += (size * ydistance)
			areaPos := core.GetAreaPosByGPos(self.Pos)
			if areaPos.X >= W || areaPos.Y >= H || areaPos.X < 0 || areaPos.Y < 0 {
				return false
			}

			if area := self.areaPr.(IArea).GetArea(areaPos.X, areaPos.Y); area != nil {
				self.areaPr = area
				err, _ := self.areaPr.(IArea).CanAddObj(self)
				if err == nil {
					// find
					log.Debug("*********************", self.Pos)
					return true
				}
			} else {
				failIndex++
			}
			return false
		}
		for ydistance = -distance; ydistance < distance; ydistance++ {
			self.Pos = pos
			if findFun() {
				return nil
			}
		}
		xdistance = distance
		for ydistance = -distance; ydistance < distance; ydistance++ {
			self.Pos = pos
			if findFun() {
				return nil
			}
		}

		ydistance = -distance
		for xdistance = -distance; xdistance < distance; xdistance++ {
			self.Pos = pos
			if findFun() {
				return nil
			}
		}

		ydistance = distance
		for xdistance = -distance; xdistance < distance; xdistance++ {
			self.Pos = pos
			if findFun() {
				return nil
			}
		}
		if failIndex >= MaxIndex {
			count++
			if count > 4 {
				return xerror.New(fmt.Sprintf("not find %s", self.Pos.ToString()), (int32)(config.ERROR_COMMON))
			}
		}
	}
	return nil
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

func (self *Obj) GetDirection() core.DIRECTION {
	return self.direction
}

func (self *Obj) SetDirection(direction core.DIRECTION) {
	self.Status = STATUS_MOVE
	self.direction = direction
}

func (self *Obj) SetSpeed(speed int32) {
	self.speed = speed
}

func (self *Obj) GetSpeed() int32 {
	return self.speed
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

func GetObj2ObjDistance(obj1 Obj, obj2 Obj) (int32, int32) {
	return core.AbsInt32(obj1.Pos.X, obj2.Pos.X), core.AbsInt32(obj1.Pos.Y, obj2.Pos.Y)
}

func GetMaxObj() int32 {
	MAXOBJ += 1
	return MAXOBJ
}

func GenPObj(objP *Obj) *net.PObj {
	return &net.PObj{
		Id:        objP.ObjID,
		Type:      objP.Type,
		Name:      objP.Name,
		Pos:       core.Pos2PPos(objP.GetArea(), objP.Pos),
		Status:    (int32)(objP.Status),
		Direction: (int32)(objP.GetDirection()),
		Speed:     objP.GetSpeed()}
}

func init() {
	MAXOBJ = 0
	// todo 持久化 存庫
}
