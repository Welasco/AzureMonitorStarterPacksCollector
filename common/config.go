package common

// config file struct definition config_collector.ini
type Config struct {
	NginxCollector struct {
		LogPath           string
		Url               string
		ScrapeIntervalsec int
	}
}
