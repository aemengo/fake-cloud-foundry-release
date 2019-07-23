package main

import (
	"github.com/aemengo/fake-cloud-foundry/api"
	cfg "github.com/aemengo/fake-cloud-foundry/config"
	"github.com/aemengo/fake-cloud-foundry/db"
	"github.com/aemengo/fake-cloud-foundry/uaa"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("[USAGE] %s <config-path>", os.Args[0])
	}

	logger := log.New(os.Stdout, "[MAIN]", log.LstdFlags)

	config, err := cfg.New(os.Args[1])
	expectNoError(err)

	var (
		database  = db.New(config)
		apiServer = api.New(config, database)
		uaaServer = uaa.New(config)
		sigs      = make(chan os.Signal, 1)
	)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go launchServer("api", "8080", apiServer.Router(), logger)
	go launchServer("uaa", "8081", uaaServer.Router(), logger)

	<-sigs
}

func launchServer(name string, port string, handler http.Handler, logger *log.Logger) {
	logger.Printf("Launching %s server on :%s...\n", name, port)
	expectNoError(http.ListenAndServe(":"+port, handler))
}

func expectNoError(err error) {
	if err != nil {
		log.Fatalf("Failed to initialize: %s\n", err)
	}
}
