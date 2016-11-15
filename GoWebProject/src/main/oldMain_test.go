package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetContentType_KnownExtension_ReturnsMimeType(t *testing.T)  {
	file := "video.mp4"
	actualContentType := getContentType(file)
	expectedContentType := extensionToContentType[".mp4"]

	assert.Equal(t, expectedContentType, actualContentType)
}

func TestGetContentType_UnknownExtension_ReturnsMimeType(t *testing.T)  {
	file := "video.ElCiupos"
	actualContentType := getContentType(file)

	assert.Equal(t, defaultContentType, actualContentType)
}
