package main

import "embed"

//go:embed themes/light/*
var LightThemeFS embed.FS

//go:embed admin/dist/*
var AdminFS embed.FS
