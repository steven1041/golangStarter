package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"golangStarter/settings"
	"io"
	"net/http"
)

type WeChat struct {
	AppId       string `json:"app_id"`
	AppSecret   string `json:"app_secret"`
	AccessToken string `json:"access_token"`
}

type SnsOauth2 struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	Openid       string `json:"openid"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type AccessTokenErrorResponse struct {
	ErrMsg  string `json:"err_msg"`
	ErrCode string `json:"err_code"`
}

// 授权
func (weChat *WeChat) GetAuthUrl(redirectUrl string) string {

	oauth2Url := fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect",
		weChat.AppId, redirectUrl)
	return oauth2Url
}

// 通过code换取网页授权access_token
func (weChat *WeChat) GetWxOpenIdFromOauth2(code string) (*SnsOauth2, error) {

	requestLine := settings.Conf.WeChatConfig.Url + "appid=" + settings.Conf.WeChatConfig.AppId + "&secret=" + settings.Conf.WeChatConfig.Secret + "&js_code=" + code + settings.Conf.WeChatConfig.HttpTail

	resp, err := http.Get(requestLine)
	if err != nil || resp.StatusCode != http.StatusOK {
		zap.L().Error("发送get请求获取 openid 错误", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		zap.L().Error("发送get请求获取 openid 读取返回body错误", zap.Error(err))
		return nil, err
	}
	if bytes.Contains(body, []byte("errcode")) {
		ater := AccessTokenErrorResponse{}
		err = json.Unmarshal(body, &ater)
		if err != nil {
			zap.L().Error("发送get请求获取 openid 的错误信息", zap.Error(err))
			return nil, err
		}
		return nil, fmt.Errorf("%s", ater.ErrMsg)
	} else {
		atr := SnsOauth2{}
		err = json.Unmarshal(body, &atr)
		if err != nil {
			zap.L().Error("发送get请求获取 openid 返回数据json解析错误", zap.Error(err))
			return nil, err
		}
		return &atr, nil
	}
}
