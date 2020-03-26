package area

import (
	"game/config"
	"game/core"
	"game/map/obj"
	xerror "lib/error"
	"proto/net"
)

/// true 碰撞 false 不会碰撞
func CheckObjCollide(obj1, obj2 *obj.Obj) bool {
	if obj1 == obj2 {
		return false
	}
	pos1 := core.Pos2PPos(obj1.GetArea(), obj1.GetPos())
	pos2 := core.Pos2PPos(obj2.GetArea(), obj2.GetPos())
	cover1 := obj1.GetSize()
	cover2 := obj2.GetSize()
	return CheckPosCollide(obj1, obj2, pos1, pos2, cover1, cover2)
}

func CheckPosCollide(obj1, obj2 *obj.Obj, pos1 *net.PPosition, pos2 *net.PPosition, cover1, cover2 core.Position) bool {
	return core.AbsInt32(pos1.X, pos2.X) < (cover1.X+cover2.X)/2 &&
		core.AbsInt32(pos1.Y, pos2.Y) < (cover1.Y+cover2.Y)/2
}

func (self *Area) CheckObjCollide(obj1 *obj.Obj) (*xerror.XError, *obj.Obj) {
	var objT *obj.Obj = nil
	fun := func(objP *obj.Obj) bool {
		if CheckObjCollide(objP, obj1) {
			objT = objP
			return true
		}
		return false
	}
	self.ObjCollideObjsFac(obj1, fun)
	if objT != nil {
		return xerror.New("objsrc: "+obj1.ToString()+
			" objaim: "+objT.ToString()+" Collide", (int32)(config.ERROR_COMMON)), objT
	}

	return nil, objT
}

func (self *Area) ObjCollideObjsFac(objP *obj.Obj, fun func(Obj1 *obj.Obj) bool) {
	for _, item := range self.objs.GetAll() {
		exit := fun(item.(*obj.Obj))
		if exit {
			return
		}
	}
	switch obj.GetOverDirection(objP) {
	case core.DIRECTION_RIGHT:
		if AreaR := self.GetRightArea(); AreaR != nil {
			for _, itemT := range AreaR.objs.GetAll() {
				exit := fun(itemT.(*obj.Obj))
				if exit {
					return
				}
			}
		}
	case core.DIRECTION_LEFT:
		if AreaL := self.GetLeftArea(); AreaL != nil {
			for _, itemT := range AreaL.objs.GetAll() {
				exit := fun(itemT.(*obj.Obj))
				if exit {
					return
				}
			}
		}
	case core.DIRECTION_UP:
		if AreaT := self.GetUpArea(); AreaT != nil {
			for _, itemT := range AreaT.objs.GetAll() {
				exit := fun(itemT.(*obj.Obj))
				if exit {
					return
				}
			}
		}
	case core.DIRECTION_DOWN:
		if AreaB := self.GetDownArea(); AreaB != nil {
			for _, itemT := range AreaB.objs.GetAll() {
				exit := fun(itemT.(*obj.Obj))
				if exit {
					return
				}
			}
		}
	default:
		// nothing
	}
}
