package utils

import (
	"fmt"
	"log"
	"os"
)

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

	return nil // Directory is readable
}
