package handler

import (
	"context"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"

	"github.com/songjiayang/cog-cluster/pkg/cog"
	"github.com/songjiayang/cog-cluster/pkg/logger"
)

func PredictionProcess(ctx context.Context, t *asynq.Task) error {
	taskID := t.ResultWriter().TaskID()
	logger.Log().Info("resolve task", zap.String("task_id", taskID))

	if err := cog.GetClient().Predict(taskID, t.Payload()); err != nil {
		logger.Log().Error("cog process faield")
		return err
	}

	logger.Log().Info("processed")
	return nil
}
