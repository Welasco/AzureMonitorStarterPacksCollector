package main

import (
	"fmt"
	"time"

	"github.com/Welasco/AzureMonitorStarterPacksCollector/collectors"
	"github.com/Welasco/AzureMonitorStarterPacksCollector/common"
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
}

func main() {
	fmt.Println("Starting collectors...")
	var logcollector []collectors.LogCollector
	logcollector = append(logcollector, collectors.Newnginx_log(cfg))

	for _, collector := range logcollector {
		go collector.Start()
	}

	fmt.Println("Sleeping for 10 seconds GetStatus")
	time.Sleep(10 * time.Second)
	for _, collector := range logcollector {
		fmt.Println(collector.GetStatus())
	}

	fmt.Println("Sleeping for 50 seconds before stopping the collector")
	time.Sleep(50 * time.Second)

	for _, collector := range logcollector {
		collector.Stop()
	}

	for _, collector := range logcollector {
		fmt.Println(collector.GetStatus())
	}

	fmt.Println("Sleeping for 50 seconds after stop the collector")
	time.Sleep(50 * time.Second)

	// Add config file for all collectors - Done
	// Config file must support the specifics of each collector, URL, File, etc - Done
	// Error handling everywhere - Almost Done
	// Add a test module
	// Add a log module
	// add comments

}
