package logic

import (
	"golangStarter/dao/mysql"
	"golangStarter/models"
	"golangStarter/pkg/jwt"
	"golangStarter/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	//1.判断用户是否存在
	err = mysql.CheckUserExit(p.Username)
	if err != nil {
		return err
	}
	//2.生成UID
	userID := snowflake.GenID()
	user := &models.User{
		UserID:   userID,
		UserName: p.Username,
		Password: p.Password,
	}
	//3.保存进数据库
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		UserName: p.Username,
		Password: p.Password,
	}
	//传递的是指针，就能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	//生成JWT
	token, err := jwt.GenToken(user.UserID, user.UserName)
	if err != nil {
		return
	}
	user.Token = token
	return
}