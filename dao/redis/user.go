package redis

import "golangStarter/pkg/jwt"

func SetToken(openid string, token string) (err error) {
	err = client.Set(KeyUserString+openid, token, jwt.TokenExpireDuration).Err()
	if err != nil {
		return
	}
	return nil
}

func CheckToken(openid string, token string) bool {
	tokenInRedis, err := client.Get(KeyUserString + openid).Result()
	if err != nil || token != tokenInRedis {
		return false
	} else {
		return true
	}
}
