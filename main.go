package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("message.txt")
	if err != nil {
		log.Fatal("error", err)
	}
	str := ""
	for {
		data := make([]byte, 8)
		n, err := file.Read(data)
		if err != nil {
			break
		}
		data = data[:n]
		if i := bytes.IndexByte(data, '\n'); i != -1 {
			str += string(data[:i])
			data = data[i+1:]

			fmt.Printf("read: %s\n", str)
			str = ""

		}
		str += string(data)

	}
	if len(str) != 0 {
		fmt.Printf("read: %s\n", str)
	}
}
