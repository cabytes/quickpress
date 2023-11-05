package wp

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Load(mux *chi.Mux, theme *Theme) {

	RenderNotFound := func(w http.ResponseWriter) {

		w.WriteHeader(http.StatusNotFound)

		err := theme.Render(w, "404.html", D{
			"title": "Page not found",
		})

		if err != nil {
			log.Default().Println(err)
		}
	}

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {

		posts, _ := GetPosts()

		err := theme.Render(w, "index.html", D{
			"title": "Home",
			"posts": posts,
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
			"title": "Post",
			"post":  post,
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

	mux.Route("/admin", AdminLoader)
}
