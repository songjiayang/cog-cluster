package handler

import (
	"io"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
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
	// TODO:
	// 1. resolve callback data
	// 2. save task result to redis and s3
	// 3. update task status
}
