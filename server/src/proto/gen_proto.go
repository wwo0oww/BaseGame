package proto

import (
	"bytes"
	"encoding/binary"
	"game"
	"lib/cellnet"
	"lib/cellnet/codec"
	_ "lib/cellnet/codec/binary"
	_ "lib/cellnet/codec/protobuf"
	"proto/net"
	"proto/rpc"
	"reflect"
	"sync/atomic"
)

var key2id map[interface{}]int32
var id2key map[int32]interface{}
var start_index int32

func init() {
	start_index = 1000
	key2id = make(map[interface{}]int32)
	id2key = make(map[int32]interface{})
	Add((*net.TestEchoACK)(nil))
	Add((*net.MLoginTos)(nil))
	Add((*net.MLoginToc)(nil))
	Add((*net.MHeartbeatTos)(nil))
	Add((*net.MHeartbeatToc)(nil))
	Add((*rpc.PMapMsg)(nil))
}

func AddMap(key interface{}) {
	var NewID = atomic.AddInt32(&start_index, 1)
	id2key[NewID] = key
	key2id[key] = NewID
}

func Add(key interface{}) {
	var NewID = atomic.AddInt32(&start_index, 1)

	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("protobuf"),
		Type:  reflect.TypeOf(key).Elem(),
		ID:    (int)(NewID),
	})
}

func EncodeMapMsg(rawPlayer interface{}, raw interface{}) interface{} {
	switch raw.(type) {
	case rpc.PMapMsg: // rpc 消息
		return raw
	default: // 本地裸包处理
		MsgID := key2id[raw]
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.LittleEndian, raw)

		return &rpc.PMapMsg{Player: game.Player2PPlayer(rawPlayer), MsgId: MsgID, Data: buf.Bytes()}
	}
}

func DecodeMapMsg(raw interface{}) (Msg interface{}) {
	switch raw.(type) {
	case rpc.PMapMsg: // rpc 消息
		data := raw.(rpc.PMapMsg)
		buf := new(bytes.Buffer)
		buf.Write(data.Data)
		Msg = reflect.New(reflect.TypeOf(id2key[data.MsgId]).Elem())
		binary.Read(buf, binary.LittleEndian, &Msg)
		return
	default: // 本地裸包处理
		Msg = raw
	}
	return
}
