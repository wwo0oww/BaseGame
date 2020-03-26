package main

import (
	"sync"
	"sync/atomic"
	"time"
)

func test11() {
	var arr = []int32{0, 0}
	go func() {
		var i int32 = 0
		for i < 10 {
			println("t1")
			atomic.StoreInt32(&arr[1], arr[1]+1)
			i++
		}
		println(arr[1])
	}()
	go func() {
		var i int32 = 0
		for i < 10 {
			println("t2")
			atomic.StoreInt32(&arr[1], arr[1]+1)
			i++
		}
		println(arr[1])
	}()
	for {
		time.Sleep(1)
	}
}

func main() {
	var cond *sync.Cond
	var gurad sync.Mutex
	cond = sync.NewCond(&gurad)
	go func() {
		for {
			//gurad.Lock()
			time.Sleep(time.Second)
			cond.Signal()
			//gurad.Unlock()
		}
	}()
	go func() {
		for {
			gurad.Lock()
			cond.Wait()
			gurad.Unlock()
			println("x")
		}
	}()
	for {
		time.Sleep(1)
	}
}

func test1() {
	var cond *sync.Cond
	var gurad sync.Mutex
	var loopSync sync.WaitGroup
	type tt struct {
		Name string
	}
	var m map[int32]*tt = make(map[int32]*tt)
	m[1] = &tt{Name: "hhh"}
	cond = sync.NewCond(&gurad)
	go func() {
		gurad.Lock()
		cond.Wait()
		println("do sth")
		loopSync.Done()
		gurad.Unlock()
	}()
	go func() {
		//cond.Signal()
		for i := 0; i < 1000; i++ {
			//gurad.Lock()
			m[1].Name = "tt1"
			//gurad.Unlock()
			time.Sleep(10)
		}
		println(m[1].Name)
	}()
	go func() {
		//cond.Signal()
		for i := 0; i < 1000; i++ {
			//gurad.Lock()
			m[1].Name = "tt2"
			//gurad.Unlock()
			time.Sleep(10)
		}
		println(m[1].Name)
		cond.Signal()
	}()
	loopSync.Add(1)
	go func() {
		loopSync.Wait()
		println("t3")
	}()
	for {
		time.Sleep(1)
	}
}

func test() {
	var i int32 = 0
	t1 := time.Now().Nanosecond()
	go func() {
		for {
			atomic.StoreInt32(&i, i+1)
			time.Sleep(time.Second / 10)
		}
	}()
	go func() {
		for {
			if i%10 == 0 {
				t2 := time.Now().Nanosecond()
				println("", i, t2-t1)
				t1 = t2
			}
		}
	}()
	for i < 1000 {
		time.Sleep(time.Second / 100)
	}

}
