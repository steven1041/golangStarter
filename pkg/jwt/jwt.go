package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenExpireDuration = time.Hour * 24

var mySecret = []byte("夏天夏天悄悄的过去")

// MyClaims jwt包的jwt.StandardClaims只包含了官方字段
// 如果想要保存更多的信息，都可以添加到这个结构体中
type MyClaims struct {
	Openid   string `json:"openid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成token
func GenToken(openId string, username string) (string, error) {
	c := MyClaims{
		openId,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "bluebell",
		},
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	//使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(mySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenStr string) (*MyClaims, error) {
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenStr, mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}

	return nil, errors.New("invalid token")
}
