package wp

import (
	"archive/zip"
	"embed"
	"encoding/json"
	"html/template"
	"io"
	"io/fs"
	"net/http"
	"path/filepath"
	"plugin"
	"strings"
)

type D map[string]any

type ThemeFS interface {
	fs.FS
	fs.ReadDirFS
	fs.ReadFileFS
}

type Theme struct {
	fs ThemeFS
}

func (t *Theme) ReadMetadata() (meta map[string]any, err error) {

	meta = make(map[string]any)

	f, err := t.fs.Open("metadata.json")

	if err != nil {
		return meta, err
	}

	data, err := io.ReadAll(f)

	if err != nil {
		return meta, err
	}

	err = json.Unmarshal(data, &meta)

	return
}

func (t *Theme) Name() string {

	meta, _ := t.ReadMetadata()

	if v, exists := meta["name"]; exists {
		return v.(string)
	}

	return "N/A"
}

func (t *Theme) RenderAsset(w http.ResponseWriter, asset string) error {

	WriteMimeType(w, asset)

	return t.Render(w, strings.TrimPrefix(asset, "/"), nil)
}

func (t *Theme) Render(w io.Writer, view string, data map[string]any) error {

	f, err := t.fs.Open(view)

	if err != nil {
		return err
	}

	tpld, err := io.ReadAll(f)

	if err != nil {
		return err
	}

	entries, err := t.fs.ReadDir("partials")

	if err != nil {
		return err
	}

	var tpl = template.New("")

	for _, entry := range entries {
		if filepath.Ext(entry.Name()) == ".html" {

			name := entry.Name()

			tpl = tpl.New(
				filepath.Base(name),
			)

			partialData, err := t.fs.ReadFile("partials/" + name)

			if err != nil {
				panic(err)
			}

			nTpl, err := tpl.Parse(string(partialData))

			if err != nil {
				panic(err)
			}

			tpl = nTpl
		}
	}

	tpl, err = tpl.Parse(string(tpld))

	if err != nil {
		return err
	}

	return tpl.Execute(w, data)
}

func NewTheme(themeFS ThemeFS) *Theme {
	return &Theme{themeFS}
}

func NewThemeFromPlugin(path string) (t *Theme, err error) {

	plug, err := plugin.Open(path)

	if err != nil {
		return t, err
	}

	sym, err := plug.Lookup("Files")

	if err != nil {
		return t, err
	}

	files := sym.(*embed.FS)

	t = NewTheme(NewFakeEmbedFallback("./themes/clean/", *files))

	return
}

type ThemeZipFS struct {
	*zip.ReadCloser
}

func (tzfs *ThemeZipFS) ReadDir(name string) (entries []fs.DirEntry, err error) {
	return
}

func (tzfs *ThemeZipFS) ReadFile(name string) (data []byte, err error) {
	return
}

func NewThemeFromZip(path string) (t *Theme, err error) {

	r, err := zip.OpenReader(path)

	if err != nil {
		return t, err
	}

	t = NewTheme(&ThemeZipFS{r})

	return
}
