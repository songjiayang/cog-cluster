package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/songjiayang/cog-cluster/pkg/cog"
	"github.com/songjiayang/cog-cluster/pkg/logger"
	"github.com/songjiayang/cog-cluster/pkg/queue"
)

func PredictionGet(ctx *gin.Context) {
	taskID := ctx.Param("prediction_id")
	logger.Log().Info("resolve task id", zap.String("task_id", taskID))
	// TODO:
	// 1. get result from redis
	// 2. response to user
}

func PredictionCreate(ctx *gin.Context) {
	payload, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		logger.Log().Error("load request body with error", zap.Error(err))
		ctx.Status(http.StatusBadRequest)
		return
	}

	taskID, err := queue.Enqueue(queue.PredictionTask, payload)
	if err != nil {
		logger.Log().Error("add task with failed", zap.Error(err))
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"task_id": taskID,
	})
}

func PredictionCallback(ctx *gin.Context) {
	taskID := ctx.Param("prediction_id")
	logger.Log().Info("resolve task id", zap.String("task_id", taskID))
	payload, _ := io.ReadAll(ctx.Request.Body)
	logger.Log().Info("cog-server webhook", zap.String("body", string(payload)))

	var output cog.Output
	if err := json.Unmarshal(payload, &output); err != nil {
		logger.Log().Error("resolve cog-server webhook data failed", zap.Error(err))
		return
	}

	// TODO:
	// send data to redis
	if output.IsSuccess() {
		logger.Log().Info("task predict success", zap.Any("task", output))
	}
}
