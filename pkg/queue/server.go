package queue

import (
	"github.com/hibiken/asynq"
	"github.com/songjiayang/cog-cluster/pkg/redis"
)

var (
	srv *asynq.Server
)

func NewServer() *asynq.Server {
	srv = asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: redis.GetRedisAddr(),
		},
		asynq.Config{
			Concurrency: 1,
		},
	)

	return srv
}
