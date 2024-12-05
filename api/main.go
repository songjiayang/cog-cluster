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

	srv.GET("/v1/predictions/:prediction_id", handler.PredictionGet)
	srv.POST("/v1/predictions", handler.PredictionCreate)
	srv.POST("/v1/models/:namespace/:model_name/predictions", handler.PredictionCreate)

	// inner api
	srv.POST("/inner/predictions/:prediction_id/callback", handler.PredictionCallback)

	srv.Run("0.0.0.0:8000")
}
