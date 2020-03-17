package area

import (
	"errors"
	"game/core"
	mMsg "game/map/msg"
	"game/map/obj"
	"lib/log"
)

func (self *Area) AddObj(objP *obj.Obj, nCount int32) error {
	//posT := objP.Pos
	if err, obj1 := self.CanAddObj(objP); err != nil {
		if obj1 != nil {
			if nCount == 0 {
				return errors.New("area_msg.go add obj fault,over time")
			}
			NewObj := objP.ChangePosForCollide(obj1)
			self.mapPr.(interface{ AddMessage(interface{}) }).AddMessage(&mMsg.JoinObj{
				Obj: NewObj, Times: nCount - 1})
			//log.Debug("碰撞:", posT.ToString(), "<=>", obj1.Pos.ToString(), "new ", NewObj.Pos.ToString())
			return nil
		}
		log.Debug(err.Error())
		return err
	}
	self.DoAddObj(objP)
	return nil
}

func (self *Area) DoAddObj(objP *obj.Obj) {
	objP.Init(self)
	//log.Debug(self.ToString(), "[objlen:",len(self.objs), "] add Obj"+objP.ToString(), objP.Pos.ToString())
	self.objs[objP.ObjID] = objP
	if objP.Type == obj.OBJTYPE_PLAYER {
		self.players[objP.ObjID] = objP
	}
}

func (self *Area) DoRemoveObj(objP *obj.Obj) {
	delete(self.objs, objP.ObjID)
}

func (self *Area) CanAddObj(newObj *obj.Obj) (error, *obj.Obj) {
	switch newObj.GetObjType() {
	case obj.OBJTYPE_PLAYER:
		if self.players[newObj.ObjID] != nil {
			return errors.New("JoinObj error :already exist player " + newObj.ToString()), nil
		}
		var obj *obj.Obj
		var err error
		if err, obj = self.CheckObjCollide(newObj); err == nil {

		}
		return err, obj
	default:
	}
	return self.CheckObjCollide(newObj)
}

func (self *Area) CheckObjCollide(obj1 *obj.Obj) (error, *obj.Obj) {
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
		return errors.New("objsrc: " + obj1.ToString() +
			" objaim: " + objT.ToString() + " Collide"), objT
	}

	return nil, objT
}

func (self *Area) ObjCollideObjsFac(objP *obj.Obj, fun func(Obj1 *obj.Obj) bool) {
	for _, objs := range self.objs {
		exit := fun(objs)
		if exit {
			return
		}
	}
	switch obj.GetOverDirection(objP) {
	case core.DIRECTION_RIGHT:
		if AreaR := self.GetRightArea(); AreaR != nil {
			for _, objT := range AreaR.objs {
				exit := fun(objT)
				if exit {
					return
				}
			}
		}
	case core.DIRECTION_LEFT:
		if AreaL := self.GetLeftArea(); AreaL != nil {
			for _, objT := range AreaL.objs {
				exit := fun(objT)
				if exit {
					return
				}
			}
		}
	case core.DIRECTION_UP:
		if AreaT := self.GetUpArea(); AreaT != nil {
			for _, objT := range AreaT.objs {
				exit := fun(objT)
				if exit {
					return
				}
			}
		}
	case core.DIRECTION_DOWN:
		if AreaB := self.GetDownArea(); AreaB != nil {
			for _, objT := range AreaB.objs {
				exit := fun(objT)
				if exit {
					return
				}
			}
		}
	default:
		// nothing
	}
}
