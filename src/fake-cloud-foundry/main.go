package main

import (
	"fmt"
	"github.com/aemengo/fake-cloud-foundry/api"
	cfg "github.com/aemengo/fake-cloud-foundry/config"
	"github.com/aemengo/fake-cloud-foundry/db"
	"github.com/aemengo/fake-cloud-foundry/router"
	"github.com/aemengo/fake-cloud-foundry/uaa"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("[USAGE] %s <config-path>", os.Args[0])
	}

	rand.Seed(time.Now().UTC().UnixNano())

	logger := log.New(os.Stdout, "[MAIN]", log.LstdFlags)

	config, err := cfg.New(os.Args[1])
	expectNoError(err)

	var (
		routerSwitch = router.New()
		database     = db.New(config)
		apiServer    = api.New(config, database)
		uaaServer    = uaa.New(config, database)
	)

	routerSwitch.Add(fmt.Sprintf("api.%s", config.Domain()), apiServer.Router())
	routerSwitch.Add(fmt.Sprintf("uaa.%s", config.Domain()), uaaServer.Router())

	logger.Printf("Launching server on :%s...\n", config.Port)
	expectNoError(http.ListenAndServe(":"+config.Port, routerSwitch))
}

func expectNoError(err error) {
	if err != nil {
		log.Fatalf("Failed to initialize: %s\n", err)
	}
}
