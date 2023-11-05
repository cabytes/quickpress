package main

import (
	"cabytes/wordpost/wp"
	"embed"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//go:embed themes/clean/*
var cleanThemeFS embed.FS

var defaultTheme = wp.NewTheme(wp.NewFakeEmbedFallback("./themes/clean/", cleanThemeFS))

func main() {

	mux := chi.NewMux()

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		defaultTheme.Render(w, "index.html", wp.D{
			"title": "Homepage",
		})
	})

	mux.Get("/{post}", func(w http.ResponseWriter, r *http.Request) {
		defaultTheme.Render(w, "post.html", wp.D{
			"title": "Post",
			"slug":  chi.URLParam(r, "post"),
		})
	})

	mux.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) {
		defaultTheme.RenderAsset(w, r.RequestURI)
	})

	mux.Route("/admin", wp.AdminLoader)

	fmt.Println("Working with theme:", defaultTheme.Name())

	http.ListenAndServe(":8085", mux)
}
