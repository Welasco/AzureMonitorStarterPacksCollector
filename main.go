package main

import (
	"fmt"
	"time"

	"github.com/Welasco/AzureMonitorStarterPacksCollector/collectors"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(time.Now().Format("2006-01-02T15:04:05"))

	var logcollector []collectors.LogCollector

	logcollector = append(logcollector, collectors.Newnginx_log())

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

	// Add config file for all collectors
	// Config file must support the specifics of each collector, URL, File, etc
	// Error handling everywhere
	// Add a test module
	// Add a log module
	// check if we must need a type.go file for each collector
	// review the pointers

}
