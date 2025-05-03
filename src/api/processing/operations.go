package processing

import (
	"fmt"
	"io"
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

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
	}
}

func CheckConvertedContent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		parsedURLData := utils.ParseURLPath(vars["name"])
		pathInfo := utils.GetContentPath(parsedURLData)
		sourcePath := filepath.Join(pathInfo, vars["item"])
		outputFilePath := removeExtensions(sourcePath) + ".kepub.epub"

		if _, err := os.Stat(outputFilePath); err == nil {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
			return
		}

		w.WriteHeader(http.StatusNotFound)
	}
}

func DownloadConvertedContent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		parsedURLData := utils.ParseURLPath(vars["name"])
		pathInfo := utils.GetContentPath(parsedURLData)
		sourcePath := filepath.Join(pathInfo, vars["item"])

		log.Printf("source path: %s", sourcePath)

		outputFilePath := removeExtensions(sourcePath) + ".kepub.epub"

		if _, err := os.Stat(outputFilePath); os.IsNotExist(err) {
			// Log if the file is not found
			log.Printf("File not found: %s", outputFilePath)
			http.Error(w, "Converted file not found", http.StatusNotFound)
			return
		}

		log.Printf("Found file to download: %s", outputFilePath)

		w.Header().Set("Content-Type", "application/epub+zip") // Adjust this to the appropriate content type
		w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(outputFilePath))

		// Open the file and copy its contents to the response
		file, err := os.Open(outputFilePath)
		if err != nil {
			log.Printf("Error opening file: %v", err)
			http.Error(w, "Error opening file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		fileInfo, err := file.Stat()
		if err != nil {
			log.Printf("Error getting file size: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

		_, err = io.Copy(w, file)
		if err != nil {
			log.Printf("Error writing file content: %v", err)
			http.Error(w, "Error writing file", http.StatusInternalServerError)
			return
		}

		// File download initiated successfully
		log.Printf("File download initiated: %s", outputFilePath)
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

func getFileSize(filePath string) int64 {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Printf("Error getting file size: %v", err)
		return 0
	}
	return fileInfo.Size()
}
