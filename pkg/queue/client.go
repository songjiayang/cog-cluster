package queue

import (
	"fmt"

	"github.com/hibiken/asynq"

	"github.com/songjiayang/cog-cluster/pkg/redis"
)

var client *asynq.Client

func InitClient() {
	client = asynq.NewClient(asynq.RedisClientOpt{Addr: redis.GetRedisAddr()})
}

func GetClient() *asynq.Client {
	return client
}

func GetPredictionTaskQueue(taskType string) string {
	return fmt.Sprintf("predictions:%s", taskType)
}

func Enqueue(taskType string, payload []byte) (string, error) {
	info, err := client.Enqueue(asynq.NewTask(taskType, payload))
	if err != nil {
		return "", err
	}
	return info.ID, nil
}
