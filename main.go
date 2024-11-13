package main

import (
	"fmt"
	"github.com/lexantus/useless_project/server"
	"log"
)

func main() {
	fmt.Println("Running server...")
	s := server.NewServer()
	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
