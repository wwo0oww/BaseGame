package area

import (
	"game/map/obj"
)

func (self *Area) JoinObj(newObj *obj.Obj, nCount int32) error {
	return self.AddObj(newObj, nCount)
}
