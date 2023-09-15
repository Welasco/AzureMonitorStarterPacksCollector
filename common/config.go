package common

type Config struct {
	// Main section
	Main struct {
		LogPath string
	}
	// NginxCollector section
	NginxCollector NginxCollector
	// sub-section NginxCollectorWebsite
	NginxCollectorWebsite NginxCollectorWebsite
}

type NginxCollector struct {
	LogPath string
	// sub-section NginxCollectorWebsite
	//NginxCollectorWebsite NginxCollectorWebsite
}

type NginxCollectorWebsite map[string]*WebSite

type WebSite struct {
	Url               string
	ScrapeIntervalsec int
}

// config file struct definition config_collector.ini
// type Config struct {
// 	// Main section
// 	Main struct {
// 		LogPath string
// 	}
// 	// NginxCollector section
// 	NginxCollector struct {
// 		LogPath string
// 		// sub-section NginxCollectorWebsite
// 		NginxCollectorWebsite map[string]*struct {
// 			SiteName          string
// 			Url               string
// 			ScrapeIntervalsec int
// 		}
// 	}
// }

// type NginxCollectorWebsite map[string]*struct {
// 	SiteName          string
// 	Url               string
// 	ScrapeIntervalsec int
// }

// type NginxCollector struct {
// 	LogPath string
// 	// sub-section NginxCollectorWebsite
// 	NginxCollectorWebsite NginxCollectorWebsite
// }
