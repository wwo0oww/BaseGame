package obj

import (
	"game/config"
	"game/core"
	"math/rand"
	"proto/net"
)

func (self *Obj) Update() {
	//println(self.Pos.ToString())
	switch self.Status {
	case STATUS_NONE:
		//todo
	case STATUS_MOVE:
		self.Move()
	default:
		//todo
	}
}
func (self *Obj) Move() {
	self.StartMove(self.direction)
}
func (self *Obj) StartMove(direction core.DIRECTION) core.DIRECTION {
	//todo
	//log.Debug(" ", core.Pos2PPos(self.areaPr, self.Pos).X, " ", core.Pos2PPos(self.areaPr, self.Pos).Y)

	switch direction {
	case core.DIRECTION_RIGHT:
		return self.MoveRight()
	case core.DIRECTION_LEFT:
		return self.MoveLeft()
	case core.DIRECTION_UP:
		return self.MoveUp()
	case core.DIRECTION_DOWN:
		return self.MoveDown()
	}
	return core.DIRECTION_STOP
}

func (self *Obj) ChangeDirection(direction core.DIRECTION) core.DIRECTION {
	//log.Debug("pos:", self.Pos.ToString(), "area:", self.areaPr.(IArea).GetPos().X, "_", self.areaPr.(IArea).GetPos().X, self.direction)
	// todo 非player，未处理，
	if self.direction != core.DIRECTION_STOP && self.direction != core.DIRECTION_NONE {
		if self.Type == T_OBJ.NPC().TEST().Value() {
			self.direction = core.Directions[rand.Intn(4)]
			self.areaPr.(IArea).TrySendMsgToRegPlayersByPos(&net.MObjUpdateToc{
				Type:    (int32)(core.ObjUpdateType_Upt),
				ObjInfo: []*net.PObj{GenPObj(self)}}, core.Pos2PPos(self.areaPr, self.Pos))
		} else {
			self.StopMove()
		}
	}
	return core.DIRECTION_STOP
}

func (self *Obj) StopMove() {
	self.Status = STATUS_NONE
	self.direction = core.DIRECTION_NONE
	// 告诉所有注册玩家 ，当前obj修改
	self.areaPr.(IArea).TrySendMsgToRegPlayersByPos(&net.MObjUpdateToc{
		Type:    (int32)(core.ObjUpdateType_Upt),
		ObjInfo: []*net.PObj{GenPObj(self)}}, core.Pos2PPos(self.areaPr, self.Pos))
}

func (self *Obj) MoveRight() core.DIRECTION {
	NewPos := core.Position{X: self.Pos.X + self.speed, Y: self.Pos.Y}
	OldPos := self.Pos
	self.Pos = NewPos

	if NewPos.X >= config.GetAreaSize() {
		areasT := self.areaPr.(IArea).G_GetRightAreas()
		if len(areasT) == 0 {
			self.Pos = OldPos
			return self.ChangeDirection(core.DIRECTION_RIGHT)
		}
		for _, areaT := range areasT {
			if err, _ := areaT.(IArea).CanMoveObj(self); err != nil {
				self.Pos = OldPos
				return self.ChangeDirection(core.DIRECTION_RIGHT)
			}
		}
		areaPr := self.areaPr
		areasT[0].(IArea).DoAddObj(self, core.DIRECTION_RIGHT)
		areaPr.(IArea).DoRemoveObj(self, core.DIRECTION_RIGHT)
		return core.DIRECTION_RIGHT
	}
	if err, _ := self.areaPr.(IArea).CanMoveObj(self); err != nil {
		self.Pos = OldPos
		return self.ChangeDirection(core.DIRECTION_RIGHT)
	}
	return core.DIRECTION_RIGHT
}
func (self *Obj) MoveLeft() core.DIRECTION {
	NewPos := core.Position{X: self.Pos.X - self.speed, Y: self.Pos.Y}
	OldPos := self.Pos
	self.Pos = NewPos

	if NewPos.X < 0 {
		areasT := self.areaPr.(IArea).G_GetLeftAreas()
		if len(areasT) == 0 {
			self.Pos = OldPos
			return self.ChangeDirection(core.DIRECTION_LEFT)
		}
		for _, areaT := range areasT {
			if err, _ := areaT.(IArea).CanMoveObj(self); err != nil {
				self.Pos = OldPos
				return self.ChangeDirection(core.DIRECTION_LEFT)
			}
		}
		areaPr := self.areaPr
		areasT[0].(IArea).DoAddObj(self, core.DIRECTION_LEFT)
		areaPr.(IArea).DoRemoveObj(self, core.DIRECTION_LEFT)
		return core.DIRECTION_LEFT
	}
	if err, _ := self.areaPr.(IArea).CanMoveObj(self); err != nil {
		self.Pos = OldPos
		return self.ChangeDirection(core.DIRECTION_LEFT)
	}
	return core.DIRECTION_LEFT
}
func (self *Obj) MoveUp() core.DIRECTION {
	NewPos := core.Position{X: self.Pos.X, Y: self.Pos.Y + self.speed}
	OldPos := self.Pos
	self.Pos = NewPos

	if NewPos.Y >= config.GetAreaSize() {
		areasT := self.areaPr.(IArea).G_GetUpAreas()
		if len(areasT) == 0 {
			self.Pos = OldPos
			return self.ChangeDirection(core.DIRECTION_UP)
		}
		for _, areaT := range areasT {
			if err, _ := areaT.(IArea).CanMoveObj(self); err != nil {
				self.Pos = OldPos
				return self.ChangeDirection(core.DIRECTION_UP)
			}
		}
		areaPr := self.areaPr
		areasT[0].(IArea).DoAddObj(self, core.DIRECTION_UP)
		areaPr.(IArea).DoRemoveObj(self, core.DIRECTION_UP)
		return core.DIRECTION_UP
	}
	if err, _ := self.areaPr.(IArea).CanMoveObj(self); err != nil {
		self.Pos = OldPos
		return self.ChangeDirection(core.DIRECTION_UP)
	}
	return core.DIRECTION_UP
}
func (self *Obj) MoveDown() core.DIRECTION {
	NewPos := core.Position{X: self.Pos.X, Y: self.Pos.Y - self.speed}
	OldPos := self.Pos
	self.Pos = NewPos

	if NewPos.Y < 0 {
		areasT := self.areaPr.(IArea).G_GetDownAreas()
		if len(areasT) == 0 {
			self.Pos = OldPos
			return self.ChangeDirection(core.DIRECTION_DOWN)
		}
		for _, areaT := range areasT {
			if err, _ := areaT.(IArea).CanMoveObj(self); err != nil {
				self.Pos = OldPos
				return self.ChangeDirection(core.DIRECTION_DOWN)
			}
		}
		areaPr := self.areaPr
		areasT[0].(IArea).DoAddObj(self, core.DIRECTION_DOWN)
		areaPr.(IArea).DoRemoveObj(self, core.DIRECTION_DOWN)
		return core.DIRECTION_DOWN
	}
	if err, _ := self.areaPr.(IArea).CanMoveObj(self); err != nil {
		self.Pos = OldPos
		return self.ChangeDirection(core.DIRECTION_DOWN)
	}
	return core.DIRECTION_DOWN
}
