package main

import (
	"cabytes/wordpost/wp"
	"embed"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//go:embed themes/clean/*
var cleanThemeFS embed.FS

var defaultTheme = wp.NewTheme(wp.NewFakeEmbedFallback("./themes/clean/", cleanThemeFS))

func main() {

	mux := chi.NewMux()

	wp.Load(mux, defaultTheme)

	http.ListenAndServe(":8085", mux)
}
