package core

import (
	"sync"
)

type SMap struct {
	Data      map[interface{}]interface{}
	listGuard sync.Mutex
	listCond  *sync.Cond
}

func (self *SMap) GetAll() map[interface{}]interface{} {
	self.listGuard.Lock()
	buf := make(map[interface{}]interface{})
	for K, V := range self.Data {
		buf[K] = V
	}
	self.listGuard.Unlock()
	return buf
}

func (self *SMap) Add(k interface{}, v interface{}) {
	self.listGuard.Lock()
	self.Data[k] = v
	self.listGuard.Unlock()

	self.listCond.Signal()
}

func (self *SMap) Get(k interface{}) interface{} {
	self.listGuard.Lock()
	v := self.Data[k]
	self.listGuard.Unlock()

	return v
}

func (self *SMap) Delete(k interface{}) {
	self.listGuard.Lock()
	delete(self.Data, k)
	self.listGuard.Unlock()
}

func NewSMap() *SMap {
	self := &SMap{}
	self.listCond = sync.NewCond(&self.listGuard)
	self.Data = make(map[interface{}]interface{})
	return self
}
