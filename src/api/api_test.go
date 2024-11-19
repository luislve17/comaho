package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert" // Popular assertion library for cleaner test outputs
)

// TestGetMediaPath tests the getMediaPath function
func TestGetMediaPath(t *testing.T) {
	// Backup and defer restoration of COMAHO_MEDIA_PATH env variable
	originalPath := os.Getenv("COMAHO_MEDIA_PATH")
	defer os.Setenv("COMAHO_MEDIA_PATH", originalPath)

	// Set up the test environment
	t.Setenv("COMAHO_MEDIA_PATH", "/app/media")

	// Create a fake directory structure if necessary
	os.MkdirAll("/app/media", 0755)  // Mock directory creation
	defer os.RemoveAll("/app/media") // Cleanup after test

	// Execute the function
	path, err := getMediaPath()

	// Validate the results
	assert.NoError(t, err, "Expected no error when the media path exists and is accessible")
	assert.Equal(t, "/app/media", path, "Expected media path to match the environment variable")
}
