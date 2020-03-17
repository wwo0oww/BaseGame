package config

import "lib/node"

const HeartBeat int32 = 10

func GetGameCode() string {
	return "arrival"
}
func GetAgentCode() string {
	return "debug"
}
func GetServerID() int16 {
	return 1
}

func GetNodeName() string {
	return "game1"
}

func GetNodeType() node.NODE_TYPE {
	return node.Game
}

func GetMapNodeName() string {
	return GetNodeName()
}

func GetMapWH() (int32, int32) {
	return 100, 100
}

func MaxMapWorldX() int32 {
	W, _ := GetMapWH()
	return W * GetAreaSize()
}

func MaxMapWorldY() int32 {
	_, H := GetMapWH()
	return H * GetAreaSize()
}

func GetAreaSize() int32 {
	return 100
}

// 必须是 GetMapWH()的W*H的整数倍
func GetSingleThreadAreaNum() int {
	return 20
}
