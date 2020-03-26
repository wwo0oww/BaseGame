package area

import (
	"game/config"
	mMsg "game/map/msg"
	"game/map/obj"
	xerror "lib/error"
	"lib/log"
	"proto/net"
)

func (self *Area) JoinObj(newObj *obj.Obj, nCount int32) *xerror.XError {
	return self.AddObj(newObj, nCount)
}

func (self *Area) TrySendMsgToRegPlayersByPos(msg interface{}, position *net.PPosition) {
	//X, Y := config.MapShowSize()
	//for _, obj := range self.GetRegObjs() {
	//	pos := core.Pos2PPos(obj.GetArea(), obj.Pos)
	//	if core.AbsInt32(pos.X, position.X) < X && core.AbsInt32(pos.Y, position.Y) < Y {
	//		self.mapPr.(IMap).SendMessage(&AreaMsg{obj.ObjID, msg})
	//	}
	//}
	for _, obj := range self.GetRegObjs() {
		self.mapPr.(IMap).SendMessage(&AreaMsg{obj.ObjID, msg})
	}
}

func (self *Area) SendMsg(objP *obj.Obj, msg interface{}) {
	self.mapPr.(IMap).SendMessage(&AreaMsg{objP.ObjID, msg})
}

func (self *Area) DealRecMsg(msg interface{}) {
	var err *xerror.XError
	switch msg.(type) {
	case *mMsg.JoinObj:
		err = self.JoinObj(msg.(*mMsg.JoinObj).Obj, msg.(*mMsg.JoinObj).Times)
	case *mMsg.ObjMove:
		objT := self.objs.Get(msg.(*mMsg.ObjMove).Obj.ObjID)
		err = self.ObjMove(objT.(*obj.Obj), msg.(*mMsg.ObjMove).Direction)
	default:
		err = xerror.New("not define msg"+msg.(mMsg.IMapMsg).ToString(), (int32)(config.ERROR_COMMON))
	}
	if err != nil {
		log.Debug(err.Error())
	}
}
