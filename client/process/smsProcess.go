package process

import (
	"SimpleChat/client/utils"
	"SimpleChat/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.Id = curUser.Id
	smsMes.UserStatus = curUser.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("fail Marshal")
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("fail Marshal")
	}

	tf := &utils.Transfer{
		Conn: curUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("fail WritePkg. Error:", err)
	}
	return
}
