package processing

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/luislve17/comaho/utils"
)

func DownloadContent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parsedURLData := utils.ParseURLPath(r.PathValue("name"))
		pathInfo := utils.GetContentPath(parsedURLData)
		sourcePath := filepath.Join(pathInfo, r.PathValue("item"))
		downloadPath := ConvertComic2Ebook(sourcePath)
		log.Printf("Downloading: %s", downloadPath)

		if _, err := os.Stat(downloadPath); os.IsNotExist(err) {
			log.Printf("File not found: %s", downloadPath)
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}

		w.Header().Set("HX-Redirect", downloadPath)
		w.WriteHeader(http.StatusOK)
	}
}

func ConvertComic2Ebook(comicPath string) string {
	outputFilePath := removeExtensions(comicPath) + ".kepub.epub"
	log.Printf("Expected output file: %s", outputFilePath)
	command := fmt.Sprintf("kcc-c2e.py -p KoL -m '%s' -o '%s'", comicPath, outputFilePath)
	log.Printf("Attempting conversion: %s", command)
	cmd := exec.Command("sh", "-c", command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error: %v", err)
		return ""
	}
	log.Printf(string(output))
	return outputFilePath
}

func removeExtensions(filePath string) string {
	for {
		ext := filepath.Ext(filePath)
		if ext == "" {
			break
		}
		filePath = filePath[:len(filePath)-len(ext)]
	}
	return filePath
}
