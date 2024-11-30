package handler

import (
	"context"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"

	"github.com/songjiayang/cog-cluster/pkg/logger"
)

func PredictionProcess(ctx context.Context, t *asynq.Task) error {
	taskID := t.ResultWriter().TaskID()
	logger.Log().Info("resolve task", zap.String("task_id", taskID))
	// TODO:
	// 1. send request to cog-server
	// 2. cog-api callback
	return nil
}
