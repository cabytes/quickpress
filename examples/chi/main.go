package main

import (
	"cabytes/zine"
	"cabytes/zine/themes/light"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type MyUser struct{}

func (mu *MyUser) ID() string       { return "" }
func (mu *MyUser) Username() string { return "" }
func (mu *MyUser) Password() string { return "" }
func (mu *MyUser) Name() string     { return "" }

func main() {

	mux := chi.NewMux()

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/blog", http.StatusMovedPermanently)
	})

	app, err := zine.New(
		zine.DataPath("../../data"),
		zine.BaseHref("/blog"),
		zine.LoadTheme("../../themes/light/", light.Files),
	)

	if err != nil {
		panic(err)
	}

	mux.Handle("/blog*", app)

	http.ListenAndServe(":8000", mux)
}
