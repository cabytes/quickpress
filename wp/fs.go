package wp

import (
	"embed"
	"io"
	"io/fs"
	"os"
	"path"
)

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
		return fe.embeded.ReadDir(fullPath)
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
