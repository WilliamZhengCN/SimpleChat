package process

import (
	"SimpleChat/client/model"
	"SimpleChat/common/message"
	"fmt"
)

var onlineUsers map[string]*message.User = make(map[string]*message.User, 100)
var curUser model.CurUser

func updateUserStatus(notifyMes *message.NotifyUserStatusMes) {
	notifyUserInfo, ok := onlineUsers[notifyMes.UserId]

	if !ok {
		notifyUserInfo = &message.User{
			Id: notifyMes.UserId,
		}
	}
	notifyUserInfo.UserStatus = notifyMes.Status
	onlineUsers[notifyMes.UserId] = notifyUserInfo

	showOnlineUser()
}

func showOnlineUser() {
	fmt.Println("Current online user:")
	for id := range onlineUsers {
		fmt.Println(id)
	}
}
