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

		post, err := GetPostBySlug(chi.URLParam(r, "slug"))

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		theme.Render(w, "post.html", D{
			"title": "Post",
			"post":  post,
		})
	})

	mux.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) {
		theme.RenderAsset(w, r.RequestURI)
	})

	mux.Route("/admin", AdminLoader)
}
