package redis

import "github.com/songjiayang/cog-cluster/pkg/util"

var redisAddr = "127.0.0.1:6379"

func GetRedisAddr() string {
	redisAddr = util.GetEnvOr("REDIS_ADDR", redisAddr)
	return redisAddr
}
