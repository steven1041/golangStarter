package redis

// redis的key使用命名空间的方式,方便查询和拆分
const (
	KeyPrefix     = "bluebell:"
	KeyUserString = "user:"
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
