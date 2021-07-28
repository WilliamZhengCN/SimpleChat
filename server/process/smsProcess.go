package process

import (
	"SimpleChat/common/message"
	"SimpleChat/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("Fail to Unmarshal. Err: ", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("Fail to Unmarshal. Err: ", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		if id == smsMes.Id {
			continue
		}

		this.SendMesToUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("Fail to WritePkg. Error:", err)
	}
}
