package cmd

import (
	"fmt"
	"os"

	"log"
)

func Execute() {
	dir := "."
	if len(os.Args) >= 2 {
		dir = os.Args[1]
	}

	fileInfo, err := os.Stat(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}

	if !fileInfo.IsDir() {
		fmt.Println(fileInfo.Name())
		return
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(os.Stderr, "Error reading directory: %v\n", err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
