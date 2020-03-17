package tests

import (
	"fmt"
	"lib/cellnet"
	_ "lib/cellnet/codec/httpform"
	_ "lib/cellnet/codec/httpjson"
	"lib/cellnet/peer"
	httppeer "lib/cellnet/peer/http"
	"lib/cellnet/proc"
	_ "lib/cellnet/proc/http"
	"reflect"
	"testing"
)

const httpTestAddr = "127.0.0.1:8081"

func TestHttp(t *testing.T) {

	p := peer.NewGenericPeer("http.Acceptor", "httpserver", httpTestAddr, nil)

	proc.BindProcessorHandler(p, "http", func(raw cellnet.Event) {

		if matcher, ok := raw.Session().(httppeer.RequestMatcher); ok {
			switch {
			case matcher.Match("GET", "/hello_get"):

				// 默认返回json
				raw.Session().Send(&httppeer.MessageRespond{
					Msg: &HttpEchoACK{
						Token: "get",
					},
				})
			case matcher.Match("POST", "/hello_post"):

				// 默认返回json
				raw.Session().Send(&httppeer.MessageRespond{
					Msg: &HttpEchoACK{
						Token: "post",
					},
				})

			}
		}

	})

	p.Start()

	requestThenValid(t, "GET", "/hello_get", &HttpEchoREQ{
		UserName: "kitty_get",
	}, &HttpEchoACK{
		Token: "get",
	})

	requestThenValid(t, "POST", "/hello_post", &HttpEchoREQ{
		UserName: "kitty_post",
	}, &HttpEchoACK{
		Token: "post",
	})

	p.Stop()
}

func requestThenValid(t *testing.T, method, path string, req, expectACK interface{}) {

	p := peer.NewGenericPeer("http.Connector", "httpclient", httpTestAddr, nil).(cellnet.HTTPConnector)

	ackMsg := reflect.New(reflect.TypeOf(expectACK).Elem()).Interface()

	err := p.Request(method, path, &cellnet.HTTPRequest{
		REQMsg: req,
		ACKMsg: ackMsg,
	})

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if !reflect.DeepEqual(ackMsg, expectACK) {
		t.Log("unexpect token result", err)
		t.FailNow()
	}

}

type HttpEchoREQ struct {
	UserName string
}

type HttpEchoACK struct {
	Token  string
	Status int32
}

func (self *HttpEchoREQ) String() string { return fmt.Sprintf("%+v", *self) }
func (self *HttpEchoACK) String() string { return fmt.Sprintf("%+v", *self) }
