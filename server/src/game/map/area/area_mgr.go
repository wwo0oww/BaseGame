package area

import (
	"game/config"
	"game/core"
	mMsg "game/map/msg"
	"game/map/obj"
	xerror "lib/error"
	"lib/log"
	"proto/net"
)

///////////////////////////////////////////////////////////////////////////////////////////////
func (self *Area) TryAddObjAgain(objP *obj.Obj, nCount int32) {
	if nCount == 0 {
		return
	}
	var xerr *xerror.XError
	/// ChangePosForCollide 随机一定范围找n次  MakeSureObjPos 全图找位置
	// 注意这里的查找是异步的，但是真实加入操作是同步的，如果同时大量加入同一位置，会出现竞争同一坐标的情况
	xerr = objP.MakeSureObjPos(self)
	if xerr == nil {
		self.mapPr.(interface{ AddMessage(interface{}) }).AddMessage(&mMsg.JoinObj{
			Obj: objP, Times: nCount - 1})
	}
	return
}

/// addobj
func (self *Area) AddObj(objP *obj.Obj, nCount int32) *xerror.XError {
	//posT := objP.Pos
	if err, obj1 := self.CanAddObj(objP); err != nil {
		if obj1 != nil {
			// 重新进入地图
			if err.Code() == (int32)(config.ERROR_MAP_PLAYER_ALEARDY) {

				self.DoAddObj1(obj1)
				return err
			}
			if objP.PlayerFlag == 1 {
				go self.TryAddObjAgain(objP, nCount)
			}
		}

		return err
	}
	self.DoAddObj(objP, core.DIRECTION_NONE)
	return nil
}

func (self *Area) DoAddObj(objP *obj.Obj, direction core.DIRECTION) {
	objP.Init(self)
	self.objs.Add(objP.ObjID, objP)
	//log.Debug(self.ToString(), "[objlen:", len(self.objs.GetAll()), "] add Obj"+objP.ToString(), objP.Pos.ToString(),
	//	core.Pos2PPos(objP.GetArea(), objP.Pos))

	// 告诉所有注册玩家 ，当前obj增加
	self.TrySendMsgToRegPlayersByPos(&net.MObjUpdateToc{
		Type:    (int32)(core.ObjUpdateType_Add),
		ObjInfo: []*net.PObj{obj.GenPObj(objP)}}, core.Pos2PPos(objP.GetArea(), objP.Pos))

	if objP.PlayerFlag == 1 {

		if !self.mapPr.(IMap).AddPlayer2Map(objP.ObjID, self) {
			self.DoAddObj1(objP)
		}

		log.Debug("增加player: ", objP.Pos.ToString())

		// 将当前obj注册到周围area
		self.players[objP.ObjID] = objP
		switch direction {
		case core.DIRECTION_NONE:
			self.RegObjToArea(objP, self.GetArea(self.position.X+1, self.position.Y-1))
			self.RegObjToArea(objP, self.GetArea(self.position.X+1, self.position.Y))
			self.RegObjToArea(objP, self.GetArea(self.position.X+1, self.position.Y+1))

			self.RegObjToArea(objP, self.GetArea(self.position.X-1, self.position.Y-1))
			self.RegObjToArea(objP, self.GetArea(self.position.X-1, self.position.Y))
			self.RegObjToArea(objP, self.GetArea(self.position.X-1, self.position.Y+1))

			self.RegObjToArea(objP, self.GetArea(self.position.X, self.position.Y-1))
			self.RegObj(objP)
			self.RegObjToArea(objP, self.GetArea(self.position.X, self.position.Y+1))
		case core.DIRECTION_RIGHT:
			self.RegObjToArea(objP, self.GetArea(self.position.X+1, self.position.Y-1))
			self.RegObjToArea(objP, self.GetArea(self.position.X+1, self.position.Y))
			self.RegObjToArea(objP, self.GetArea(self.position.X+1, self.position.Y+1))
		case core.DIRECTION_LEFT:
			self.RegObjToArea(objP, self.GetArea(self.position.X-1, self.position.Y-1))
			self.RegObjToArea(objP, self.GetArea(self.position.X-1, self.position.Y))
			self.RegObjToArea(objP, self.GetArea(self.position.X-1, self.position.Y+1))
		case core.DIRECTION_UP:
			self.RegObjToArea(objP, self.GetArea(self.position.X-1, self.position.Y+1))
			self.RegObjToArea(objP, self.GetArea(self.position.X, self.position.Y+1))
			self.RegObjToArea(objP, self.GetArea(self.position.X+1, self.position.Y+1))
		case core.DIRECTION_DOWN:
			self.RegObjToArea(objP, self.GetArea(self.position.X-1, self.position.Y-1))
			self.RegObjToArea(objP, self.GetArea(self.position.X, self.position.Y-1))
			self.RegObjToArea(objP, self.GetArea(self.position.X+1, self.position.Y-1))
		}

	}

}

func (self *Area) DoAddObj1(objP *obj.Obj) {
	if objP.PlayerFlag == 1 {
		self.SendMsg(objP, &net.MMapPlayerToc{ObjInfo: obj.GenPObj(objP)})
		// 将周围area的obj信息同步给obj的client
		self.SendMapObjs(objP)
	}
}

//

