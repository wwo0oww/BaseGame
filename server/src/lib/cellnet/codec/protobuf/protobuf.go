package protobuf

import (
	"github.com/gogo/protobuf/proto"
	"lib/cellnet"
	"lib/cellnet/codec"
)

type protobuf struct {
}

func (self *protobuf) Name() string {
	return "protobuf"
}

func (self *protobuf) MimeType() string {
	return "application/binary"
}

func (self *protobuf) Encode(msgObj interface{}, ctx cellnet.ContextSet) (data interface{}, err error) {

	return proto.Marshal(msgObj.(proto.Message))

}

func (self *protobuf) Decode(data interface{}, msgObj interface{}) error {

	return proto.Unmarshal(data.([]byte), msgObj.(proto.Message))
}

func init() {

	codec.RegisterCodec(new(protobuf))
}
