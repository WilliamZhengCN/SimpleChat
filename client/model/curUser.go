package model

import (
	"SimpleChat/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
