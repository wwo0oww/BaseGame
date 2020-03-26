package client

import (
	"lib/cellnet"
	"lib/log"
)

type Client struct {
	session    cellnet.Session
	status     int32
	Heartindex int32
	PlayerPr   interface{}
}

func (self *Client) Session() cellnet.Session {
	return self.session
}
func (self *Client) Start(session cellnet.Session) {
	self.session = session
}
func (self *Client) Stop(err error) {
	self.session.Close()
	log.Debug(err.Error())
	self.session = nil
}
