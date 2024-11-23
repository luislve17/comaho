package contentpage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMediaPath_GenericPath(t *testing.T) {
	path := "/ABZ-123-Name_Including_Spaces"
	parsedURL := parseURLPath(path)

	assert.Equal(t, "ABZ", *parsedURL.Type, "Id type don't match")
	assert.Equal(t, "123", *parsedURL.ID, "Id value don't match")
	assert.Equal(t, "Name_Including_Spaces", parsedURL.Name, "Content name don't match")
}

func TestParseMediaPath_MissingId(t *testing.T) {
	path := "/-Name_Including_Spaces"
	parsedURL := parseURLPath(path)

	assert.Nil(t, parsedURL.Type, "Id type was supposed to be nil")
	assert.Nil(t, parsedURL.ID, "Id value was supposed to be nil")
	assert.Equal(t, "Name_Including_Spaces", parsedURL.Name, "Content name don't match")
}

func TestParseMediaPath_BadFormatting(t *testing.T) {
	path := "/Name_withBad-Format"
	parsedURL := parseURLPath(path)

	assert.Nil(t, parsedURL.Type, "Id type was supposed to be nil")
	assert.Nil(t, parsedURL.ID, "Id value was supposed to be nil")
	assert.Equal(t, "", parsedURL.Name, "Content name should be empty")
}
