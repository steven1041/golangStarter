package redis

// redis的key使用命名空间的方式,方便查询和拆分
const (
	KeyPrefix              = "bluebell:"
	KeyPostTimeZSet        = "post:time"
	KeyPostScoreZSet       = "post:score"
	KeyPostVotedZSetPrefix = "post:voted:" //参数是post_id
	KeyCommunitySetPrefix  = "community:"  //set;保存每个分区下贴子的id

)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
