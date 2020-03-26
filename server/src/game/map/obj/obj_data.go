package obj

// obj类型
type ObjType int32

const (
	OBJ_NPC ObjType = 1
	OBJ_BUILDING ObjType = 2
)

// obj状态
type ObjStatus int32

const (
	STATUS_NONE   ObjStatus = 1
	STATUS_MOVE   ObjStatus = 2
	STATUS_ATTACK ObjStatus = 4
)
