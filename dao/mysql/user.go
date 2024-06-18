package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"golangStarter/models"
)

// 把每一步数据库操作封装成函数,被Logic层根据业务需求来调用
const secret = "stone.com"

// CheckUserExit 检查用户是滞存在
func CheckUserExit(username string) (err error) {
	sqlStr := `select count(user_id) from user where username =?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 向数据库插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	//对密码进行加密
	password := encryptPassword(user.Password)
	//执行SQL语句入库
	sqlStr := `insert into user(user_id,username,password) values (?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.UserName, password)
	return

}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id,username,password from user where username=?`
	err = db.Get(user, sqlStr, user.UserName)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		//查询数据库失败
		return err
	}
	//判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}

func GetUserById(userId int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := "select user_id,username from user where user_id=?"
	err = db.Get(user, sqlStr, userId)
	return
}
