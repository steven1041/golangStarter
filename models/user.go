package models

type User struct {
	UserID   int64  `db:"user_id"`
	UserName string `db:"username"`
	Password string `db:"password"`
	Token    string
}

type WxUser struct {
	OpenID   string `db:"open_id"`
	NickName string `db:"nickname"`
	Avatar   string `db:"avatar"`
	Gender   string `db:"gender"`
	City     string `db:"city"`
	Province string `db:"province"`
	Country  string `db:"country"`
	Language string `db:"language"`
	Token    string
}
