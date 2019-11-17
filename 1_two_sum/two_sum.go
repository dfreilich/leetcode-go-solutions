package main

import (
	"fmt"
	"log"
)

func main() {
	if err := Run(); err != nil {
		log.Fatalf("Failed to run: %s", err)
	}
	fmt.Println("Successfully ran")
}

func Run() error {
	return nil
}
