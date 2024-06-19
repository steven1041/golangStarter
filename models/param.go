package models

//定义请求参数的结构体

const (
	OrderType  = "time"
	OrderScore = "score"
)

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamMiniProgrammerRegister struct {
	Code      string `json:"code"`
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
	/**
	put custom code here
	e.g.
		Email			string		`json:"email" binding:"required"`
		Password 		string		`json:"password" binding:"required"`
	*/
}
