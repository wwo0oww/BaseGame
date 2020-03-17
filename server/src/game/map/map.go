package emap

import (
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

type LoopSync struct {
	WaitStatus     int32
	WaitRelayFrame int32
	WaitCond       *sync.Cond
	WaitGurad      sync.Mutex
}

type Map struct {
	// 玩家到地图映射
	PlayerMapping map[string]interface{}

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
}

func GetMapInstance() *Map {
	if m_map == nil {
		newMap()
	}
	return m_map
}

func newMap() *Map {
	m_map = &Map{sendQueue: cellnet.NewPipe(), recQueue: cellnet.NewPipe()}
	m_map.Start()
	return m_map
}

func (self *Map) ToString() string {
	return "map"
}

func (self *Map) sendLoop() {
	self.WaitForRunning()
	var writeList []interface{}

	for {
		writeList = writeList[0:0]
		exit := self.sendQueue.Pick(&writeList)

		// 遍历要发送的数据
		for _, msg := range writeList {
			self.SendMessage(msg)
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
		/// 将map分为多个区域，根据具体map的长W 宽H 分配具体的层数，如单位区域的长宽为4*4，每次下面只有4个,总大小为64*64，则有
		/// 4。根据player进程玩家的position找到具体的，map位置
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
			time.Sleep(time.Second * 1)
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
	//todo
}

func (self *Map) DispatchMsg(raw interface{}) {
	msg := proto.DecodeMapMsg(raw)
	if msg == nil {
		return
	}
	pos := msg.(mMsg.IMapMsg).GetAreaPos()
	x, y := config.GetMapWH()
	if pos.X >= x || pos.Y >= y {
		log.Debug("map to area error ", pos.X, pos.Y)
		return
	}
	self.areas[pos.X][pos.Y].AddMsg(msg)
}

func (self *Map) GetAreas() map[int32]map[int32]*area.Area {
	return self.areas
}

func Test() {
	go func() {
		var nCount int32
		for {
			nCount++
			time.Sleep(time.Second)
			println("do_",nCount)
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second)
			println(len(GetMapInstance().GetRecQueue().GetList()))
		}
	}()
	for i := 0; i < 100; i++ {
		log.Debug("add 500")
		for j := 0; j < 100; j++ {
			NewObj := &obj.Obj{
				ObjID: obj.GetMaxObj(),
				Size:  core.Position{X: 5, Y: 5},
				Pos:   core.Position{X: int32(i), Y: int32(i)}}
			NewObj.SetDirection(core.Directions[rand.Intn(4)])
			NewObj.SetSpeed(2)
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
		log.Debug("test")
		Test()
	}()
}
