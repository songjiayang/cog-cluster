package redis

import (
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/songjiayang/cog-cluster/pkg/util"
)

var (
	redisAddr = "127.0.0.1:6379"
	rdb       *redis.Client
)

func GetRedisAddr() string {
	redisAddr = util.GetEnvOr("REDIS_ADDR", redisAddr)
	return redisAddr
}

func GetDB() *redis.Client {
	sync.OnceFunc(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr: GetRedisAddr(),
		})
	})()

	return rdb
}
