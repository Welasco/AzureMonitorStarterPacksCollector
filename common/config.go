package common

// config file struct definition config_collector.ini
type Config struct {
	// Main section
	Main struct {
		LogPath string
	}
	// NginxCollector section
	NginxCollector struct {
		LogPath           string
		Url               string
		ScrapeIntervalsec int
	}
}
