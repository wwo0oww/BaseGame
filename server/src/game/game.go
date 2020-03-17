package game

import (
	"game/config"
	"game/db"
	"game/sql"
	"lib/node"
)

var MaxPlayerNum int32 = 0

var self_node *node.Node
func StartGame(){
	sql.WaitForRunning()
	MaxPlayerNum = db.GetMaxPlayerNum()

	self_node = node.NewNode(config.GetNodeName(),config.GetNodeType())
}

func Node() *node.Node{
	return self_node
}
