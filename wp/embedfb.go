package wp

import (
	"embed"
	"io"
	"io/fs"
	"os"
	"path"
)

type FakeEmbedFallback struct {
	path    string
	embeded embed.FS
}

func NewFakeEmbedFallback(path string, embeded embed.FS) *FakeEmbedFallback {
	return &FakeEmbedFallback{path, embeded}
}

func (fe *FakeEmbedFallback) Open(name string) (f fs.File, err error) {

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

func (fe *FakeEmbedFallback) ReadDir(name string) (entries []fs.DirEntry, err error) {
	fullPath := path.Clean(path.Join(fe.path, name))
	return os.ReadDir(fullPath)
}

func (fe *FakeEmbedFallback) ReadFile(name string) (data []byte, err error) {
	f, err := fe.Open(name)
	if err != nil {
		return data, err
	}
	return io.ReadAll(f)
}
