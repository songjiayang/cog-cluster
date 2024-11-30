package main

import (
	"github.com/songjiayang/cog-cluster/pkg/logger"
	"github.com/songjiayang/cog-cluster/pkg/server"
)

func main() {
	logger.Init()

	srv := server.New("cog-agent")
	srv.Run("0.0.0.0:8001")
}
