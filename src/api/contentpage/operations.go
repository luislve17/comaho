package contentpage

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/luislve17/comaho/utils"
	"github.com/tidwall/gjson"
)

func ServeContentPage(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		contentPageData := getContentPageData(r)
		err := utils.RenderTemplate(w, tmpl, "content-index", contentPageData)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	}
}

func getContentPageData(req *http.Request) ContentPageData {
	fullURL := req.URL.String()
	parsedURLData := parseURLPath(fullURL)
	refreshMetadata(parsedURLData)

	contentSummary := ContentSummary{
		ImgURL:    "",
		Name:      "",
		Published: "",
		Authors:   []string{},
		Genres:    []string{},
	}

	content := []string{}

	rawMetadata, metadataErr := readMetadataFromFile(parsedURLData)
	existentContent, contentErr := readAvailableContentFromFile(parsedURLData)

	if metadataErr != nil || rawMetadata == nil || contentErr != nil || existentContent == nil {
		log.Printf("Error reading metadata: %v||%v", metadataErr, contentErr)
		return ContentPageData{
			Summary: contentSummary,
			Content: content,
		}
	}

	contentSummary = ContentSummary{
		ImgURL:    gjson.Get(*rawMetadata, "data.images.jpg.image_url").Str,
		Name:      gjson.Get(*rawMetadata, "data.title").Str,
		Published: extractDate(gjson.Get(*rawMetadata, "data.published.from").Str),
		Authors:   getAuthorsNames(*rawMetadata),
		Genres:    getGenres(*rawMetadata),
	}
	content = existentContent
	return ContentPageData{
		Summary: contentSummary,
		Content: content,
	}
}

func parseURLPath(path string) ParsedURL {
	// Define regex to match the components of the URL
	log.Printf("Entered: %s", path)
	regex := regexp.MustCompile(`^/(?:([A-Z]+)-)?(?:([0-9]+)-)?-?([A-Za-z0-9_]+)$`)

	// Match the URL against the regex
	matches := regex.FindStringSubmatch(path)
	if len(matches) == 4 {
		var idType, id *string
		if matches[1] != "" {
			idType = &matches[1]
		}
		if matches[2] != "" {
			id = &matches[2]
		}
		name := matches[3]

		// Always return the name component (matches[3])
		return ParsedURL{
			Type: idType,
			ID:   id,
			Name: name,
		}
	}

	return ParsedURL{
		Type: nil,
		ID:   nil,
		Name: "",
	}
}

func refreshMetadata(parsedURLData ParsedURL) error {
	if parsedURLData.Type == nil || *parsedURLData.Type != "MAL" || parsedURLData.ID == nil {
		fmt.Println("Type is not MAL or ID is missing. Skipping metadata refresh.")
		return nil
	}

	folderName := fmt.Sprintf("(%s-%s) %s", *parsedURLData.Type, *parsedURLData.ID, parsedURLData.Name)
	metadataPath := filepath.Join("media", folderName, "metadata.json")

	if _, err := os.Stat(metadataPath); err == nil {
		log.Printf("Metadata file for %s already exist. Skipping API fetching.", metadataPath)
		return nil
	}

	apiData, err := fetchMALMetadata(*parsedURLData.ID)
	if err != nil {
		return fmt.Errorf("failed to fetch metadata: %w", err)
	}

	err = writeMetadataToFile(parsedURLData, apiData)
	if err != nil {
		return fmt.Errorf("failed to write metadata to file: %w", err)
	}

	fmt.Println("Metadata successfully refreshed.")
	return nil
}

func fetchMALMetadata(id string) ([]byte, error) {
	apiURL := fmt.Sprintf("https://api.jikan.moe/v4/manga/%s", id)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make API request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read API response: %w", err)
	}

	return body, nil
}

func writeMetadataToFile(parsedURLData ParsedURL, data []byte) error {
	// Construct the folder path
	folderName := fmt.Sprintf("(%s-%s) %s", *parsedURLData.Type, *parsedURLData.ID, parsedURLData.Name)
	mediaFolder := filepath.Join("media", folderName)

	// Check if the folder exists
	if _, err := os.Stat(mediaFolder); os.IsNotExist(err) {
		return fmt.Errorf("media folder %s does not exist", mediaFolder)
	}

	// Pretty-print the JSON
	var prettyData map[string]interface{}
	if err := json.Unmarshal(data, &prettyData); err != nil {
		return fmt.Errorf("failed to unmarshal JSON data: %w", err)
	}

	prettyJSON, err := json.MarshalIndent(prettyData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to pretty-print JSON: %w", err)
	}

	// Write the pretty JSON to a file
	outputFile := filepath.Join(mediaFolder, "metadata.json")
	err = os.WriteFile(outputFile, prettyJSON, 0644)
	if err != nil {
		return fmt.Errorf("failed to write metadata file: %w", err)
	}

	return nil
}

func readMetadataFromFile(parsedURLData ParsedURL) (*string, error) {
	if parsedURLData.Type == nil || parsedURLData.ID == nil {
		return nil, nil
	}

	folderName := fmt.Sprintf("(%s-%s) %s", *parsedURLData.Type, *parsedURLData.ID, parsedURLData.Name)
	metadataFilePath := filepath.Join("media", folderName, "metadata.json")

	if _, err := os.Stat(metadataFilePath); os.IsNotExist(err) {
		return nil, nil // Metadata file does not exist
	}

	data, err := os.ReadFile(metadataFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read metadata file: %w", err)
	}

	rawData := string(data)
	return &rawData, nil
}

func readAvailableContentFromFile(parsedURLData ParsedURL) ([]string, error) {
	if parsedURLData.Type == nil || parsedURLData.ID == nil {
		return nil, nil
	}

	folderName := fmt.Sprintf("(%s-%s) %s", *parsedURLData.Type, *parsedURLData.ID, parsedURLData.Name)
	contentPath := filepath.Join("media", folderName)

	if _, err := os.Stat(contentPath); os.IsNotExist(err) {
		return nil, err
	}

	supportedExtensions := map[string]bool{
		".zip": true,
		".cbz": true,
		".rar": true,
		".cbr": true,
	}

	var files []string

	err := filepath.Walk(contentPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Check if it's a file and has a valid extension
		if !info.IsDir() && supportedExtensions[strings.ToLower(filepath.Ext(info.Name()))] {
			files = append(files, filepath.Base(path))
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func getAuthorsNames(metadata string) []string {
	var authors []string
	gjson.Get(metadata, "data.authors").ForEach(func(_, author gjson.Result) bool {
		authors = append(authors, author.Get("name").String())
		return true
	})
	return authors
}

func extractDate(datetime string) string {
	parsedTime, err := time.Parse(time.RFC3339, datetime)
	if err != nil {
		return ""
	}
	return parsedTime.Format("2006-01-02")
}

func getGenres(metadata string) []string {
	var genres []string
	gjson.Get(metadata, "data.genres").ForEach(func(_, genre gjson.Result) bool {
		genres = append(genres, genre.Get("name").String())
		return true
	})
	return genres
}
