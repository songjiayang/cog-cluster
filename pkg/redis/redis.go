package redis

import "os"

var redisAddr = "127.0.0.1:6379"

func GetRedisAddr() string {
	addr := os.Getenv("REDIS_ADDR")
	if addr != "" {
		redisAddr = addr
	}

	return redisAddr
}
