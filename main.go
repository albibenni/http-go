package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("message.txt")
	if err != nil {
		log.Fatal("error", err)

	}
}
