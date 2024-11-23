package api

import (
	"os"
	"testing"

	"github.com/luislve17/comaho/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetMediaPath_NoEnvVar(t *testing.T) {
	os.Unsetenv("COMAHO_MEDIA_PATH")
	os.Unsetenv("COMAHO_DOCKER_VOLUME_PATH")

	fsChecker := utils.CanReadDir
	mediaPath, err := getMediaPath(fsChecker)

	assert.ErrorIs(t, err, ErrEnvVarNotSet)
	assert.Equal(t, "null", mediaPath)
}

func TestGetMediaPath_UnreachablePath(t *testing.T) {
	os.Unsetenv("COMAHO_DOCKER_VOLUME_PATH")

	// Create temp directory and set permissions to make it unreadable
	tempDir := t.TempDir()
	err := os.Chmod(tempDir, 0000)
	assert.NoError(t, err, "Failed to set permissions for testing")

	// Temp directory should not be readable
	info, statErr := os.Stat(tempDir)
	assert.NoError(t, statErr, "TempDir should exist")
	assert.False(t, info.Mode().Perm()&0400 != 0, "Directory should not be readable")

	os.Setenv("COMAHO_MEDIA_PATH", tempDir)

	mediaPath, err := getMediaPath(utils.CanReadDir)

	assert.Error(t, err, "Expected an error for an unreadable path")
	assert.Equal(t, tempDir, mediaPath, "Expected media path to be set, when path exists but is unreadable")
}

func TestGetMediaPath_ValidPath(t *testing.T) {
	os.Unsetenv("COMAHO_DOCKER_VOLUME_PATH")

	tempDir := t.TempDir()
	os.Setenv("COMAHO_MEDIA_PATH", tempDir)

	mediaPath, err := getMediaPath(utils.CanReadDir)

	assert.NoError(t, err, "Expected no error for a valid path")
	assert.Equal(t, tempDir, mediaPath, "Expected media path to match the temporary directory")
}
