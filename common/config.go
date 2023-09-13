package common

type Config struct {
	NginxCollector struct {
		LogPath           string
		Url               string
		ScrapeIntervalsec int
	}
}
