package wp

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Load(mux *chi.Mux, theme *Theme) {

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {

		theme.Render(w, "index.html", D{
			"title": "Home",
			"theme": GetConfig("theme"),
		})
	})

	mux.Get("/{slug}", func(w http.ResponseWriter, r *http.Request) {
		// Check is a page or a post
		theme.Render(w, "post.html", D{
			"title": "Post",
			"slug":  chi.URLParam(r, "slug"),
		})
	})

	mux.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) {
		theme.RenderAsset(w, r.RequestURI)
	})

	mux.Route("/admin", AdminLoader)
}
