package api

import (
	"fmt"
	"log"
	"os"

	"github.com/luislve17/comaho/utils"
)

type DashboardData struct {
	MediaPath   string
	MediaEnvVar string
	MediaErrMsg string
}

var ErrEnvVarNotSet = fmt.Errorf("environment variable COMAHO_MEDIA_PATH not set")

func getMediaPath(fsChecker func(string) error) (string, error) {
	envMediaPath := os.Getenv("COMAHO_MEDIA_PATH")
	if envMediaPath == "" {
		log.Println(ErrEnvVarNotSet.Error())
		return "null", ErrEnvVarNotSet
	}

	if err := fsChecker(envMediaPath); err != nil {
		log.Printf("Path is not accessible: %s\n", err.Error())
		return envMediaPath, err
	}

	return envMediaPath, nil
}

func getDashboardData() DashboardData {
	mediaPath, mediaAccessErr := getMediaPath(utils.CanReadDir)

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
