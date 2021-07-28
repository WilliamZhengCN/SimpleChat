package process

import (
	"FirstProject/SimpleChat/common/message"
	"encoding/json"
	"fmt"
)

func ShowGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	 err := json.Unmarshal([]byte(mes.Data), & smsMes)
	if err != nil{
		fmt.Println("Fail to Unmarshal. Error: ", err)	
	}
	
	fmt.Printf("User %s said to all : %s \n", smsMes.Id, smsMes.Content)
}
