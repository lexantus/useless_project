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
	// Ensure the database is reachable
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
	// Check for errors in scanning
	if err != nil {
		// If the error is not nil, check if it's a "no rows found" error
		if errors.Is(err, sql.ErrNoRows) {
			// Return a 404 if no rows are found for the given milliseconds
			c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		} else {
			// Return a 500 for any other errors
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "details": err.Error()})
		}
		return
	}
	// c.JSON(http.StatusOK, gin.H{"message": "pong"})
	c.JSON(http.StatusOK, activity)
}

func (s *Server) Start() error {
	router := gin.Default()
	// router.Static("/", http.Dir("assets"))
	router.GET("/activity/:milliseconds", s.activityHandler)
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
