package main

import (
	"cabytes/wordpost/wp"
	. "cabytes/wordpost/wp"
)

func main() {

	SetDefaultTheme(
		NewTheme(
			NewEmbedFallbackFS(
				"./themes/light/",
				LightThemeFS,
			),
		),
	)

	SetAdminFS(AdminFS)

	wp.Run()

	quickpress.Listen
}
