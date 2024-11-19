package api

import (
	"log"
	"os"

	"github.com/luislve17/comaho/utils"
)

type DashboardData struct {
	MediaPath   string
	MediaEnvVar string
	MediaErrMsg string
}

func getMediaPath() (string, error) {
	MEDIA_PATH := "/app/media"
	err := utils.CanReadDir(MEDIA_PATH)
	if err != nil {
		log.Printf("Couldn't access path. %s...\n", err.Error())
		return "nil", err
	}
	return MEDIA_PATH, nil
}

func getDashboardData() DashboardData {
	mediaPath, mediaAccessErr := getMediaPath()

	var mediaErrMsg string
	if mediaAccessErr != nil {
		mediaErrMsg = mediaAccessErr.Error()
	} else {
		mediaErrMsg = "" // No error, so set to an empty string
	}

	return DashboardData{
		MediaPath:   mediaPath,
		MediaEnvVar: os.Getenv("COMAHO_MEDIA_PATH"),
		MediaErrMsg: mediaErrMsg,
	}
}
