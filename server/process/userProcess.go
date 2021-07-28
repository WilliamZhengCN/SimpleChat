package process

import (
	message "FirstProject/SimpleChat/common/message"
	"FirstProject/SimpleChat/server/model"
	"FirstProject/SimpleChat/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId string
}

func (this *UserProcess) NotifyOnlineUsers(userId string) {
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		up.NotifyOnlineUser(userId)
	}
}

func (this *UserProcess) NotifyOnlineUser(userId string) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyMes message.NotifyUserStatusMes
	notifyMes.UserId = userId
	notifyMes.Status = message.UserOnline

	data, err := json.Marshal(notifyMes)
	if err != nil {
		fmt.Println("Fail Marshal. Error:", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("Fail Marshal. Error:", err)
		return
	}

	ts := utils.Transfer{
		Conn: this.Conn,
	}
	err = ts.WritePkg(data)
	if err != nil {
		fmt.Println("Fail to send message. Error: ", err)
		return
	}
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("fail to Unmarshal. Error: ", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.StatusCode = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.StatusCode = 506
			registerResMes.Error = "Inner error. Error: " + err.Error()
		}
	} else {
		registerResMes.StatusCode = 200
		fmt.Println("Register sccuess")
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		return err
	}
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		return err
	}

	transfer := &utils.Transfer{
		Conn: this.Conn,
	}
	err = transfer.WritePkg(data)

	return err
}

func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("fail to Unmarshal. Error: ", err)
		return
	}
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes

	user, err := model.MyUserDao.Login(loginMes.Id, loginMes.Password)
	if err != nil {
		if err == model.ERROR_USER_NOTEXIST {
			loginResMes.StatusCode = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.StatusCode = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.StatusCode = 505
			loginResMes.Error = "server inner error"
		}
	} else {
		loginResMes.StatusCode = 200
		this.UserId = loginMes.Id
		fmt.Println("Current user ", user)
		userMgr.AddOnlineUser(this)
		this.NotifyOnlineUsers(loginMes.Id)
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		return err
	}
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		return err
	}

	transfer := &utils.Transfer{
		Conn: this.Conn,
	}
	err = transfer.WritePkg(data)

	return err
}
