package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Server struct {
	db *sql.DB
}

type Activity struct {
	ID         int    `json:"id"`
	TimespanMs int    `json:"timespan_ms"`
	Desc       string `json:"description"`
}

func NewServer() (*Server, error) {
	connStr := "postgres://postgres:password@db:5432/timespan?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return &Server{db: db}, nil
}

func (s *Server) activityHandler(c *gin.Context) {
	milliseconds := c.Param("milliseconds")
	r := s.db.QueryRow("SELECT id, timespan_ms, description from activity WHERE timespan_ms = $1", milliseconds)
	var activity Activity
	err := r.Scan(&activity.ID, &activity.TimespanMs, &activity.Desc)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "details": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, activity)
}

func (s *Server) Start() error {
	router := gin.Default()
	router.GET("/activity/:milliseconds", s.activityHandler)
	router.StaticFS("/assets", http.Dir("assets"))
	router.StaticFile("/index.html", "./assets/index.html")
	return router.Run(":5555")
}

func main() {
	fmt.Println("Hello I am server.go")
	srv, err := NewServer()
	if err != nil {
		log.Fatalf("Failed to initialize server %v", err)
	}
	if err := srv.Start(); err != nil {
		log.Fatalf("Server failed to start %v", err)
	}
}
