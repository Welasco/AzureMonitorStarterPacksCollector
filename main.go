package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
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

	// Graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	var wg sync.WaitGroup
	wg.Add(1)
	go gracefulShutdown(ctx, stop, &logcollector, &wg)

	for _, collector := range logcollector {
		wg.Add(1)
		go collector.Start(&wg)
	}

	logger.Info("Sleeping for 11 seconds GetStatus")
	time.Sleep(11 * time.Second)
	for _, collector := range logcollector {
		logger.Info("Current status:", collector.GetStatus())
	}

	// logger.Info("Sleeping for 9 seconds before stopping the collector")
	// time.Sleep(9 * time.Second)

	// for _, collector := range logcollector {
	// 	collector.Stop()
	// }

	for _, collector := range logcollector {
		logger.Info("Current status:", collector.GetStatus())
	}

	logger.Info("Sleeping for 5 seconds after stop the collector")
	time.Sleep(5 * time.Second)

	logger.Info("Reach Wait() Exiting...")
	wg.Wait()
	fmt.Println("Main done")
	// Add config file for all collectors - Done
	// Config file must support the specifics of each collector, URL, File, etc - Done
	// Error handling everywhere - Almost Done still need to test at the go routine level to catch the error during the start of the collector
	// add comments - Done
	// Add a test module
	// Add a log module - Done
	// graceful shutdown
	// add multi site nginx support

}

func gracefulShutdown(ctx context.Context, stop context.CancelFunc, logcollector *[]collectors.LogCollector, wg *sync.WaitGroup) {
	<-ctx.Done()

	for _, collector := range *logcollector {
		collector.Stop()
	}
	logger.Info("Shutting down...")
	stop()
	wg.Done()
}
