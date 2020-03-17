package peer

import (
	"fmt"
	"lib/cellnet"
	"sort"
)

type PeerCreateFunc func() cellnet.Peer

var peerByName = map[string]PeerCreateFunc{}

// 注册Peer创建器
func RegisterPeerCreator(f PeerCreateFunc) {

	// 临时实例化一个，获取类型
	dummyPeer := f()

	if _, ok := peerByName[dummyPeer.TypeName()]; ok {
		panic("duplicate peer type: " + dummyPeer.TypeName())
	}

	peerByName[dummyPeer.TypeName()] = f
}

// Peer创建器列表
func PeerCreatorList() (ret []string) {

	for name := range peerByName {
		ret = append(ret, name)
	}

	sort.Strings(ret)
	return
}

// cellnet自带的peer对应包
func getPackageByPeerName(name string) string {
	switch name {
	case "tcp.Connector", "tcp.Acceptor", "tcp.SyncConnector":
		return "lib/cellnet/peer/tcp"
	case "udp.Connector", "udp.Acceptor":
		return "lib/cellnet/peer/udp"
	case "gorillaws.Acceptor", "gorillaws.Connector", "gorillaws.SyncConnector":
		return "lib/cellnet/peer/gorillaws"
	case "http.Connector", "http.Acceptor":
		return "lib/cellnet/peer/http"
	case "redix.Connector":
		return "lib/cellnet/peer/redix"
	case "mysql.Connector":
		return "lib/cellnet/peer/mysql"
	default:
		return "package/to/your/peer"
	}
}

// 创建一个Peer
func NewPeer(peerType string) cellnet.Peer {
	peerCreator := peerByName[peerType]
	if peerCreator == nil {
		panic(fmt.Sprintf("peer type not found '%s'\ntry to add code below:\nimport (\n  _ \"%s\"\n)\n\n",
			peerType,
			getPackageByPeerName(peerType)))
	}

	return peerCreator()
}

// 创建Peer后，设置基本属性


func NewGenericPeer(peerType, name, addr string, q cellnet.EventQueue) cellnet.GenericPeer {

	p := NewPeer(peerType)
	gp := p.(cellnet.GenericPeer)
	gp.SetName(name)
	gp.SetAddress(addr)
	gp.SetQueue(q)
	return gp
}
