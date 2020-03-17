package protoplus

import (
	"lib/cellnet"
	"lib/cellnet/codec"
	"lib/protoplus/proto"
)

type protoplus struct {
}

func (self *protoplus) Name() string {
	return "protoplus"
}

func (self *protoplus) MimeType() string {
	return "application/binary"
}

func (self *protoplus) Encode(msgObj interface{}, ctx cellnet.ContextSet) (data interface{}, err error) {

	return proto.Marshal(msgObj)

}

func (self *protoplus) Decode(data interface{}, msgObj interface{}) error {

	return proto.Unmarshal(data.([]byte), msgObj)
}

func init() {

	codec.RegisterCodec(new(protoplus))
}
