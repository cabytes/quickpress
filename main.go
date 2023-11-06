package main

import (
	. "cabytes/wordpost/wp"
	"embed"
)

//go:embed themes/light/*
var cleanThemeFS embed.FS

//go:embed admin/dist/*
var adminFS embed.FS

func main() {

	SetDefaultTheme(
		NewTheme(
			NewFakeEmbedFallback(
				"./themes/light/",
				cleanThemeFS,
			),
		),
	)

	SetAdminFS(adminFS)

	SetupCLI()
}
