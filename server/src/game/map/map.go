package emap

import (
	"fmt"
	"game/config"
	"game/core"
	"game/map/area"
	mMsg "game/map/msg"
	"game/map/obj"
	"lib/cellnet"
	"lib/log"
	"math/rand"
	"proto"
	"sync"
	"sync/atomic"
	"time"
)

const (
	STATUS_EXIT int32 = iota
	STATUS_PREPARE
	STATUS_RUNING
)
const (
	LOOP_LOCK int32 = iota
	LOOP_UNLOCK
	LOOP_CLOSE
)

var m_map *Map = nil

type WaitStatusType int32

const (
	WAITSTATUSTYPE_NOING int32 = iota
	WAITSTATUSTYPE_DONE
	WAITSTATUSTYPE_FROZEN
)

type IPlayer interface {
	SendMsgToMap(Msg interface{})
	GetPlayerID() int32
}

type LoopSync struct {
	WaitStatus     int32
	MapFrame       uint32
	WaitRelayFrame int32
	WaitCond       *sync.Cond
	WaitGurad      sync.Mutex
}

type Map struct {
	// 玩家到地图映射
	PlayerMapping *core.SMap

	players [] *obj.Obj
	// 发送队列
	sendQueue *cellnet.Pipe

	//接收队列
	recQueue *cellnet.Pipe

	// 退出同步器
	exitSync sync.WaitGroup

	// map状态
	status int32

	// 包含的所有area
	areas map[int32]map[int32]*area.Area

	// Loop Lock
	loopCond  *sync.Cond
	loopGurad sync.Mutex
	//
	loopSync map[int32]*LoopSync
	//
	loopGroup sync.WaitGroup

	// 当前帧数
	nFrameCount uint32
}

func GetMapInstance() *Map {
	if m_map == nil {
		newMap()
	}
	return m_map
}

func newMap() *Map {
	m_map = &Map{
		sendQueue:     cellnet.NewPipe(),
		recQueue:      cellnet.NewPipe(),
		PlayerMapping: core.NewSMap()}
	m_map.Start()
	mMsg.Map = m_map
	return m_map
}

func (self *Map) ToString() string {
	return "map"
}

func (self *Map) GetAreas() map[int32]map[int32]*area.Area {
	return self.areas
}

func (self *Map) GetPlayer2Map() *core.SMap {
	return self.PlayerMapping
}

func (self *Map) AddPlayer2Map(ObjID int32, area *area.Area) (BHas bool) {
	if self.PlayerMapping.Get(ObjID) == nil {
		BHas = false
	} else {
		BHas = true
	}
	self.PlayerMapping.Add(ObjID, area)
	return
}

func (self *Map) DelPlayer4Map(ObjID int32) {
	//todo 因为player总是从一个area移动到另外一个area,目前 不错处理【以后出现player不积极等情况，从地图移除再做处理】
}

func (self *Map) sendLoop() {
	self.WaitForRunning()
	var writeList []interface{}

	for {
		writeList = writeList[0:0]
		exit := self.sendQueue.Pick(&writeList)

		// 遍历要发送的数据
		for _, msg := range writeList {
			self.SendMsgToMap(msg)
		}

		if exit {
			break
		}
	}

	// 通知完成
	self.exitSync.Done()
}

func (self *Map) recvLoop() {
	self.WaitForRunning()
	var writeList []interface{}

	for {
		writeList = writeList[0:0]
		exit := self.recQueue.Pick(&writeList)
		// 遍历要处理的数据
		for _, msg := range writeList {
			self.DispatchMsg(msg)
		}
		if exit {
			break
		}
	}

	// 通知完成
	self.exitSync.Done()
}

func (self *Map) WaitForRunning() {
	for {
		if self.status == STATUS_RUNING {
			break
		}
		time.Sleep(100000)
	}
}

func (self *Map) Start() {
	atomic.StoreInt32(&self.status, STATUS_PREPARE)
	self.loopCond = sync.NewCond(&self.loopGurad)
	self.recQueue.Reset()
	self.sendQueue.Reset()

	self.exitSync.Add(2)

	go func() {

		self.exitSync.Wait()

		self.Close()
	}()

	go self.sendLoop()
	go self.recvLoop()

	self.InitArea()

	go func() {
		for {
			time.Sleep(time.Second * 20)
			nCount := 0
			for i, _ := range self.areas {
				for j, _ := range self.areas[i] {
					len := len(self.areas[i][j].GetObjs())
					nCount += len
					if len != 0 {
						log.Debug(self.areas[i][j].ToString(), ":", len)
					}
				}
			}
			log.Debug(self.ToString(), " total obj:", nCount)
		}
	}()
}

