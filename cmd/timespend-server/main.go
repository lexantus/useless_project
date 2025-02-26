package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lexantus/useless_project/configs"
	"github.com/lexantus/useless_project/internal/handlers"
	"github.com/lexantus/useless_project/internal/repos"
	"log"
	"net/http"
)

func main() {
	activityRepo, err := repos.NewActivityRepo(*configs.GetConfig())
	if err != nil {
		log.Fatalf("Failed to create activityRepo %v", err)
	}

	activityHandler := handlers.NewActivityHandler(activityRepo)

	router := gin.Default()
	router.GET("/activity/:milliseconds", activityHandler.Handler)
	router.StaticFS("/assets", http.Dir("assets"))
	router.StaticFile("/index.html", "./assets/index.html")
	err = router.Run(":5555")

	if err != nil {
		log.Fatalf("Can not run server %v", err)
	}
}
