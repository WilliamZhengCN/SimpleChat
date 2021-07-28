package process

import (
	"SimpleChat/client/utils"
	"SimpleChat/common/message"
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func ShowMenu() {
	fmt.Println("********************")
	fmt.Println("1. Show Online User")
	fmt.Println("2. Send Message")
	fmt.Println("3. Message List")
	fmt.Println("4. Exist")
	fmt.Println("Please select 1--4")
	var key int
	fmt.Scanln(&key)
	smsProcess := &SmsProcess{}
	switch key {
	case 1:
		showOnlineUser()
	case 2:
		fmt.Println("Type the message you wanted to send: ")
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n')
		smsProcess.SendGroupMes(input)
	case 3:
		fmt.Println("message list")
	case 4:
		fmt.Println("exist")
		os.Exit(0)
	default:
		fmt.Printf("Error command")
	}
}

func serverProcessMes(conn net.Conn) {
	ts := utils.Transfer{
		Conn: conn,
	}
	fmt.Println("client is waitting for the message from server")
	for {
		mes, err := ts.ReadPkg()
		if err != nil {
			fmt.Println("Fail to read message from server. Error : ", err)
		}
		fmt.Println(mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			ShowGroupMes(&mes)
		default:
			fmt.Println("Unsupport message type")
		}
	}
}
