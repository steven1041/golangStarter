package middlewares

import (
	"github.com/gin-gonic/gin"
	"golangStarter/controller"
	"golangStarter/pkg/jwt"
	"strings"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//客户端携带Token有三种方式1.放在请求头 2.放在请求体 3.放在URI
		//这里假设 Token放在Header的Authorization中,并使用Bearer开头
		//Authorization:Bearer xxx.xxx.xxx
		//这里的具体实现方式要依据实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedAuth)
			c.Abort()
			return
		}
		//按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		//parts[1]是获取到的token,我们使用之前定义好的解析JWT函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
		}
		//将当前请求的user_id信息保存到请求的上下文c上
		c.Set(controller.CtxUserIDKey, mc.UserID)
		c.Next()
	}
}
