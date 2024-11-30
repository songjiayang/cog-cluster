package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/songjiayang/cog-cluster/pkg/logger"
	"go.uber.org/zap"
)

type Server struct {
	name string
	app  *gin.Engine
}

type Router struct {
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
}

func New(name string) *Server {
	return &Server{
		name: name,
		app:  gin.Default(),
	}
}

func (s *Server) Resources(routers []Router) {
	for _, router := range routers {
		s.app.Handle(router.Method, router.Path, router.HandlerFunc)
	}

	s.app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func (s *Server) Run(addr ...string) {
	logger.Log().Info("server start ...", zap.String("name", s.name))
	s.app.Run(addr...)
}
