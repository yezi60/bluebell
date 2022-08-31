package redis

import (
	"bluebell/settings"
	"fmt"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

func Init(conf *settings.RedisConfig) (err error) {

	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			conf.Host,
			conf.Port),
		Password:     conf.Password,
		DB:           conf.DB,
		PoolSize:     conf.PoolSize,
		MinIdleConns: conf.MinIdleConns,
	})

	_, err = client.Ping().Result()
	if err != nil {
		zap.L().Error("redis connect fail:", zap.Error(err))
		return
	}
	return
}

func Close() {
	_ = client.Close()
}
