package main

import (
	"cabytes/wordpost/wp"
	"embed"
)

//go:embed themes/clean/*
var cleanThemeFS embed.FS

func main() {
	wp.SetDefaultTheme(wp.NewTheme(wp.NewFakeEmbedFallback("./themes/clean/", cleanThemeFS)))
	wp.SetupCLI()
}
