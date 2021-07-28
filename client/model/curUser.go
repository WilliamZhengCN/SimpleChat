package model

import (
	"FirstProject/SimpleChat/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
