package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

const (
	UserOnline = iota
	UserOffLine
	UserBusy
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginResMes struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
	UserIds    []string
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	StatusCode int    `json:"statusCode"` //400 user exist  200
	Error      string `json:"error"`
}

type NotifyUserStatusMes struct {
	UserId string `json:"userId"`
	Status int    `json:"status"`
}

type SmsMes struct {
	Content string `json:"content"`
	User
}
