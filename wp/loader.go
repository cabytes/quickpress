package wp

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Load(mux *chi.Mux, theme *Theme) {

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
			w.WriteHeader(http.StatusNotFound)
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

	mux.Route("/admin", AdminLoader)
}
