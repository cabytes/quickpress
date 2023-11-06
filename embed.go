package zine

import "embed"

//go:embed themes/light/*
var lightThemeFS embed.FS

//go:embed admin/dist/*
var adminFS embed.FS
