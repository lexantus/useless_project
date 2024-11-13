package server

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func Start() {
  r := gin.Default()
  r.Get("/activity", func(c *gin.Context) {
    c.JSON(http.StatusOk, gin.H{
	message: "pong"
    })
  })
  r.Run()
}
