package main

import (
	"github.com/songjiayang/cog-cluster/api/handler"
	"github.com/songjiayang/cog-cluster/pkg/logger"
	"github.com/songjiayang/cog-cluster/pkg/queue"
	"github.com/songjiayang/cog-cluster/pkg/server"
)

func main() {
	logger.Init()
	queue.InitClient()

	srv := server.New("cog-api")

	srv.GET("/predictions/:prediction_id", handler.PredictionGet)
	srv.POST("/predictions", handler.PredictionCreate)

	srv.Run("0.0.0.0:8000")
}
