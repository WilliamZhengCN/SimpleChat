package process

import (
	"SimpleChat/client/utils"
	message "SimpleChat/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {
}

func (this *UserProcess) LoginIn(user message.User) (err error) {
	confhelper := utils.ConfigureHelper{}
	conf := confhelper.GetConf()
	conn, err := net.Dial(conf.ConnectType, conf.ServerHost+":"+conf.ServerPort)
	fmt.Println(conf)
	if err != nil {
		fmt.Println("Fail to connect to server. Err: ", err, conn)
		return
	}
	defer conn.Close()

	var mess message.Message
	mess.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.Id = user.Id
	loginMes.Password = user.Password

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json format error. error = ", err)
		return
	}
	mess.Data = string(data)

	realData, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("json format error. error = ", err)
		return
	}

	transfer := utils.Transfer{
		Conn: conn,
	}
	err = transfer.WritePkg(realData)
	if err != nil {
		fmt.Println("fail to register. Error: ", err)
	}

	mes, err := transfer.ReadPkg()
	if err != nil {
		return err
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.StatusCode == 200 {
		curUser.Conn = conn
		curUser.Id = user.Id
		curUser.UserStatus = message.UserOnline
		fmt.Println("login success.Current online user.")
		for _, val := range loginResMes.UserIds {
			fmt.Println("User Id :\t", val)
			user := &message.User{
				Id:         val,
				UserStatus: message.UserOnline,
			}
			onlineUsers[val] = user
		}

		go serverProcessMes(conn)

		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}

func (this *UserProcess) Register(id string, password string, userName string) (err error) {
	confhelper := utils.ConfigureHelper{}
	conf := confhelper.GetConf()
	conn, err := net.Dial(conf.ConnectType, conf.ServerHost+":"+conf.ServerPort)

	if err != nil {
		fmt.Println("Fail to connect to server. Err: ", err, conn)
		return
	}
	defer conn.Close()

	var mess message.Message
	mess.Type = message.RegisterMesType
	var rgMes message.RegisterMes
	rgMes.User.Id = id
	rgMes.User.UserName = userName
	rgMes.User.Password = password

	data, err := json.Marshal(rgMes)
	if err != nil {
		fmt.Println("json format error. error = ", err)
		return
	}
	mess.Data = string(data)

	realData, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("json format error. error = ", err)
		return
	}

	transfer := utils.Transfer{
		Conn: conn,
	}
	err = transfer.WritePkg(realData)
	if err != nil {
		fmt.Println("fail to register. Error: ", err)
	}

	registerMes, err := transfer.ReadPkg()
	if err != nil {
		fmt.Println("fail to register. Error: ", err)
	}

	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(registerMes.Data), &registerResMes)
	if registerResMes.StatusCode == 200 {
		fmt.Println("register success. Plese login in")
		os.Exit(0)
	} else {
		fmt.Print(registerResMes.Error)
		os.Exit(0)
	}
	return
}
