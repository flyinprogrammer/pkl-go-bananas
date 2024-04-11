package main

import (
	"context"
	"fmt"
	"github.com/flyinprogrammer/pkl-go-bananas/gen/appconfig"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Error getting user config directory: %v", err)
	}
	configFile := filepath.Join(cfgDir, "pkl-go-bananas/config.pkl")
	log.Printf("Looking for Config file in: %s", configFile)

	cfg, err := appconfig.LoadFromPath(context.TODO(), "pkl/dev/config.pkl")
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	fmt.Println("Let's print out your network data:")
	for _, computer := range cfg.Computers {
		fmt.Printf("Computer: %s\n", computer.Name)
		for _, ip := range computer.IpAddress {
			fmt.Printf("\tIP: %s\n", ip)
		}
	}
	value, exists := os.LookupEnv("PKL_GO_BANANAS_SLEEP")
	if exists {
		sleepTime, err := time.ParseDuration(value)
		if err != nil {
			log.Fatalf("Error parsing sleep time (https://pkg.go.dev/maze.io/x/duration#ParseDuration): %v", err)
		}
		log.Printf("Sleeping for %s.", sleepTime)
		time.Sleep(sleepTime)
	}
}
