package node

type NODE_TYPE int

const Center NODE_TYPE = 0
const Game NODE_TYPE = 1
const Gateway NODE_TYPE = 2

type Node struct {
	name  string
	ntype NODE_TYPE
}

func NewNode(nodeName string,ntype NODE_TYPE) *Node{
	node := &Node{}

	node.SetName(nodeName)

	node.SetType(ntype)

	node.Start()

	return node
}

func (self *Node) GetName() string {
	return self.name
}

func (self *Node) SetName(name string) {
	self.name = name
}

func (self *Node) SetType(ntype NODE_TYPE) {
	self.ntype = ntype
}

func (self *Node) GetType() NODE_TYPE {
	return self.ntype
}


func (self *Node) Start(){
	//todo 连接其他节点 并保持心跳
}
