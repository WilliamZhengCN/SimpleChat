package main

import (
	"FirstProject/SimpleChat/client/model"
	"FirstProject/SimpleChat/client/process"
	"fmt"
)

func main() {
	ShowMainMenu()
}

func ShowMainMenu() {
	var cmd int
	fmt.Println("*********************************")
	fmt.Println("Welcome to Simple Chat:")
	fmt.Println("\t1 Login In")
	fmt.Println("\t2 Register")
	fmt.Println("\t3 Exist")
	fmt.Println("Please type 1--3 for operation:")
	fmt.Println("*********************************")

	for {
		fmt.Scanf("%d\n", &cmd)
		switch cmd {
		case 1:
			HandleLoginIn()
		case 2:
			var id string
			var password string
			var userName string
			fmt.Println("input User ID:")
			fmt.Scanln(&id)
			fmt.Println("input Password:")
			fmt.Scanln(&password)
			fmt.Println("input User Name:")
			fmt.Scanln(&userName)
			userProcess := &process.UserProcess{}
			userProcess.Register(id, password, userName)
		case 3:
		default:
			fmt.Println("Input Error.")
		}
	}
}

func HandleLoginIn() {
	var id string
	var password string
	fmt.Println("User ID:")
	fmt.Scanln(&id)
	fmt.Println("Password:")
	fmt.Scanln(&password)
	userInfo := model.ChatUser{
		Id:       id,
		Password: password,
	}
	userProcess := &process.UserProcess{}
	err := userProcess.LoginIn(userInfo)
	if err != nil {
		fmt.Println("Fail to login int. Error: ", err)
	}
}
