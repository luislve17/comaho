package dashboard

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/luislve17/comaho/utils"
)

func ServeDashboard(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		dashboardData := getDashboardData()
		err := utils.RenderTemplate(w, tmpl, dashboardData)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	}
}

func getDashboardData() DashboardData {
	mediaPath, mediaAccessErr := getMediaPath(utils.CanReadDir)
	mediaErrMsg := parseErrMsg(mediaAccessErr)
	envMediaPath := os.Getenv("COMAHO_MEDIA_PATH")
	dockerVolumePath := os.Getenv("COMAHO_DOCKER_VOLUME_PATH")

	comaData, contentErr := getLibraryData(mediaPath, dockerVolumePath)
	contentErrMsg := parseErrMsg(contentErr)

	result := DashboardData{
		EnvData{
			MediaPath:              mediaPath,
			MediaPathEnvVar:        envMediaPath,
			DockerVolumePathEnvVar: dockerVolumePath,
			ErrMsg:                 mediaErrMsg,
		},
		LibData{
			ComicAndMangaData: comaData,
			ErrMsg:            contentErrMsg,
		},
	}
	return result
}

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

func parseErrMsg(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

func getLibraryData(localMediaPath string, dockerMediaPath string) ([]CoMaData, error) {
	var mediaPath string
	if localMediaPath != "" {
		mediaPath = localMediaPath
	} else if dockerMediaPath != "" {
		mediaPath = dockerMediaPath
	}

	var ComicAndMangaData []CoMaData
	entries, err := os.ReadDir(mediaPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			info, _ := os.Stat(filepath.Join(mediaPath, entry.Name()))
			id, name := getIdAndNameFromDir(entry.Name())
			ComicAndMangaData = append(ComicAndMangaData, CoMaData{
				Id:           id,
				Name:         name,
				LastModified: info.ModTime().Local().Format("2006-01-02 03:04:05"),
			})
		}
	}

	return ComicAndMangaData, nil
}

func getIdAndNameFromDir(input string) (string, string) {
	re := regexp.MustCompile(`^\(([^)]+)\)\s*(.+)$`)
	matches := re.FindStringSubmatch(strings.TrimSpace(input))

	if len(matches) == 3 {
		return matches[1], matches[2]
	}
	return "", strings.TrimSpace(input)
}
