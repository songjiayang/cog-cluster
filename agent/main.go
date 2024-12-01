package main

import (
	"log"

	"github.com/hibiken/asynq"

	"github.com/songjiayang/cog-cluster/agent/handler"
	"github.com/songjiayang/cog-cluster/pkg/logger"
	"github.com/songjiayang/cog-cluster/pkg/queue"
)

func main() {
	logger.Init()

	srv := queue.NewServer()
	mux := asynq.NewServeMux()

	mux.HandleFunc(queue.PredictionTask, handler.PredictionProcess)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
