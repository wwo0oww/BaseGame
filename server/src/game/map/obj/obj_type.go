package obj

import "game/map/obj/npc"

const LevelNum int32 = 1000

var T_OBJ *T_Obj

type T_Obj struct {
	value    int32
	npc      *T_NPC
	building *T_BUILDING
}

func (self *T_Obj) Value() int32 {
	if self.value == 0 {
		self.value = 1
	}
	return self.value
}

func (self *T_Obj) NPC() *T_NPC {
	if self.npc == nil {
		self.npc = &T_NPC{base: self}
	}
	return self.npc
}

func (self *T_Obj) BUILDING() *T_BUILDING {
	if self.building == nil {
		self.building = &T_BUILDING{base: self}
	}
	return self.building
}

type T_NPC struct {
	base   *T_Obj
	value  int32
	player *T_NPC_PLAYER
	test   *T_NPC_TEST
}

func (self *T_NPC) Value() int32 {
	if self.value == 0 {
		self.value = self.base.Value()*LevelNum + (int32)(OBJ_NPC)
	}
	return self.value
}

func (self *T_NPC) Player() *T_NPC_PLAYER {
	if self.player == nil {
		self.player = &T_NPC_PLAYER{base: self}
	}
	return self.player
}

func (self *T_NPC) TEST() *T_NPC_TEST {
	if self.test == nil {
		self.test = &T_NPC_TEST{base: self}
	}
	return self.test
}

type T_NPC_PLAYER struct {
	base  *T_NPC
	value int32
}

func (self *T_NPC_PLAYER) Value() int32 {
	if self.value == 0 {
		self.value = self.base.Value()*LevelNum + (int32)(npc.NPC_PLAYER)
	}
	return self.value
}

type T_NPC_TEST struct {
	base  *T_NPC
	value int32
}

func (self *T_NPC_TEST) Value() int32 {
	if self.value == 0 {
		self.value = self.base.Value()*LevelNum + (int32)(npc.NPC_TEST)
	}
	return self.value
}

type T_BUILDING struct {
	base  *T_Obj
	value int32
}

func (self *T_BUILDING) Value() int32 {
	if self.value == 0 {
		self.value = self.base.Value()*LevelNum + (int32)(OBJ_BUILDING)
	}
	return self.value
}
func init() {
	T_OBJ = &T_Obj{}
}
