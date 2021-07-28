package model

import (
	"SimpleChat/common/message"
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return userDao
}

func (this *UserDao) GetUserById(conn redis.Conn, id string) (user *User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			fmt.Println("User not exist by id ", id)
			err = ERROR_USER_NOTEXIST
		}
		return
	}
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json format error. Error:", err)
	}
	return
}

func (this *UserDao) Login(id string, password string) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.GetUserById(conn, id)
	if err != nil {
		return
	}
	if user.Password == password {
		fmt.Println("success to login")
	} else {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDao) Register(userInfo *message.User) (err error) {
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.GetUserById(conn, userInfo.Id)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	data, err := json.Marshal(userInfo)
	if err != nil {
		return
	}

	_, err = conn.Do("HSet", "users", userInfo.Id, string(data))
	if err != nil {
		fmt.Println("Fail to register user. Error:", err)
		return
	}

	return
}
