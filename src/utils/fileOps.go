package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type ParsedURL struct {
	Type *string
	ID   *string
	Name string
}

func CanReadDir(path string) error {
	info, err := os.Stat(path)
	log.Printf("Attempting accessing %s\n", path)
	if err != nil {
		return fmt.Errorf("error accessing path: %w", err)
	}

	if !info.IsDir() {
		return fmt.Errorf("path is not a directory")
	}

	_, err = os.Open(path)
	if err != nil {
		return fmt.Errorf("directory is not readable: %w", err)
	}

	log.Printf("Directory is readable %s\n", path)
	return nil
}

func ParseDirPath(path string) ParsedURL {
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

func ParseURLPath(path string) ParsedURL {
	log.Printf("Entered: %s", path)
	regex := regexp.MustCompile(`^([A-Za-z]+)?-(\d+)?-?([A-Za-z0-9_]+)$`)

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

func GetContentPath(parsedURLData ParsedURL) string {
	var contentPath string
	if parsedURLData.Type == nil || parsedURLData.ID == nil {
		contentPath = filepath.Join("media", parsedURLData.Name)
	} else {

		folderName := fmt.Sprintf("(%s-%s) %s", *parsedURLData.Type, *parsedURLData.ID, parsedURLData.Name)
		contentPath = filepath.Join("media", folderName)
	}
	return contentPath
}
