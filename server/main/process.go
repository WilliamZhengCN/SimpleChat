package main

import (
	message "SimpleChat/common/message"
	"SimpleChat/server/process"
	"SimpleChat/server/utils"
	"fmt"
	"io"
	"net"
)

type Process struct {
	Conn net.Conn
}

func (this *Process) serverProcessMes(mes *message.Message) (err error) {

	fmt.Print("mes = ", mes)
	switch mes.Type {
	case message.LoginMesType:

		userProcess := &process.UserProcess{
			Conn: this.Conn,
		}
		err = userProcess.ServerProcessLogin(mes)
	case message.RegisterMesType:
		userProcess := &process.UserProcess{
			Conn: this.Conn,
		}
		err = userProcess.ServerProcessRegister(mes)
	case message.SmsMesType:
		smsprocess := process.SmsProcess{}
		smsprocess.SendGroupMes(mes)
	default:
		fmt.Println("unknow message type from client")
	}
	return err
}

func (this *Process) processInpoint() (err error) {
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err == io.EOF {
			fmt.Println("custom connection has been closed")
			return err
		}
		if err != nil {
			fmt.Println("fail to read message from server. Error:", err)
			return err
		}
		fmt.Println("received message is : ", mes)

		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