func (self *Area) DoRemoveObj(objP *obj.Obj, direction core.DIRECTION) {
	self.objs.Delete(objP.ObjID)

	//self.TrySendMsgToRegPlayersByPos(&net.MObjUpdateToc{
	//	Type:    (int32)(core.ObjUpdateType_Del),
	//	ObjInfo: obj.GenPObj(objP)}, objP.Pos)
	if objP.PlayerFlag == 1 {
		switch direction {
		case core.DIRECTION_NONE:
			self.UnRegObjToArea(objP, self.GetArea(self.position.X+1, self.position.Y-1))
			self.UnRegObjToArea(objP, self.GetArea(self.position.X+1, self.position.Y))
			self.UnRegObjToArea(objP, self.GetArea(self.position.X+1, self.position.Y+1))

			self.UnRegObjToArea(objP, self.GetArea(self.position.X-1, self.position.Y-1))
			self.UnRegObjToArea(objP, self.GetArea(self.position.X-1, self.position.Y))
			self.UnRegObjToArea(objP, self.GetArea(self.position.X-1, self.position.Y+1))

			self.UnRegObjToArea(objP, self.GetArea(self.position.X, self.position.Y-1))
			self.UnRegObj(objP)
			self.UnRegObjToArea(objP, self.GetArea(self.position.X, self.position.Y+1))
		case core.DIRECTION_LEFT:
			self.UnRegObjToArea(objP, self.GetArea(self.position.X+1, self.position.Y-1))
			self.UnRegObjToArea(objP, self.GetArea(self.position.X+1, self.position.Y))
			self.UnRegObjToArea(objP, self.GetArea(self.position.X+1, self.position.Y+1))
		case core.DIRECTION_RIGHT:
			self.UnRegObjToArea(objP, self.GetArea(self.position.X-1, self.position.Y-1))
			self.UnRegObjToArea(objP, self.GetArea(self.position.X-1, self.position.Y))
			self.UnRegObjToArea(objP, self.GetArea(self.position.X-1, self.position.Y+1))
		case core.DIRECTION_DOWN:
			self.UnRegObjToArea(objP, self.GetArea(self.position.X-1, self.position.Y+1))
			self.UnRegObjToArea(objP, self.GetArea(self.position.X, self.position.Y+1))
			self.UnRegObjToArea(objP, self.GetArea(self.position.X+1, self.position.Y+1))
		case core.DIRECTION_UP:
			self.UnRegObjToArea(objP, self.GetArea(self.position.X-1, self.position.Y-1))
			self.UnRegObjToArea(objP, self.GetArea(self.position.X, self.position.Y-1))
			self.UnRegObjToArea(objP, self.GetArea(self.position.X+1, self.position.Y-1))
		}

	}
}

func (self *Area) UnRegObjToArea(objP *obj.Obj, area interface{}) {
	if area == nil {
		return
	}
	area.(*Area).UnRegObj(objP)
}

func (self *Area) RegObjToArea(objP *obj.Obj, area interface{}) {
	if area == nil {
		return
	}
	area.(*Area).RegObj(objP)
}

func (self *Area) GetRegObjs() map[int32]*obj.Obj {
	Buff := make(map[int32]*obj.Obj)
	self.regLock.Lock()
	for K, V := range self.regPlayers {
		Buff[K] = V
	}
	self.regLock.Unlock()
	return Buff
}

func (self *Area) RegObj(objP *obj.Obj) {
	self.regLock.Lock()
	self.regPlayers[objP.ObjID] = objP
	self.regLock.Unlock()
}
func (self *Area) UnRegObj(objP *obj.Obj) {
	self.regLock.Lock()
	delete(self.regPlayers, objP.ObjID)
	self.regLock.Unlock()
}

func (self *Area) CanMoveObj(newObj *obj.Obj) (*xerror.XError, *obj.Obj) {
	return self.DoCanAddObj(newObj, true)
}

func (self *Area) CanAddObj(newObj *obj.Obj) (*xerror.XError, *obj.Obj) {
	return self.DoCanAddObj(newObj, false)
}

func (self *Area) DoCanAddObj(newObj *obj.Obj, IsMove bool) (*xerror.XError, *obj.Obj) {
	funT := func() (*xerror.XError, *obj.Obj) {
		Pos := newObj.Pos
		Area := newObj.GetArea()
		newObj.Init(self)
		err, obj := self.CheckObjCollide(newObj)
		newObj.SetArea(Area)
		newObj.Pos = Pos
		return err, obj
	}
	switch newObj.PlayerFlag {
	case 1:
		if !IsMove && self.players[newObj.ObjID] != nil {
			return xerror.New("JoinObj error :already exist player "+newObj.ToString(), (int32)(config.ERROR_MAP_PLAYER_ALEARDY)), self.players[newObj.ObjID]
		}
		return funT()
	default:
		return funT()
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////
/// obj move

func (self *Area) ObjMove(objP *obj.Obj, direction core.DIRECTION) *xerror.XError {
	if direction == core.DIRECTION_STOP {
		if objP.Status == obj.STATUS_MOVE {
			objP.StopMove()
		}
	} else {
		direction := objP.StartMove(direction)
		objP.SetDirection(direction)
		if direction != core.DIRECTION_STOP && direction != core.DIRECTION_NONE {
			objP.Status = obj.STATUS_MOVE
		} else {
			objP.Status = obj.STATUS_NONE
		}
		objP.GetArea().(*Area).TrySendMsgToRegPlayersByPos(&net.MObjUpdateToc{
			Type:    (int32)(core.ObjUpdateType_Upt),
			ObjInfo: []*net.PObj{obj.GenPObj(objP)}}, core.Pos2PPos(objP.GetArea(), objP.Pos))
	}
	return nil
}
