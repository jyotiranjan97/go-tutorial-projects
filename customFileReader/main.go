package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := os.Args[1]

	file, error := os.Open(filePath)

	if error != nil {
		fmt.Println("Error", error)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
