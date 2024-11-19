package api

import (
	"fmt"
	"log"
	"os"

	"github.com/luislve17/comaho/utils"
)

type DashboardData struct {
	MediaPath              string
	MediaPathEnvVar        string
	DockerVolumePathEnvVar string
	ErrMsg                 string
}

var ErrEnvVarNotSet = fmt.Errorf("environment variable COMAHO_MEDIA_PATH not set")

func getMediaPath(fsChecker func(string) error) (string, error) {
	envMediaPath := os.Getenv("COMAHO_MEDIA_PATH")
	if envMediaPath != "" {
		if err := fsChecker(envMediaPath); err != nil {
			log.Printf("Local path is not accessible: %s\n", err.Error())
			return envMediaPath, err
		}
		return envMediaPath, nil
	}

	dockerVolumePath := os.Getenv("COMAHO_DOCKER_VOLUME_PATH")
	if dockerVolumePath != "" {
		if err := fsChecker("/app/media"); err != nil {
			log.Printf("Docker path is not accessible: %s\n", err.Error())
			return dockerVolumePath, err
		}
		return "/app/media", nil
	}

	log.Println(ErrEnvVarNotSet.Error())
	return "null", ErrEnvVarNotSet
}

func getDashboardData() DashboardData {
	mediaPath, mediaAccessErr := getMediaPath(utils.CanReadDir)

	var mediaErrMsg string
	if mediaAccessErr != nil {
		mediaErrMsg = mediaAccessErr.Error()
	} else {
		mediaErrMsg = "" // No error, so set to an empty string
	}
	envMediaPath := os.Getenv("COMAHO_MEDIA_PATH")
	dockerVolumePath := os.Getenv("COMAHO_DOCKER_VOLUME_PATH")

	return DashboardData{
		MediaPath:              mediaPath,
		MediaPathEnvVar:        envMediaPath,
		DockerVolumePathEnvVar: dockerVolumePath,
		ErrMsg:                 mediaErrMsg,
	}
}
