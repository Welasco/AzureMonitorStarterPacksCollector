package main

import (
	"fmt"
	"time"

	"github.com/Welasco/AzureMonitorStarterPacksCollector/collectors"
	"github.com/Welasco/AzureMonitorStarterPacksCollector/common"
	logger "github.com/Welasco/AzureMonitorStarterPacksCollector/common/logger"
	"gopkg.in/gcfg.v1"
)

var cfg common.Config

// LoadConfig function
func LoadConfig() {
	fmt.Println("Loading config file...")

	var err error
	if err = gcfg.ReadFileInto(&cfg, "config_collector.ini"); err != nil {
		fmt.Println("Unable to open or unrecognized entries at config_collector.ini")
		//fmt.Println(err)
		panic(err)
	}
}

// Go init function
// Loading config file
func init() {
	fmt.Println("Initializing...")

	LoadConfig()
	logger.Init(cfg.Main.LogPath)
}

func main() {
	//logger.Info("Starting collectors...")
	logger.Info("Starting collectors...")

	var logcollector []collectors.LogCollector
	logcollector = append(logcollector, collectors.Newnginx_log(cfg))

	for _, collector := range logcollector {
		go collector.Start()
	}

	logger.Info("Sleeping for 10 seconds GetStatus")
	time.Sleep(10 * time.Second)
	for _, collector := range logcollector {
		logger.Info(collector.GetStatus())
	}

	logger.Info("Sleeping for 50 seconds before stopping the collector")
	time.Sleep(50 * time.Second)

	for _, collector := range logcollector {
		collector.Stop()
	}

	for _, collector := range logcollector {
		logger.Info(collector.GetStatus())
	}

	logger.Info("Sleeping for 50 seconds after stop the collector")
	time.Sleep(50 * time.Second)

	// Add config file for all collectors - Done
	// Config file must support the specifics of each collector, URL, File, etc - Done
	// Error handling everywhere - Almost Done still need to test at the go routine level to catch the error during the start of the collector
	// add comments - Done
	// Add a test module
	// Add a log module - Done
	// graceful shutdown

}
