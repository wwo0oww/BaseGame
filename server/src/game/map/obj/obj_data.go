package obj

// obj类型
type ObjType int32
const (
	OBJTYPE_PLAYER ObjType = iota
	OBJTYPE_NPC
)


// obj状态
type ObjStatus int32
const (
	STATUS_NONE ObjStatus = iota
	STATUS_MOVE
)
