package area

import (
	"game/config"
	"game/core"
	"game/map/obj"
)

func CheckObjCollide(obj1, obj2 *obj.Obj) bool {
	if obj1 == obj2{
		return false
	}
	pos1 := obj1.GetPos()
	pos2 := obj2.GetPos()
	cover1 := obj1.GetSize()
	cover2 := obj2.GetSize()
	return CheckPosCollide(pos1, pos2, cover1, cover2)
}
func CheckPosCollide(pos1, pos2, cover1, cover2 core.Position) bool {
	return core.AbsInt32(pos1.X, pos2.X) <= (cover1.X+cover2.X) ||
		core.AbsInt32(pos1.Y, pos2.Y) <= (cover1.Y+cover2.Y)
}

func GetAreaPos(pos core.Position) core.Position {
	size := config.GetAreaSize()
	return core.Position{X: pos.X / size, Y: pos.X / size}
}
