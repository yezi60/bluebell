package redis

import (
	"bluebell/models"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func getIDsFromKey(key string, page, size int64) ([]string, error) {
	// 2. 根据查询的索引点位置
	start := (page - 1) * size
	end := start + size - 1

	// 3. ZREVRANGE 按分数从大到小的顺序查询指定数量的元素
	return client.ZRevRange(key, start, end).Result()
}

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {

	// 从redis获取id
	// 1.根据用户请求中携带的order参数确定要查的redis key

	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}

	// 2. 确定查询的索引起始点
	return getIDsFromKey(key, p.Page, p.Size)
}

// 根据ids查询每篇帖子的赞成票的数据
func GetPostVoteData(ids []string) (data []int64, err error) {

	// 造成大量的RTT
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedZSetPrefix + id)
	//	// 查找key中分数是1的元素的数量 -> 统计每篇帖子的赞成票的数量
	//	v1 := client.ZCount(key, "1", "1").Val()
	//	data = append(data, v1)
	//}

	//使用pipeline优化

	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPrefix + id)
		pipeline.ZCount(key, "1", "1")
	}

	cmdrs, err := pipeline.Exec()
	if err != nil {
		return
	}

	data = make([]int64, 0, len(cmdrs))

	for _, cmdr := range cmdrs {
		v := cmdr.(*redis.IntCmd).Val()
		data = append(data, v)
	}

	return
}

// GetCommunityPostIDsInOrder 按社区根据ids
func GetCommunityPostIDsInOrder(p *models.ParamPostList) ([]string, error) {

	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}

	// 使用 zinterstore 把分区的帖子set与帖子分数的 zset 生成一个新的zset
	// 针对新的 zset 按之前的逻辑取数据

	// 社区的key bluebell:community:community_id
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityID)))

	// 社区的key bluebell:post:score/time:community_id
	// 利用缓存key减少zinterstore执行的次数
	key := orderKey + strconv.Itoa(int(p.CommunityID))

	// 不存在，需要计算
	if client.Exists(key).Val() < 1 {
		pipeline := client.Pipeline()

		// key作为后面的合并后的键名
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX", // 求两边的最大值
		}, cKey, orderKey) // zinterstore 计算

		pipeline.Expire(key, 60*time.Second) // 设置超时时间
		if _, err := pipeline.Exec(); err != nil {
			return nil, nil
		}
	}

	// 存在的话就根据key查询ids
	return getIDsFromKey(key, p.Page, p.Size)
}
