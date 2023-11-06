package runtime

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func Handler(w http.ResponseWriter, r *http.Request) {

}

func LoadRoutes(mux *chi.Mux, theme *Theme) {

	SetupAPI(mux)

	RenderNotFound := func(w http.ResponseWriter) {

		w.WriteHeader(http.StatusNotFound)

		err := theme.Render(w, "404.html", D{
			"title": "Page not found",
		})

		if err != nil {
			log.Default().Println(err)
		}
	}
	/*
		mux.Get("/admin/*", func(w http.ResponseWriter, r *http.Request) {

			path := strings.TrimPrefix(r.RequestURI, "/admin/")

			if path == "" {
				path = "index.html"
			}

			WriteMimeType(w, path)

			path = "admin/dist/" + path

			data, _ := adminFS.ReadFile(path)

			w.Write(data)
		})
	*/
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {

		posts, _ := GetPosts()

		err := theme.Render(w, "index.html", D{
			"title":      BlogTitle("Home"),
			"blog_title": GetConfig("blog_title"),
			"posts":      posts,
		})

		if err != nil {
			log.Default().Println(err)
		}
	})

	mux.Get("/{slug}", func(w http.ResponseWriter, r *http.Request) {

		post, err := GetPostBySlug(chi.URLParam(r, "slug"))

		if err != nil {
			RenderNotFound(w)
			return
		}

		err = theme.Render(w, "post.html", D{
			"title":      BlogTitle(post.Title),
			"blog_title": GetConfig("blog_title"),
			"post":       post,
		})

		if err != nil {
			log.Default().Println(err)
		}
	})

	mux.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) {

		err := theme.RenderAsset(w, r.RequestURI)

		if err != nil {
			log.Default().Println(err)
		}
	})

	mux.Get("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?><urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"></urlset>`))
	})

	mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		RenderNotFound(w)
	})
}

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
