package redis

import (
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePreVote     = 432 // 每一票值多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeater   = errors.New("不允许重复相同值进行投票")
)

/*
	投票的几种情况：
	direction = 1时，有两种情况：
		1. 之前没有投过票，现在投赞成票    --> 更新分数和投票记录  差值的绝对值：1 +432
		2. 之前投反对票，现在改投赞成票    --> 更新分数和投票记录  差值的绝对值：2 +432*2
	direction = 0时，有两种情况：
		1. 之前投过反对票，现在取消投票    --> 更新分数和投票记录  差值的绝对值：1 +432
		2. 之前投过赞成票，现在取消投票    --> 更新分数和投票记录  差值的绝对值：1 -432
	direction = -1，有两种情况
		1. 之前没有投过票，现在投反对票    --> 更新分数和投票记录  差值的绝对值：1 -432
		2. 之前投过赞成票，现在改投反对票   --> 更新分数和投票记录 差值的绝对值：2 -432*2

投票的限制：
	每个帖子自发表之日起，一个星期内允许用户投票，超过一个星期就不允许再投了
	1. 到期后将redis中保存的赞成票以及反对票数存储到mysql中
	2. 到期之后删除那个（KeyPostVotedZsetPF）

*/

// CreatePost 添加到redis
func CreatePost(postID, communityID int64) error {
	// 同时成功，需要有一个事务操作
	pipeline := client.TxPipeline()

	// 帖子创建时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 帖子分数，基准值是时间，方便后面投票逻辑
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 更新：把帖子id加入社区set中
	cKey := getRedisKey(KeyCommunitySetPF) + strconv.Itoa(int(communityID))
	pipeline.SAdd(cKey, postID)

	// 事务执行
	_, err := pipeline.Exec()

	return err
}

func VoteForPost(userID, postID string, value float64) error {
	// 1. 判断投票的限制

	// 去redis取帖子发布时间
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds { //大于一周
		return ErrVoteTimeExpire
	}

	// 2. 更新帖子的分数
	// 先查当前用户给当前帖子的投票记录
	ov := client.ZScore(getRedisKey(KeyPostVotedZSetPrefix+postID), userID).Val()

	// 如果这一次投票的值和之前保存的值一致，就提示不允许重复投票
	if value == ov {
		return ErrVoteRepeater
	}

	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}

	diff := math.Abs(ov - value) // 计算两次投票的差值

	// 使用事务进行操作
	pipeline := client.TxPipeline()

	// 非常精简的操作
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePreVote, postID)

	// 3. 记录用户为该帖子投过票    key <-> redis.Z{value, score}
	if value == 0 {
		client.ZRem(getRedisKey(KeyPostVotedZSetPrefix+postID), userID)
	} else {
		client.ZAdd(getRedisKey(KeyPostVotedZSetPrefix+postID), redis.Z{
			Score:  value, // 当前用户投赞成还是反对
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
