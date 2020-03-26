package cellnet

import (
	"sync"
)

// 不限制大小，添加不发生阻塞，接收阻塞等待
type Pipe struct {
	list      []interface{}
	listGuard sync.Mutex
	listCond  *sync.Cond
}

func (self *Pipe) GetList() []interface{} {
	self.listGuard.Lock()
	var buf []interface{}
	for _, item := range self.list {
		buf = append(buf, item)
	}
	self.listGuard.Unlock()
	return buf
}

// 添加时不会发送阻塞
func (self *Pipe) Add(msg interface{}) {
	self.listGuard.Lock()
	self.list = append(self.list, msg)
	self.listGuard.Unlock()

	self.listCond.Signal()
}

func (self *Pipe) Reset() {
	self.list = self.list[0:0]
}

// 如果没有数据，发生阻塞
func (self *Pipe) Pick(retList *[]interface{}) (exit bool) {
	return self.DoPick(retList, true, 0)
}

// 如果没有数据，不会发生阻塞
func (self *Pipe) All(retList *[]interface{}) (exit bool) {
	return self.DoPick(retList, false, 0)
}

func (self *Pipe) Get(retList *[]interface{}, nNum int32) (exit bool) {
	return self.DoPick(retList, false, nNum)
}

func (self *Pipe) DoPick(retList *[]interface{}, IsLock bool, nNum int32) (exit bool) {

	self.listGuard.Lock()

	if IsLock {
		for len(self.list) == 0 { // 加for =>解除阻塞后，立马在检查长度
			self.listCond.Wait()
		}
	}

	self.listGuard.Unlock()

	self.listGuard.Lock()

	// 复制出队列
	var nCount int32 = 0
	for _, data := range self.list {
		if data == nil {
			exit = true
			break
		} else {
			*retList = append(*retList, data)
		}
		nCount++
		if nNum != 0 && nCount >= nNum {
			break
		}
	}

	self.list = self.list[nCount:]
	self.listGuard.Unlock()

	return
}

func NewPipe() *Pipe {
	self := &Pipe{}
	self.listCond = sync.NewCond(&self.listGuard)

	return self
}
