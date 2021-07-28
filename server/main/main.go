package main

import (
	"SimpleChat/server/model"
	"SimpleChat/server/utils"
	"fmt"
	"net"
	"time"
)

func main() {
	confhelper := utils.ConfigureHelper{}
	conf := confhelper.GetConf()
	//init pool for redis
	initPool(conf.RedisHost+":"+conf.RedisPort, 12, 0, 100*time.Second)
	initUserDao()

	listen, err := net.Listen(conf.ListenType, conf.ServerHost+":"+conf.ServerPort)
	defer listen.Close()
	if err != nil {
		fmt.Println("Fail to listen port 8765. Err:", err)
		return
	}
	for {
		fmt.Println("Server start to listen the port 8765")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Fail to accept. Error: ", err)
			return
		}
		go process1(conn)
	}
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func process1(conn net.Conn) {
	defer conn.Close()
	process := &Process{
		Conn: conn,
	}
	err := process.processInpoint()
	if err != nil {
		fmt.Print("Process error ", err)
		return
	}
}
