package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() error {
	r := gin.Default()
	r.StaticFS("/", http.Dir("assets"))
	//r.GET("/activity/:milliseconds?sex=male", activityHandler)
	return r.Run(":5555")
}

func main() {
	fmt.Println("Hello I am server")
	srv := NewServer()
	srv.Start()
}
