package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/songjiayang/cog-cluster/pkg/cog"
	"github.com/songjiayang/cog-cluster/pkg/logger"
	"github.com/songjiayang/cog-cluster/pkg/queue"
	"github.com/songjiayang/cog-cluster/pkg/redis"
)

func PredictionGet(ctx *gin.Context) {
	taskID := ctx.Param("prediction_id")
	logger.Log().Info("resolve task id", zap.String("task_id", taskID))
	output, err := redis.GetDB().Get(ctx, redis.TaskOutputKey(taskID)).Bytes()
	if err != nil {
		logger.Log().Info("resolve task output with error", zap.Error(err))
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.Data(200, "application/json; charset=utf-8", output)
}

func PredictionCreate(ctx *gin.Context) {
	var input PredictionInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		logger.Log().Error("load request body with error", zap.Error(err))
		ctx.Status(http.StatusBadRequest)
		return
	}

	// TODO: version validate
	if input.Version == "" {
		logger.Log().Error("empty version input")
		ctx.Status(http.StatusBadRequest)
		return
	}

	taskInput := input.Marshal()
	taskQueue := queue.GetPredictionTaskQueue(input.Version)
	taskID, err := queue.Enqueue(taskQueue, taskInput)
	if err != nil {
		logger.Log().Error("add task with failed", zap.Error(err))
		ctx.Status(http.StatusInternalServerError)
		return
	}
	logger.Log().Info("add a new task", zap.String("queue", taskQueue), zap.String("task", string(taskInput)))

	// for sync request
	if ctx.GetHeader("Prefer") == "wait" {
		ticker := time.NewTicker(250 * time.Millisecond)
		timeout, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
	LOOP:
		for {
			select {
			case <-ticker.C:
				output, err := redis.GetDB().Get(ctx, redis.TaskOutputKey(taskID)).Bytes()
				if err == nil {
					ctx.Data(200, "application/json; charset=utf-8", output)
					return
				}
			case <-timeout.Done():
				logger.Log().Warn("predict timeout", zap.String("task_id", taskID))
				break LOOP
			}
		}
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

	if output.IsSuccess() {
		logger.Log().Info("task predict success", zap.Any("task", output))
		data, _ := json.Marshal(output)
		// set output value to redis
		redis.GetDB().Set(ctx, redis.TaskOutputKey(taskID), data, 24*time.Hour)
	}
}
