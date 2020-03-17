package tests

import (
	"fmt"
	"lib/cellnet"
	"lib/cellnet/codec"
	_ "lib/cellnet/codec/binary"
	"lib/cellnet/util"
	"reflect"
)

//type TestData struct{
//	x int32
//	y int32
//}

type TestEchoACK1 struct{

	Msg   string
	Value []byte
}

type TestEchoACK struct {
	Msg   string
	Value int32
}

func (self *TestEchoACK) String() string { return fmt.Sprintf("%+v", *self) }

func init() {
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*TestEchoACK)(nil)).Elem(),
		ID:    int(util.StringHash("tests.TestEchoACK")),
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*TestEchoACK1)(nil)).Elem(),
		ID:    int(util.StringHash("tests.TestEchoACK1")),
	})

}
