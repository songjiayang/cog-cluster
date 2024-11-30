package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/songjiayang/cog-cluster/pkg/logger"
	"go.uber.org/zap"
)

type Server struct {
	name string
	*gin.Engine
}

func New(name string) *Server {
	return &Server{
		name: name,

		Engine: gin.Default(),
	}
}

func (s *Server) Run(addr ...string) {
	s.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	logger.Log().Info("server start ...", zap.String("name", s.name))
	s.Engine.Run(addr...)
}
