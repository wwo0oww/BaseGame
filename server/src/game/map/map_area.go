package emap

import (
	"fmt"
	"game/config"
	"game/core"
	"game/map/area"
	"lib/log"
	"sync"
	"sync/atomic"
	"time"
)

func (self *Map) InitArea() {
	self.areas = make(map[int32]map[int32]*area.Area)
	W, H := config.GetMapWH()
	var i, j int32
	for i = 0; i < W; i++ {
		self.areas[i] = make(map[int32]*area.Area)
		for j = 0; j < H; j++ {
			newArea := area.NewArea(self, core.Position{X: i, Y: j})
			self.areas[i][j] = newArea
			newArea.Start()
		}
	}
	go self.loopArea()
}

func (self *Map) loopArea() {
	self.initLoop()
	go self.startLoop()
	atomic.StoreInt32(&self.status, STATUS_RUNING)
}

func (self *Map) initLoop() {
	SingleAreaNum := config.GetSingleThreadAreaNum()
	self.loopSync = make(map[int32]*LoopSync)
	var buff []*area.Area
	var iCount int = 0
	var nCount int32 = 0
	self.loopGroup.Add(1)
	for i, _ := range self.areas {
		for j, _ := range self.areas[i] {
			buff = append(buff, self.areas[i][j])
			iCount++
			if iCount == SingleAreaNum {
				iCount = 0
				var buffT []*area.Area = make([]*area.Area, SingleAreaNum, SingleAreaNum)
				copy(buffT, buff[0:SingleAreaNum])
				self.exitSync.Add(1)
				loopSync := &LoopSync{}
				loopSync.WaitStatus = WAITSTATUSTYPE_DONE
				loopSync.WaitCond = sync.NewCond(&loopSync.WaitGurad)
				self.loopSync[nCount] = loopSync
				for _, item := range buffT {
					item.SetLoopIndex(nCount)
				}
				self.loopGroup.Add(1)
				go self.doLoopArea(buffT, self.loopSync[nCount], nCount)
				nCount++
				buff = buff[0:0]
			}

		}
	}
	self.loopGroup.Done()
}

func (self *Map) startLoop() {
	var a []int32
	self.loopGroup.Wait()
	var t1, t2, t3 int
	t1 = 0
	t2 = 0
	nFrame := config.GetMapFrame()
	for {
		self.FrameAdd()

		a = a[0:0]
		t1 = time.Now().Nanosecond()
		if t1-t2 == 0 {
			t3 = (int)(time.Second) / nFrame
		} else {
			t3 = (int)(time.Second)/nFrame - (t1 - t2)
		}
		if t1-t2 != 0 {
			log.Debug("xx:", t1-t2)
		}
		if n := self.GetFrame(); n%10 == 0 {

			//log.Debug("frame ",n)
		}
		time.Sleep((time.Duration)(t3))
		t2 = time.Now().Nanosecond()
		for index, item := range self.loopSync {
			item.MapFrame = self.GetFrame()
			if item.WaitStatus == WAITSTATUSTYPE_DONE {
				item.WaitCond.Signal()
			} else { // 进入这个分支有2种情况 1.被唤醒B事件未完成， 2.事件完成了WaitStatus 已经变为了 WAITSTATUSTYPE_DONE，但是还未进入等待状态，概率极小
				if item.WaitStatus != WAITSTATUSTYPE_FROZEN {
					item.WaitCond.Signal() //处理情况2：当前B，被唤醒B： 处理 唤醒A时，A未处于等待状态
					a = append(a, index)
					log.Debug("no finished %d", index)
					item.WaitRelayFrame++
				}
			}
		}
		var s string = ""
		for _, item := range a {
			s = fmt.Sprintf("%s %d", s, item)
		}
		if s != "" {
			log.Debug("lose", s)
		}
	}
}

func (self *Map) FrozenArea(areaP *area.Area) {
	atomic.StoreInt32(&self.loopSync[areaP.GetLoopIndex()].WaitStatus, WAITSTATUSTYPE_FROZEN)

}

func (self *Map) UnFrozenArea(areaP *area.Area) {
	if self.loopSync[areaP.GetLoopIndex()].WaitStatus == WAITSTATUSTYPE_FROZEN {
		atomic.StoreInt32(&self.loopSync[areaP.GetLoopIndex()].WaitStatus, WAITSTATUSTYPE_DONE)
	}
}

func (self *Map) doLoopArea(areas []*area.Area, loopSync *LoopSync, nCount int32) {
	self.loopGroup.Done()
	for {
		loopSync.WaitGurad.Lock()
		loopSync.WaitCond.Wait()
		loopSync.WaitGurad.Unlock()
		atomic.StoreInt32(&loopSync.WaitStatus, WAITSTATUSTYPE_NOING)
		if self.status == STATUS_EXIT {
			self.exitSync.Done()
			return
		}
		for _, area := range areas {
			area.MapFrame = loopSync.MapFrame
			area.RelayFrame(loopSync.WaitRelayFrame)
			area.LoopRecMsg()
			area.Running()
			area.LoopSendMsg()
		}
		loopSync.WaitRelayFrame = 0
		atomic.StoreInt32(&loopSync.WaitStatus, WAITSTATUSTYPE_DONE)
	}

}
