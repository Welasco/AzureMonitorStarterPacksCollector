package collectors

type Nginx_log struct {
	LogPath           string
	Url               string
	ScrapeIntervalsec string
	LogFormat         string
}
