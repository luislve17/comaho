package api

import (
	"os"
)

type DashboardData struct {
	MediaPath string
}

func getMediaPath() string {
	return os.Getenv("COMAHO_MEDIA_PATH")
}

func getDashboardData() DashboardData {
	return DashboardData{
		MediaPath: getMediaPath(),
	}
}
