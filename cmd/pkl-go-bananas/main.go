package main

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/flyinprogrammer/pkl-go-bananas/gen/appconfig"
	"github.com/flyinprogrammer/pkl-go-bananas/internal"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	configFilePath := setupLocalConfigFile()
	cfg, err := appconfig.LoadFromPath(context.TODO(), configFilePath)
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}
	printNetworkData(cfg)
	sleepForContainerTesting()
}

func setupLocalConfigFile() (configFilePath string) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Error getting user config directory: %v", err)
	}

	cliCfgDirPath := filepath.Join(cfgDir, "pkl-go-bananas") + string(filepath.Separator)
	if _, err := os.Stat(cliCfgDirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(cliCfgDirPath, 0700); err != nil {
			log.Fatalf("Error creating cli config dir: %v", err)
		}
	}

	configFilePath = filepath.Join(cliCfgDirPath, "config.pkl")
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		data := internal.DefaultTemplate
		err = os.WriteFile(configFilePath, []byte(data), 0600)
		if err != nil {
			log.Fatalf("Error creating default config file: %v", err)
		}
	}
	return
}

func printNetworkData(cfg *appconfig.AppConfig) {
	fmt.Println("Let's print out your network data:")
	for _, computer := range cfg.Computers {
		fmt.Printf("Computer: %s\n", computer.Name)
		for _, ip := range computer.IpAddress {
			fmt.Printf("\tIP: %s\n", ip)
		}
	}
}

func sleepForContainerTesting() {
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
