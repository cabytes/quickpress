package wp

import (
	"embed"
	"io"
	"strings"
	"testing"
)

//go:embed test/*
var adminFS embed.FS

func TestFallbackToEmbed(t *testing.T) {

	efb := NewFakeEmbedFallback("./test", adminFS)

	f, err := efb.Open("index.html")

	if err != nil {
		t.Error(err.Error())
		return
	}

	data, err := io.ReadAll(f)

	if err != nil {
		t.Error(err.Error())
		return
	}

	if strings.IndexAny(string(data), "html") != 0 {
		t.Error("Expected content")
		return
	}
}
