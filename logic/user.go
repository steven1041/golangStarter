package logic

import (
	"golangStarter/dao/mysql"
	"golangStarter/dao/redis"
	"golangStarter/models"
	"golangStarter/pkg/jwt"
	"golangStarter/pkg/wechat"
)

type Code2Session struct {
	ErrCode    int32  `json:"err_code"`    // 错误码
	ErrMsg     string `json:"err_msg"`     // 错误信息
	Openid     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
}

func MiniProgrammerLogin(p *models.ParamMiniProgrammerLogin) (user *models.WxUser, err error) {
	user = &models.WxUser{
		NickName: p.NickName,
		Avatar:   p.Avatar,
		Gender:   p.Gender,
		Country:  p.Country,
		Province: p.Province,
		City:     p.City,
		Language: p.Language,
	}

	wechatProxy := &wechat.WeChat{}
	snsOauth2, err := wechatProxy.GetWxOpenIdFromOauth2(p.Code)
	if err != nil {
		return nil, err
	}
	if len(snsOauth2.Openid) == 0 {
		return nil, ErrorOpenidNil
	}
	user.OpenID = snsOauth2.Openid
	if err = mysql.MiniProgrammerLogin(user); err != nil {
		return nil, err
	}
	//生成JWT
	token, err := jwt.GenToken(user.OpenID, user.NickName)
	if err != nil {
		return
	}
	user.Token = token
	err = redis.SetToken(user.OpenID, token)
	if err != nil {
		return nil, err
	}
	return
}
