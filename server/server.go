package server

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/TulioGuaraldoB/springfield/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "8081",
		server: gin.Default(),
	}
}

func (s *Server) RunServer() {
	router := routes.ConfigRoutes()
	log.Fatal(router.Run(":" + s.port))
}
