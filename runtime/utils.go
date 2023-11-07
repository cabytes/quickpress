package runtime

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path"
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

type EmbedFallbackFS struct {
	path    string
	embeded embed.FS
}

func NewEmbedFallbackFS(path string, embeded embed.FS) *EmbedFallbackFS {
	return &EmbedFallbackFS{path, embeded}
}

func (fe *EmbedFallbackFS) Open(name string) (f fs.File, err error) {

	fullPath := path.Clean(path.Join(fe.path, name))

	f, err = os.OpenFile(
		fullPath,
		os.O_RDONLY,
		os.ModePerm,
	)

	if err != nil {

		f, err = fe.embeded.Open(fullPath)

		if err != nil {

			f, err = fe.embeded.Open(name)

			if err != nil {
				return f, err
			}
		}
	}

	return
}

func (fe *EmbedFallbackFS) ReadDir(name string) (entries []fs.DirEntry, err error) {
	fullPath := path.Clean(path.Join(fe.path, name))
	entries, err = os.ReadDir(fullPath)

	if err != nil {
		return fe.embeded.ReadDir(name)
	}

	return
}

func (fe *EmbedFallbackFS) ReadFile(name string) (data []byte, err error) {
	f, err := fe.Open(name)
	if err != nil {
		return data, err
	}
	return io.ReadAll(f)
}
