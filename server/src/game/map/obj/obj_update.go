package obj

import (
	"game/config"
	"game/core"
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
	//todo
	switch self.direction {
	case core.DIRECTION_RIGHT:
		self.MoveRight()
	case core.DIRECTION_LEFT:
		self.MoveLeft()
	case core.DIRECTION_UP:
		self.MoveUp()
	case core.DIRECTION_DOWN:
		self.MoveDown()
	}
}

func (self *Obj) ChangeDirection() {
	println("pos:",self.Pos.ToString(), "area:",self.areaPr.(IArea).GetPos().X,"_",self.areaPr.(IArea).GetPos().X)
	self.Status = STATUS_NONE
	self.direction = core.DIRECTION_NONE
}

func (self *Obj) MoveRight() {
	NewPos := core.Position{X: self.Pos.X + self.speed, Y: self.Pos.Y}
	OldPos := self.Pos
	self.Pos = NewPos

	if NewPos.X >= config.GetAreaSize() {
		areaT := self.areaPr.(IArea).G_GetRightArea()
		if areaT == nil {
			self.Pos = OldPos
			self.ChangeDirection()
			return
		}
		if err, _ := areaT.(IArea).CanAddObj(self); err != nil {
			self.Pos = OldPos
			self.ChangeDirection()
			return
		}
		areaT.(IArea).DoAddObj(self)
		self.areaPr.(IArea).DoRemoveObj(self)
		return
	}
	if err, _ := self.areaPr.(IArea).CanAddObj(self); err != nil {
		self.Pos = OldPos
		self.ChangeDirection()
	}
}
func (self *Obj) MoveLeft() {
	NewPos := core.Position{X: self.Pos.X - self.speed, Y: self.Pos.Y}
	OldPos := self.Pos
	self.Pos = NewPos

	if NewPos.X < 0 {
		areaT := self.areaPr.(IArea).G_GetLeftArea()
		if areaT == nil {
			self.Pos = OldPos
			self.ChangeDirection()
			return
		}

		if err, _ := areaT.(IArea).CanAddObj(self); err != nil {
			self.Pos = OldPos
			self.ChangeDirection()
			return
		}
		areaT.(IArea).DoAddObj(self)
		self.areaPr.(IArea).DoRemoveObj(self)
		return
	}
	if err, _ := self.areaPr.(IArea).CanAddObj(self); err != nil {
		self.Pos = OldPos
		self.ChangeDirection()
	}
}
func (self *Obj) MoveUp() {
	NewPos := core.Position{X: self.Pos.X, Y: self.Pos.Y + self.speed}
	OldPos := self.Pos
	self.Pos = NewPos

	if NewPos.X >= config.GetAreaSize() {
		areaT := self.areaPr.(IArea).G_GetUpArea()
		if areaT == nil {
			self.Pos = OldPos
			self.ChangeDirection()
			return
		}
		if err, _ := areaT.(IArea).CanAddObj(self); err != nil {
			self.Pos = OldPos
			self.ChangeDirection()
			return
		}
		areaT.(IArea).DoAddObj(self)
		self.areaPr.(IArea).DoRemoveObj(self)
		return
	}
	if err, _ := self.areaPr.(IArea).CanAddObj(self); err != nil {
		self.Pos = OldPos
		self.ChangeDirection()
	}
}
func (self *Obj) MoveDown() {
	NewPos := core.Position{X: self.Pos.X + self.speed, Y: self.Pos.Y}
	OldPos := self.Pos
	self.Pos = NewPos

	if NewPos.X >= config.GetAreaSize() {
		areaT := self.areaPr.(IArea).G_GetDownArea()
		if areaT == nil {
			self.Pos = OldPos
			self.ChangeDirection()
			return
		}
		if err, _ := areaT.(IArea).CanAddObj(self); err != nil {
			self.Pos = OldPos
			self.ChangeDirection()
			return
		}
		areaT.(IArea).DoAddObj(self)
		self.areaPr.(IArea).DoRemoveObj(self)
		return
	}
	if err, _ := self.areaPr.(IArea).CanAddObj(self); err != nil {
		self.Pos = OldPos
		self.ChangeDirection()
	}
}
