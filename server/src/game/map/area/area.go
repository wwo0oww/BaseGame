package area

import (
	"fmt"
	"game/config"
	"game/core"
	"game/map/obj"
	"lib/cellnet"
	"proto/net"
	"sync"
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
	SendMessage(msg interface{})
	GetArea(int32, int32) interface{}
	AddPlayer2Map(int32, *Area) bool
	DelPlayer4Map(int32)
	GetFrame() uint32
	Test() map[int32]core.Position
	GetAreas() map[int32]map[int32]*Area
}

type AreaMsg struct {
	PlayerID int32
	Msg      interface{}
}

type DataLock struct {
	listGuard sync.Mutex
	listCond  *sync.Cond
}

func (self *DataLock) Lock() {
	self.listGuard.Lock()
}
func (self *DataLock) Unlock() {
	self.listGuard.Unlock()
}

type Area struct {
	players map[int32]*obj.Obj

	position core.Position

	objs     *core.SMap
	objsLock *DataLock

	// 注册了监听的玩家
	regPlayers map[int32]*obj.Obj
	regLock    *DataLock

	// 发送队列
	//sendQueue *cellnet.Pipe

	//接收队列
	recQueue *cellnet.Pipe

	// map reference
	mapPr interface{}

	nLoopIndex int32

	MapFrame uint32
}