func (self *Map) FrameAdd() {
	atomic.AddUint32(&self.nFrameCount, 1)
}

func (self *Map) GetFrame() uint32 {
	return atomic.LoadUint32(&self.nFrameCount)
}

func (self *Map) Close() {
	//for i,j:=range self.areas{
	//	self.areas[i][j] = nil
	//}
	//todo
}

func (self *Map) GetRecQueue() *cellnet.Pipe {
	return self.recQueue
}

func (self *Map) AddMessage(data interface{}) {
	self.recQueue.Add(data)
}

func (self *Map) SendMessage(msg interface{}) {
	self.sendQueue.Add(msg)
}

func (self *Map) DispatchMsg(raw interface{}) {
	msg := proto.DecodeMapMsg(raw)
	if msg == nil {
		return
	}
	pos := msg.(mMsg.IMapMsg).GetAreaPos()

	if self.areas[pos.X][pos.Y] == nil {
		return
	}
	self.areas[pos.X][pos.Y].AddMsg(msg)
}

func (self *Map) GetPlayer() {}

func (self *Map) GetArea(x, y int32) interface{} {
	W, H := config.GetMapWH()
	if x >= W || y >= H || x < 0 || y < 0 {
		return nil
	}
	return self.areas[x][y]
}
func (self *Map) Test() map[int32]core.Position {
	return testmap
}

var testmap map[int32]core.Position

func Test() {
	testmap = make(map[int32]core.Position)
	//go func() {
	//	var nCount int32
	//	for {
	//		nCount++
	//		time.Sleep(time.Second)
	//		println("do_",nCount)
	//	}
	//}()
	//go func() {
	//	for {
	//		time.Sleep(time.Second)
	//		println(len(GetMapInstance().GetRecQueue().GetList()))
	//	}
	//}()
	w, h := config.GetMapWH()
	pos := core.Position{X: config.GetAreaSize() * w / 2, Y: config.GetAreaSize() * h / 2}
	//var i, j int32
	//var size, size1 int32
	//size = 5
	//size1 = size/2 + 1
	//for i = 0; i < (config.GetAreaSize()*w/size)-1; i++ {
	//	for j = 0; j < (config.GetAreaSize()*h/size)-1; j++ {
	//		if i != 0 && j != 0 {
	//			continue
	//		}
	//		pos1 := core.Position{X: i*size + size1, Y: j*size + size1}
	//		Id := obj.GetMaxObj()
	//		NewObj := &obj.Obj{
	//			ObjID: Id,
	//			Type:  obj.OBJTYPE_NPC,
	//			Size:  core.Position{X: size, Y: size},
	//			Pos:   pos1}
	//		GetMapInstance().Test()[Id] = pos1
	//		NewObj.SetSpeed(1)
	//		GetMapInstance().AddMessage(&mMsg.JoinObj{
	//			Obj:   NewObj,
	//			Times: 10})
	//
	//	}
	//}
	for i := 0; i < 1; i++ {
		var j int32
		for j = -1; j < 1; j++ {
			pos1 := core.Position{X: pos.X + 5, Y: pos.Y + j*5}
			Id := obj.GetMaxObj()
			NewObj := &obj.Obj{
				ObjID: Id,
				Type:  obj.T_OBJ.NPC().TEST().Value(),
				Name:  fmt.Sprintf("路人%d", Id),
				Size:  core.Position{X: 5, Y: 5},
				Pos:   pos1}
			GetMapInstance().Test()[Id] = pos1
			NewObj.SetDirection(core.Directions[rand.Intn(4)])
			NewObj.SetSpeed(1)
			GetMapInstance().AddMessage(&mMsg.JoinObj{
				Obj:   NewObj,
				Times: 10})
		}
		time.Sleep(time.Second * 3)
	}
}

func (self *Map) LoopSyncSafe(fun func(loopSync map[int32]*LoopSync)) {
	self.loopGurad.Lock()
	fun(self.loopSync)
	self.loopGurad.Unlock()
}

func init() {
	GetMapInstance()
	go func() {
		time.Sleep(time.Second)
		Test()
	}()
}
