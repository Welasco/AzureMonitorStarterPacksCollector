package collectors

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/Welasco/AzureMonitorStarterPacksCollector/common"
)

type nginx_properties struct {
	active_connections string
	server_accepts     string
	server_handled     string
	server_requests    string
	reading            string
	writing            string
	waiting            string
}

func (nlog *Nginx_log) CollectLog(bodystr string) error {

	// Regular expression patterns
	activeConnPattern := regexp.MustCompile(`Active connections: (\d+)`)
	serviceConnPattern := regexp.MustCompile(` (\d+) (\d+) (\d+)`)
	rwwPattern := regexp.MustCompile(`Reading: (\d+) Writing: (\d+) Waiting: (\d+)`)

	// Find matches in the text
	activeConnMatches := activeConnPattern.FindStringSubmatch(bodystr)
	serviceConnMatches := serviceConnPattern.FindStringSubmatch(bodystr)
	rwwMatches := rwwPattern.FindStringSubmatch(bodystr)

	// popupate the struct Nginx_log
	// nginx_log := &Nginx_log{}
	nginx_properties := nginx_properties{
		active_connections: activeConnMatches[1],
		server_accepts:     serviceConnMatches[1],
		server_handled:     serviceConnMatches[2],
		server_requests:    serviceConnMatches[3],
		reading:            rwwMatches[1],
		writing:            rwwMatches[2],
		waiting:            rwwMatches[3],
	}

	// write to csv file
	nlog.write_to_csv(&nginx_properties)

	return nil
}

func (nlog *Nginx_log) write_to_csv(nginx_log *nginx_properties) {
	file, err := os.OpenFile(nlog.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		//log.Fatal(err)
		fmt.Println("Unable to open file", nlog.LogPath)
		fmt.Println(err)
		fmt.Println("Stopping NGINX collector for now...")
		nlog.Stop()
	}

	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	//writer.Write([]string{"Active_connections", "Server_accepts", "Server_handled", "Server_requests", "Reading", "Writing", "Waiting"})
	writer.Write([]string{time.Now().Format("2006-01-02T15:04:05"), nginx_log.active_connections, nginx_log.server_accepts, nginx_log.server_handled, nginx_log.server_requests, nginx_log.reading, nginx_log.writing, nginx_log.waiting})
}

// Variable to control the status of the collector
// Running or Stopped
var shouldStop bool

func (nlog *Nginx_log) Start() error {
	shouldStop = false
	for !shouldStop {
		respbody, resp, err := common.Http_client(nlog.Url)
		if err != nil {
			fmt.Println("Failed to connect to NGINX", nlog.Url)
			fmt.Println(err)
			fmt.Println("Stopping NGINX collector for now...")
			nlog.Stop()
			return err
		}
		if resp.StatusCode != 200 {
			fmt.Println("Failed to connect to NGINX", nlog.Url)
			fmt.Println("Status code:", resp.StatusCode)
			fmt.Println("Stopping NGINX collector for now...")
			nlog.Stop()
			return err
		}
		nlog.CollectLog(respbody)
		time.Sleep(time.Duration(nlog.ScrapeIntervalsec) * time.Second)
	}

	return nil
}

func (nlog *Nginx_log) Stop() error {
	shouldStop = true
	return nil
}

func (nlog *Nginx_log) GetStatus() string {
	if !shouldStop {
		return "Running"
	} else if shouldStop {
		return "Stopped"
	}
	return "Unknown"
}

func Newnginx_log(cfg common.Config) *Nginx_log {
	var nlog Nginx_log = Nginx_log{
		LogPath:           cfg.NginxCollector.LogPath,
		Url:               cfg.NginxCollector.Url,
		ScrapeIntervalsec: cfg.NginxCollector.ScrapeIntervalsec,
	}
	return &nlog
}
