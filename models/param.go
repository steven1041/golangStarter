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

type ParamMiniProgrammerLogin struct {
	Code     string `json:"code"`
	NickName string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Gender   string `json:"gender"`
	City     string `json:"city"`
	Province string `json:"province"`
	Country  string `json:"country"`
	Language string `json:"language"`
}
