package cmd

import (
	"fmt"
	"log"
	"os"
)

func Execute() {
	currDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.ReadDir(currDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range f {
		fmt.Println(file.Name())
	}
}
