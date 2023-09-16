package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Welasco/AzureMonitorStarterPacksCollector/collectors"
	"github.com/Welasco/AzureMonitorStarterPacksCollector/common"
	logger "github.com/Welasco/AzureMonitorStarterPacksCollector/common/logger"
	"gopkg.in/gcfg.v1"
)

var cfg common.Config

// LoadConfig function
func LoadConfig() {
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

	LoadConfig()
	logger.Init(cfg.Main.LogPath, cfg.Main.LogLevel)
}

func main() {
	//logger.Info("Starting collectors...")
	logger.Info("Starting collectors...")

	// Create a slice of LogCollector interface
	var logcollector []collectors.LogCollector

	for siteName, website := range cfg.NginxCollectorWebsite {
		logcollector = append(logcollector, collectors.Newnginx_log(siteName, website, cfg.NginxCollector.LogPath))
	}

	// Sample code to add a collector
	//logcollector = append(logcollector, collectors.Newnginx_log(&cfg))

	// Graceful shutdown signals
	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1)

	// Call gracefullShutdown which gets executed when the signals SIGTERM OR SIGQUIT are received.
	// It happens at this line under the function gracefulShutdown <-ctx.Done()
	go gracefulShutdown(ctx, stop, &logcollector, &wg)

	// Start all collectors and add them to the WaitGroup
	for _, collector := range logcollector {
		wg.Add(1)
		go collector.Start(&wg)
	}

	// logger.Info("Sleeping for 11 seconds GetStatus")
	// time.Sleep(11 * time.Second)
	// for _, collector := range logcollector {
	// 	logger.Info("Current status:", collector.GetStatus())
	// }

	// logger.Info("Sleeping for 9 seconds before stopping the collector")
	// time.Sleep(9 * time.Second)

	// for _, collector := range logcollector {
	// 	collector.Stop()
	// }

	// for _, collector := range logcollector {
	// 	logger.Info("Current status:", collector.GetStatus())
	// }

	// logger.Info("Sleeping for 5 seconds after stop the collector")
	// time.Sleep(5 * time.Second)

	wg.Wait()
	logger.Info("Shutdown complete!")

	// Add config file for all collectors - Done
	// Config file must support the specifics of each collector, URL, File, etc - Done
	// Error handling everywhere - Almost Done still need to test at the go routine level to catch the error during the start of the collector *** Check with Alex where we should handle it in the main or in the collector
	// add comments - Done
	// Add a log module - Done
	// Use log file name from config file
	// graceful shutdown - Done
	// add multi site nginx support - Done
	// add log level - Done
	// Add a test module
	// Build the bin file
	// Create a deployment script/deb package

}

// Graceful shutdown function
// It gets executed when the signals SIGTERM OR SIGQUIT are received
// It stops all collectors and then stops the main loop
func gracefulShutdown(ctx context.Context, stop context.CancelFunc, logcollector *[]collectors.LogCollector, wg *sync.WaitGroup) {
	// <-ctx.Done() waits for the signals (SIGTERM OR SIGQUIT) to be received and then executes the code below
	<-ctx.Done()
	logger.Info("Received shutdown signal...")
	for _, collector := range *logcollector {
		collector.Stop()
	}
	logger.Info("Shut down!")
	stop()
	wg.Done()
}
