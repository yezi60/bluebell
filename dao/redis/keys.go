package redis

// redis key

// redis key注意使用命名空间的方法，区分不同的key，方便业务查询和拆分

const (
	Prefix                 = "bluebell:"
	KeyPostTimeZSet        = "post:time"   //zset： 帖子及发帖时间
	KeyPostScoreZSet       = "post:score"  // zset：帖子及投票的分数
	KeyPostVotedZSetPrefix = "post:voted:" // zset：记录用户及投票类型;参数是post id

	KeyCommunitySetPF = "community:" // set：保存每个分区下帖子的id
)

// 给redis key加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
