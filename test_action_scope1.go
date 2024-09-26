package main

import (
	"log"
	"os"
)

var cwd string
func init() {
	var err error
	cwd, err = os.Getwd() // NOTE: wrong!
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory = %s", cwd)
}

func main() {
	log.Printf("Current directory = %s", cwd)
}