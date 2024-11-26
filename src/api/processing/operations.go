package processing

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/luislve17/comaho/utils"
)

func ConvertContent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		parsedURLData := utils.ParseURLPath(vars["name"])
		pathInfo := utils.GetContentPath(parsedURLData)
		sourcePath := filepath.Join(pathInfo, vars["item"])

		go ConvertComic2Ebook(sourcePath)

		w.WriteHeader(http.StatusOK)
		// w.Write([]byte("Conversion started. You will be notified when it's done."))
	}
}

func ConvertComic2Ebook(comicPath string) {
	outputFilePath := removeExtensions(comicPath) + ".kepub.epub"
	log.Printf("Expected output file: %s", outputFilePath)

	if _, err := os.Stat(outputFilePath); err == nil {
		log.Printf("File already exists, skipping conversion: %s", outputFilePath)
		return
	}

	command := fmt.Sprintf("kcc-c2e.py -p KoL -m '%s' -o '%s'", comicPath, outputFilePath)
	log.Printf("Attempting conversion: %s", command)
	cmd := exec.Command("sh", "-c", command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error during conversion: %v", err)
		return
	}
	log.Printf("Conversion completed: %s", string(output))
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
