package runtime

import (
	"net/http"
	"path/filepath"
)

func WriteMimeType(w http.ResponseWriter, path string) {

	h := w.Header()
	ct := "Content-Type"

	switch filepath.Ext(path) {
	case ".html":
		h.Add(ct, "text/html")
	case ".css":
		h.Add(ct, "text/css")
	case ".js":
		h.Add(ct, "text/javascript")
	case ".png":
		h.Add(ct, "image/png")
	case ".jpg":
		h.Add(ct, "image/jpg")
	}
}
