package zine

import (
	"cabytes/zine/runtime"
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Opt interface {
	apply(app *ZineApp)
}

type User interface {
	ID() string
	Username() string
	Password() string
	Name() string
}

type ZineApp struct {
	theme    *runtime.Theme
	baseHref string
	dataPath string
}

func (za *ZineApp) setBaseHref(href string) {
	za.baseHref = href
}

func (za *ZineApp) setTheme(theme *runtime.Theme) {
	theme.SetFuncs(za.setTemplatingFuncs())
	za.theme = theme
}

func (za *ZineApp) setDataPath(path string) {
	os.MkdirAll(path, os.ModePerm)
	za.dataPath = path
}

func (za *ZineApp) setTemplatingFuncs() template.FuncMap {
	return template.FuncMap{

		"href": func(href string) string {
			return filepath.Clean(za.baseHref + "/" + href)
		},
	}
}

func (za *ZineApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.RequestURI == "/" || r.RequestURI == za.baseHref {

		posts, _ := runtime.GetPosts()

		data := runtime.D{
			"posts":      posts,
			"title":      runtime.BlogTitle("Home"),
			"blog_title": runtime.GetConfig("blog_title"),
		}

		if err := za.theme.Render(w, "index.html", data); err != nil {
			log.Default().Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	if strings.Index(r.RequestURI, za.baseHref+"/assets") == 0 {

		err := za.theme.RenderAsset(
			w,
			strings.TrimPrefix(r.RequestURI, za.baseHref),
		)

		if err != nil {
			log.Default().Println(err)
		}
	}

	if strings.Index(r.RequestURI, za.baseHref) == 0 && len(r.RequestURI) > len(za.baseHref) {

		slug := strings.TrimPrefix(r.RequestURI, za.baseHref+"/")

		post, err := runtime.GetPostBySlug(slug)

		if err != nil {
			http.NotFound(w, r)
			return
		}

		err = za.theme.Render(w, "post.html", runtime.D{
			"title":      runtime.BlogTitle(post.Title),
			"blog_title": runtime.GetConfig("blog_title"),
			"post":       post,
		})

		if err != nil {
			log.Default().Println(err)
		}
	}
}

func New(opts ...Opt) *ZineApp {
	app := &ZineApp{}
	for _, opt := range opts {
		opt.apply(app)
	}
	return app
}

type AuthHookFunc func(username, password string) User

type noopOpt struct{}

func (aho *noopOpt) apply(app *ZineApp) {}

func AuthHook(fn AuthHookFunc) Opt {
	return &noopOpt{}
}

type dataPathOpt struct{ path string }

func (opt *dataPathOpt) apply(app *ZineApp) {
	app.setDataPath(opt.path)
}

func DataPath(path string) Opt {
	return &dataPathOpt{path}
}

type loadThemeOpt struct {
	fs runtime.ThemeFS
}

func (opt *loadThemeOpt) apply(app *ZineApp) {
	app.setTheme(runtime.NewTheme(opt.fs))
}

func LoadTheme(path string, fallback embed.FS) Opt {
	return &loadThemeOpt{
		runtime.NewEmbedFallbackFS(path, fallback),
	}
}

type baseHrefOpt struct{ href string }

func (opt *baseHrefOpt) apply(app *ZineApp) {
	app.setBaseHref(opt.href)
}

func BaseHref(href string) Opt {
	return &baseHrefOpt{href}
}
