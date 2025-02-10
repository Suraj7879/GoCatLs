package cmd

import (
	"log"
	"os"
)

func Execute() {
	cla := os.Args
	fileName := cla[1]
	currDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	b, err := os.ReadFile(currDir + "/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(b)
}
