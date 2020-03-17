package area

import (
	"errors"
	"fmt"
	"game/config"
	"game/core"
	mMsg "game/map/msg"
	"game/map/obj"
	"lib/cellnet"
	"lib/log"
)

//type AreaObj stru {
//	//// 获取坐标
//	//GetPos() core.Position
//	//
//	//// 获取宽高
//	//GetWH() core.Position
//	//
//	//// 获取对象名称
//	//ToString() string
//	//
//	//// 获取类型
//	//
//	obj.Obj
//}

type ContextSet struct {
	nData map[string]int32
}

type IMap interface {
	FrozenArea(*Area)
	UnFrozenArea(*Area)
}

type Area struct {
	players map[int32]*obj.Obj

	position core.Position

	objs map[int32]*obj.Obj

	// 发送队列
	sendQueue *cellnet.Pipe

	//接收队列
	recQueue *cellnet.Pipe

	// map reference
	mapPr interface{}

	nLoopIndex int32
}

func NewArea(mapPr interface{}, position core.Position) *Area {
	return &Area{
		players:   make(map[int32]*obj.Obj),
		position:  position,
		mapPr:     mapPr,
		objs:      make(map[int32]*obj.Obj),
		sendQueue: cellnet.NewPipe(),
		recQueue:  cellnet.NewPipe()}
}

func (self *Area) Start() {
	//todo
}

func (self *Area) ToString() string {
	return fmt.Sprintf("area_%d_%d", self.position.X, self.position.Y)
}

func (self *Area) GetPos() core.Position {
	return self.position
}

func (self *Area) GetObjs() map[int32]*obj.Obj {
	return self.objs
}

func (self *Area) AddMsg(msg interface{}) {
	self.mapPr.(IMap).UnFrozenArea(self)
	self.recQueue.Add(msg)
}

func (self *Area) RelayFrame(nFrame int32) {
	if nFrame == 0 {
		return
	}
	for _, obj := range self.objs {
		obj.AddRelayFrame(nFrame)
	}
}

func (self *Area) LoopRecMsg() {
	var writeList []interface{}
	self.recQueue.Get(&writeList, 20)
	for _, msg := range writeList {
		self.DealRecMsg(msg)
	}

}

func (self *Area) DealRecMsg(msg interface{}) {
	var err error
	switch msg.(type) {
	case *mMsg.JoinObj:
		err = self.JoinObj(msg.(*mMsg.JoinObj).Obj, msg.(*mMsg.JoinObj).Times)
	default:
		err = errors.New("not define msg" + msg.(mMsg.IMapMsg).ToString())
	}
	if err != nil {
		log.Debug(err.Error())
	}
}

func (self *Area) Running() {
	bActiv := false
	for _, item := range self.objs {
		item.Update()
		item.FrameEnd()
		if item.Status != obj.STATUS_NONE {
			bActiv = true
		}
	}
	if !bActiv {
		self.mapPr.(IMap).FrozenArea(self)
	}
}

func (self *Area) LoopSendMsg() {
	// todo
}

func (self *Area) G_GetRightArea() interface{} {
	area := self.GetRightArea()
	if area == nil { // 之所以這麽寫是因爲 將一個nil值给 *Area，会变成(0x0,0x0)的格式，在一次将(0x0,0x0)，可以等于nil,赋值给interface{}，会变成
		// (0x8dd020,0x0)，不能等于nil,另外 if area := self.GetLeftArea()；area ==nil ，这种写法也会出问题，area即使等于 (0x0,0x0) 也会出现 area !=nil
		return nil
	}
	return area
}

func (self *Area) GetRightArea() *Area {
	x, _ := config.GetMapWH()
	if self.position.X+1 < x {
		areas := self.mapPr.(interface{ GetAreas() map[int32]map[int32]*Area }).GetAreas()
		return areas[self.position.X+1][self.position.Y]
	}
	return nil
}

func (self *Area) G_GetLeftArea() interface{} {
	area := self.GetLeftArea()
	if area == nil { // 之所以這麽寫是因爲 將一個nil值给 *Area，会变成(0x0,0x0)的格式，在一次将(0x0,0x0)，可以等于nil,赋值给interface{}，会变成
		// (0x8dd020,0x0)，不能等于nil,另外 if area := self.GetLeftArea()；area ==nil ，这种写法也会出问题，area即使等于 (0x0,0x0) 也会出现 area !=nil
		return nil
	}
	return area
}

func (self *Area) GetLeftArea() *Area {
	if self.position.X-1 >= 0 {
		areas := self.mapPr.(interface{ GetAreas() map[int32]map[int32]*Area }).GetAreas()
		return areas[self.position.X-1][self.position.Y]
	}
	return nil
}

func (self *Area) G_GetUpArea() interface{} {
	area := self.GetUpArea()
	if area == nil { // 之所以這麽寫是因爲 將一個nil值给 *Area，会变成(0x0,0x0)的格式，在一次将(0x0,0x0)，可以等于nil,赋值给interface{}，会变成
		// (0x8dd020,0x0)，不能等于nil,另外 if area := self.GetLeftArea()；area ==nil ，这种写法也会出问题，area即使等于 (0x0,0x0) 也会出现 area !=nil
		return nil
	}
	return area
}

func (self *Area) GetUpArea() *Area {
	_, y := config.GetMapWH()
	if self.position.Y+1 < y {
		areas := self.mapPr.(interface{ GetAreas() map[int32]map[int32]*Area }).GetAreas()
		return areas[self.position.X][self.position.Y+1]
	}
	return nil
}

func (self *Area) G_GetDownArea() interface{} {
	area := self.GetDownArea()
	if area == nil { // 之所以這麽寫是因爲 將一個nil值给 *Area，会变成(0x0,0x0)的格式，在一次将(0x0,0x0)，可以等于nil,赋值给interface{}，会变成
		// (0x8dd020,0x0)，不能等于nil,另外 if area := self.GetLeftArea()；area ==nil ，这种写法也会出问题，area即使等于 (0x0,0x0) 也会出现 area !=nil
		return nil
	}
	return area
}

func (self *Area) GetDownArea() *Area {
	if self.position.Y-1 >= 0 {
		areas := self.mapPr.(interface{ GetAreas() map[int32]map[int32]*Area }).GetAreas()
		return areas[self.position.X][self.position.Y-1]
	}
	return nil
}

func (self *Area) SetLoopIndex(nIndex int32) {
	self.nLoopIndex = nIndex
}
func (self *Area) GetLoopIndex() int32 {
	return self.nLoopIndex
}
