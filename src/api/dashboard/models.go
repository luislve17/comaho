package dashboard

import (
	"fmt"
)

type DashboardData struct {
	EnvData EnvData
	LibData LibData
}

type EnvData struct {
	MediaPath              string
	MediaPathEnvVar        string
	DockerVolumePathEnvVar string
	ErrMsg                 string
}

type LibData struct {
	ComicAndMangaData []CoMaData
	ErrMsg            string
}

type CoMaData struct {
	Id           string
	Name         string
	LastModified string
}

var ErrEnvVarNotSet = fmt.Errorf("environment variable COMAHO_MEDIA_PATH not set")
