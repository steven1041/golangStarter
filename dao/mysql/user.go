package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"golangStarter/models"
)

// 把每一步数据库操作封装成函数,被Logic层根据业务需求来调用
const secret = "stone.com"

// CheckWxUserExit  检查微信用户是否存在
func CheckWxUserExit(openid string) (any interface{}) {
	sqlStr := `select count(openid) from user where openid =?`
	var count int
	if err := db.Get(&count, sqlStr, openid); err != nil {
		return err
	}
	if count > 0 {
		return true
	} else {
		return false
	}

}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func MiniProgrammerLogin(user *models.WxUser) (err error) {
	// 检查openID是否存在
	result := CheckWxUserExit(user.OpenID)
	err, ok := result.(error)
	if ok {
		return err
	}
	//如查存在,直接更新数据库
	isExit, _ := result.(bool)
	if isExit {
		sqlStr := `UPDATE user SET nickname=?,SET avatar=?,SET gender=?,SET country=?,SET province=?,SET city=?,SET language=? WHERE openid=?`
		_, err = db.Exec(sqlStr, user.NickName, user.Avatar, user.Gender, user.Country, user.Province, user.City, user.Language, user.OpenID)
	} else { //如果不存在，插入数据库
		sqlStr := `insert into user (open_id,nickname,avatar,gender,country,province,city,language) values (?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlStr, user.OpenID, user.NickName, user.Avatar, user.Gender, user.Country, user.Province, user.City, user.Language)
		return
	}
	return
}
