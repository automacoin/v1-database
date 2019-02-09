package main

import (
	"flag"
	"log"
	"os"

	"github.com/automacoin/v1-database/api"
)

var (
	VERSION string
	BRANCH  string
	COMMIT  string
)

func main() {
	// flags
	// TODO
	// Process flags/config in their own file
	flag.Parse()

	// TODO
	// Have a unique logger across the whole service
	logger := log.New(os.Stdout, "[v1-database] ", log.LstdFlags)

	// TODO
	// Improve this message
	logger.Printf(
		"Starting AutomaCoin Database component | version: %s | branch: %s | commit: %s",
		VERSION, BRANCH, COMMIT[:6])

	// backend database
	// TODO
	// all of it!

	// set up API and go for it!
	router := api.Setup()
	api.Run(router)
}
