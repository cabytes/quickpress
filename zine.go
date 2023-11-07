package zine

import (
	"cabytes/zine/runtime"
	"database/sql"
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
	theme      *runtime.Theme
	baseHref   string
	dataPath   string
	repository *runtime.Repository
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

		posts, _ := za.repository.GetPosts()

		data := runtime.D{
			"posts":      posts,
			"title":      za.repository.BlogTitle("Home"),
			"blog_title": za.repository.GetConfig("blog_title"),
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

		return
	}

	if strings.Index(r.RequestURI, "/admin") == 0 {
		path := strings.TrimPrefix(r.RequestURI, "/admin")
		if path == "/" || path == "" {
			path = "index.html"
		}
		path = filepath.Clean("admin/dist/" + path)
		runtime.WriteMimeType(w, path)
		data, err := adminFS.ReadFile(path)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		w.Write(data)
		return
	}

	if strings.Index(r.RequestURI, za.baseHref) == 0 && len(r.RequestURI) > len(za.baseHref) {

		slug := strings.TrimPrefix(r.RequestURI, za.baseHref+"/")

		post, err := za.repository.GetPostBySlug(slug)

		if err != nil {
			http.NotFound(w, r)
			return
		}

		err = za.theme.Render(w, "post.html", runtime.D{
			"title":      za.repository.BlogTitle(post.Title),
			"blog_title": za.repository.GetConfig("blog_title"),
			"post":       post,
		})

		if err != nil {
			log.Default().Println(err)
		}
	}

}

func New(opts ...Opt) (app *ZineApp, err error) {

	app = &ZineApp{}

	for _, opt := range opts {
		opt.apply(app)
	}

	db, err := sql.Open("sqlite3", filepath.Clean(app.dataPath+"/db"))

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)

	repo, err := runtime.NewRepository(db)

	if err != nil {
		return nil, err
	}

	app.repository = repo

	return
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
