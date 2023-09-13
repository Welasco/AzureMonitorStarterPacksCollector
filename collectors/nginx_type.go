package collectors

// Nginx_log struct definition for the NGINX collector
type Nginx_log struct {
	LogPath           string
	Url               string
	ScrapeIntervalsec int
}
