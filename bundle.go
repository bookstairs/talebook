package main

import (
	"embed"
)

// This should work after the `npm run build`
//
//go:embed app/dist/*
var bundle embed.FS
