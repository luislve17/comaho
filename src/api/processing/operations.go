package processing

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/luislve17/comaho/utils"
)

func DownloadContent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parsedURLData := utils.ParseURLPath(r.PathValue("name"))
		contentPath := utils.GetContentPath(parsedURLData)
		downloadPath := filepath.Join(contentPath, r.PathValue("item"))
		log.Printf("Downloading: %s", downloadPath)

		if _, err := os.Stat(downloadPath); os.IsNotExist(err) {
			log.Printf("File not found: %s", downloadPath)
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}

		// w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(downloadPath))
		// w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("HX-Redirect", downloadPath)
		// http.ServeFile(w, r, downloadPath)
		w.WriteHeader(http.StatusOK)
	}
}
