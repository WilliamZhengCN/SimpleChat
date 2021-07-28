package process

import "fmt"

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[string]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[string]*UserProcess, 1024),
	}
}

func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

func (this *UserMgr) DeleteOnlineUser(userId string) {
	delete(this.onlineUsers, userId)
}

func (this *UserMgr) GetAllOnlineUsers() map[string]*UserProcess {
	return this.onlineUsers
}

func (this *UserMgr) GetOnlineUserById(userId string) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("user %v is not online ", userId)
		return
	}
	return
}
