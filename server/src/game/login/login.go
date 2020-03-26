package login

import (
	"game/client"
	"game/core"
	"game/db"
	mMsg "game/map/msg"
	"game/map/obj"
	"game/player"
	"proto"
	"proto/net"
)

const OP_LOGIN int32 = 1
const OP_REG int32 = 2

func Handle(tos net.MLoginTos, client *client.Client) {

	if "" == tos.Name || "" == tos.Pwd {
		client.Session().Send(&net.MLoginToc{
			Op:      tos.Op,
			Errcode: proto.LOGIN_ERROR_1,
		})
	} else {
		switch tos.Op {
		case OP_LOGIN:
			do_login(tos, client)
		default:
			do_reg(tos, client)
		}
	}
}
func do_login(tos net.MLoginTos, client *client.Client) {
	account, UserID := db.GetAccountInfo(tos.Name, tos.Pwd)
	if account == "" {
		client.Session().Send(&net.MLoginToc{
			Op:      tos.Op,
			Errcode: proto.LOGIN_ERROR_1,
		})
	} else {
		client.Session().Send(&net.MLoginToc{
			Op:      tos.Op,
			Errcode: proto.NORMORL,
		})
		player := player.NewPlayer(UserID)
		player.SetClient(client)
		obj := &obj.Obj{
			ObjID:      player.Player_id,
			Pos:        player.Pos,
			Size:       core.Position{X: 5, Y: 5},
			PlayerFlag: 1,
			Name:       account,
			Type:       obj.T_OBJ.NPC().Player().Value()}
		obj.SetSpeed(1)
		player.SendMsgToMap(&mMsg.JoinObj{Obj: obj,
			Times: 10})
	}

}
func do_reg(tos net.MLoginTos, client *client.Client) {
	account := db.GetAccountInfoByName(tos.Name)
	if account != "" {
		client.Session().Send(&net.MLoginToc{
			Op:      tos.Op,
			Errcode: proto.LOGIN_ERROR_2,
		})
		return
	}
	err := db.InsertPlayer(tos.Name, tos.Pwd, player.NewPlayerID())
	if err != nil {
		client.Session().Send(&net.MLoginToc{
			Op:      tos.Op,
			Errcode: proto.LOGIN_ERROR_0,
		})
		return
	}

	client.Session().Send(&net.MLoginToc{
		Op:      tos.Op,
		Errcode: proto.NORMORL,
	})

}
