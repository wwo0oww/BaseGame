package main

import (
	"game/core"
	"time"
)

type itt interface {
	GToString() interface{}
	ToString() *tt
}

type tt struct {
	Name string
}

func (self *tt) ToString() *tt {
	return nil
}

func (self *tt) GToString() interface{} {
	return self.ToString()
}


func main() {
	var i int32
	i = 0
	go func() {
		for {
			time.Sleep(time.Second / 10)
			println("do_", i)
			i++
		}
	}()
	for {
		time.Sleep(time.Second / 10)
	}
}

func test2() {
	p1 := core.Position{X: 1, Y: 1}
	p2 := core.Position{X: 2, Y: 2}
	p2 = p1
	p2.X = 3
	println(p2.ToString())
	println(p1.ToString())
}

func test11() {
	t := &tt{Name: "xxx"}
	test1(nil)
	test(t)
}

func test1(t interface{}) {
	println(t == nil)
}

func test(t interface{}) {
	println(t.(itt).GToString() == (interface{})(nil))
	println(t.(itt).ToString())
}
