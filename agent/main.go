package main

import (
	"log"
	"os"

	"github.com/hibiken/asynq"

	"github.com/songjiayang/cog-cluster/agent/handler"
	"github.com/songjiayang/cog-cluster/pkg/logger"
	"github.com/songjiayang/cog-cluster/pkg/queue"
)

func main() {
	logger.Init()

	srv := queue.NewServer()
	mux := asynq.NewServeMux()

	mux.HandleFunc(
		queue.GetPredictionTaskQueue(os.Getenv("COG_SERVER_TYPE")),
		handler.PredictionProcess,
	)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
