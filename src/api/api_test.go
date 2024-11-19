package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/luislve17/comaho/utils"
)

func TestGetMediaPath_NoEnvVar(t *testing.T) {
	os.Unsetenv("COMAHO_MEDIA_PATH")

	fsChecker := utils.CanReadDir
	mediaPath, err := getMediaPath(fsChecker)

	assert.ErrorIs(t, err, ErrEnvVarNotSet)
	assert.Equal(t, "null", mediaPath)
}

func TestGetMediaPath_UnreachablePath(t *testing.T) {
	tempDir := t.TempDir()

	err := os.Chmod(tempDir, 0000)
	assert.NoError(t, err, "Failed to set permissions for testing")

	os.Setenv("COMAHO_MEDIA_PATH", tempDir)

	mediaPath, err := getMediaPath(utils.CanReadDir)

	assert.Error(t, err, "Expected an error for an unreadable path")
	assert.Equal(t, tempDir, mediaPath, "Expected media path to be set, when path exists but is unreadable")
}

func TestGetMediaPath_ValidPath(t *testing.T) {
	tempDir := t.TempDir()
	os.Setenv("COMAHO_MEDIA_PATH", tempDir)

	mediaPath, err := getMediaPath(utils.CanReadDir)

	assert.NoError(t, err, "Expected no error for a valid path")
	assert.Equal(t, tempDir, mediaPath+"foo", "Expected media path to match the temporary directory")
}
