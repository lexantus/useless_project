package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() error {
	r := gin.Default()
	r.StaticFS("/", http.Dir("client"))
	return r.Run(":5555")
}
