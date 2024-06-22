package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"golangStarter/dao/mysql"
	"golangStarter/logic"
	"golangStarter/models"
)

// MiniProgrammerLoginHandler 小程序登录
func MiniProgrammerLoginHandler(c *gin.Context) {
	//1.获取请求参数及参数校验
	p := new(models.ParamMiniProgrammerLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		//判断err是不是validator.ValidationErrors类型
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeToStruct(errs.Translate(trans)))
		return
	}

	//2.业务逻辑处理
	user, err := logic.MiniProgrammerLogin(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
		} else {
			ResponseError(c, CodeInvalidPassword)
		}
		return
	}
	//3.返回响应
	ResponseSuccess(c, gin.H{
		"openid":   fmt.Sprintf("%d", user.OpenID),
		"nickname": user.NickName,
		"token":    user.Token,
	})
}
