package queue

import (
	"github.com/hibiken/asynq"

	"github.com/songjiayang/cog-cluster/pkg/redis"
)

const (
	PredictionTask = "predictions:demo"
)

var client *asynq.Client

func InitClient() {
	client = asynq.NewClient(asynq.RedisClientOpt{Addr: redis.GetRedisAddr()})
}

func GetClient() *asynq.Client {
	return client
}

func Enqueue(taskType string, payload []byte) (string, error) {
	info, err := client.Enqueue(asynq.NewTask(PredictionTask, payload))
	if err != nil {
		return "", err
	}
	return info.ID, nil
}