func NewArea(mapPr interface{}, position core.Position) *Area {
	reLock := &DataLock{}
	reLock.listCond = sync.NewCond(&reLock.listGuard)

	ObjLock := &DataLock{}
	ObjLock.listCond = sync.NewCond(&ObjLock.listGuard)
	return &Area{
		players:    make(map[int32]*obj.Obj),
		position:   position,
		mapPr:      mapPr,
		objs:       core.NewSMap(),
		regPlayers: make(map[int32]*obj.Obj),
		regLock:    reLock,
		objsLock:   ObjLock,
		//sendQueue: cellnet.NewPipe(),
		recQueue: cellnet.NewPipe()}
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

func (self *Area) GetObjs() map[interface{}]interface{} {
	return self.objs.GetAll()
}

func (self *Area) GetMap() interface{} {
	return self.mapPr
}

func (self *Area) AddMsg(msg interface{}) {
	self.mapPr.(IMap).UnFrozenArea(self)
	self.recQueue.Add(msg)
}

func (self *Area) RelayFrame(nFrame int32) {
	if nFrame == 0 {
		return
	}
	for _, item := range self.objs.GetAll() {
		item.(*obj.Obj).AddRelayFrame(nFrame)
	}
}

func (self *Area) LoopRecMsg() {
	var writeList []interface{}
	self.recQueue.Get(&writeList, 20)

	for _, msg := range writeList {
		self.DealRecMsg(msg)
	}

}

func (self *Area) GetMapFrame() uint32 {
	return self.MapFrame
}

func (self *Area) Running() {
	bActiv := false
	for _, item := range self.objs.GetAll() {
		item.(*obj.Obj).Update()
		item.(*obj.Obj).FrameEnd()
		if item.(*obj.Obj).Status != obj.STATUS_NONE {
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

func (self *Area) G_GetLeftAreas() []interface{} {
	var areas []interface{}
	area1 := self.GetArea(self.position.X-1, self.position.Y)
	if area1 == nil {
		return areas
	}
	areas = append(areas, area1)

	area2 := self.GetArea(self.position.X-1, self.position.Y-1)
	if area2 != nil {
		areas = append(areas, area2)
	}

	area3 := self.GetArea(self.position.X-1, self.position.Y+1)
	if area3 != nil {
		areas = append(areas, area3)
	}

	return areas
}

func (self *Area) G_GetRightAreas() []interface{} {
	var areas []interface{}
	area1 := self.GetArea(self.position.X+1, self.position.Y)
	if area1 == nil {
		return areas
	}
	areas = append(areas, area1)

	area2 := self.GetArea(self.position.X+1, self.position.Y-1)
	if area2 != nil {
		areas = append(areas, area2)
	}

	area3 := self.GetArea(self.position.X+1, self.position.Y+1)
	if area3 != nil {
		areas = append(areas, area3)
	}

	return areas
}
func (self *Area) G_GetUpAreas() []interface{} {
	var areas []interface{}
	area1 := self.GetArea(self.position.X, self.position.Y+1)
	if area1 == nil {
		return areas
	}
	areas = append(areas, area1)

	area2 := self.GetArea(self.position.X-1, self.position.Y+1)
	if area2 != nil {
		areas = append(areas, area2)
	}

	area3 := self.GetArea(self.position.X+1, self.position.Y+1)
	if area3 != nil {
		areas = append(areas, area3)
	}

	return areas
}
func (self *Area) G_GetDownAreas() []interface{} {
	var areas []interface{}
	area1 := self.GetArea(self.position.X, self.position.Y-1)
	if area1 == nil {
		return areas
	}
	areas = append(areas, area1)

	area2 := self.GetArea(self.position.X-1, self.position.Y-1)
	if area2 != nil {
		areas = append(areas, area2)
	}

	area3 := self.GetArea(self.position.X+1, self.position.Y-1)
	if area3 != nil {
		areas = append(areas, area3)
	}

	return areas
}

func (self *Area) GetRightArea() *Area {
	if area := self.mapPr.(IMap).GetArea(self.position.X+1, self.position.Y); area != nil {
		return area.(*Area)
	}
	return nil
}

func (self *Area) GetLeftArea() *Area {
	if area := self.mapPr.(IMap).GetArea(self.position.X-1, self.position.Y); area != nil {
		return area.(*Area)
	}
	return nil
}

func (self *Area) GetUpArea() *Area {
	if area := self.mapPr.(IMap).GetArea(self.position.X, self.position.Y+1); area != nil {
		return area.(*Area)
	}
	return nil
}

func (self *Area) GetArea(x, y int32) interface{} {
	return self.mapPr.(IMap).GetArea(x, y)
}

func (self *Area) GetDownArea() *Area {
	if area := self.mapPr.(IMap).GetArea(self.position.X, self.position.Y-1); area != nil {
		return area.(*Area)
	}
	return nil
}

func (self *Area) SetLoopIndex(nIndex int32) {
	self.nLoopIndex = nIndex
}

func (self *Area) GetLoopIndex() int32 {
	return self.nLoopIndex
}

func (self *Area) SendMapObjs(objP *obj.Obj) {
	go self.DoSendMapObjs(objP)
}

func (self *Area) DoSendMapObjs(objP *obj.Obj) {
	// 将周围area的obj信息同步给obj的client
	var objs []*net.PObj
	X, Y := config.MapShowSize()
	Iterator := func(areaP interface{}) {
		if areaP == nil {
			return
		}
		area := areaP.(*Area)
		if area != nil {
			for _, item := range area.objs.GetAll() {
				item1 := item.(*obj.Obj)
				if item1.ObjID != objP.ObjID {
					if core.AbsInt32(item1.Pos.X, objP.Pos.X) < X && core.AbsInt32(item1.Pos.Y, objP.Pos.Y) < Y {
						objs = append(objs, obj.GenPObj(item1))
						if len(objs) > 40 {
							var buf []*net.PObj
							for _, objT := range objs {
								buf = append(buf, objT)
							}
							self.SendMsg(objP, &net.MMapInfoToc{ObjInfo: buf, FrameCount: self.mapPr.(IMap).GetFrame()})
							objs = objs[0:0]
						}
					}
				}
			}
		}
	}

	Iterator(self.GetArea(self.position.X+1, self.position.Y-1))
	Iterator(self.GetArea(self.position.X+1, self.position.Y))
	Iterator(self.GetArea(self.position.X+1, self.position.Y+1))

	Iterator(self.GetArea(self.position.X-1, self.position.Y-1))
	Iterator(self.GetArea(self.position.X-1, self.position.Y))
	Iterator(self.GetArea(self.position.X-1, self.position.Y+1))

	Iterator(self.GetArea(self.position.X, self.position.Y+1))
	Iterator(self)
	Iterator(self.GetArea(self.position.X, self.position.Y-1))

	self.SendMsg(objP, &net.MMapInfoToc{ObjInfo: objs, FrameCount: self.mapPr.(IMap).GetFrame()})
}
