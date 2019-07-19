package main

import (
	cfg "github.com/aemengo/fake-cloud-foundry/config"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("[USAGE] %s <config-path>", os.Args[0])
	}

	_, err := cfg.New(os.Args[1])
	expectNoError(err)
}

func expectNoError(err error) {
	if err != nil {
		log.Fatalf("Failed to initialize: %s\n", err)
	}
}