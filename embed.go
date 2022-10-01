package main

import (
	"embed"
	"io/fs"
	"log"
)

// This should work after the `npm run build`
//
//go:embed app/dist/*
var frontend embed.FS

// This would be the default calibre library to use.
//
//go:embed library/*
var library embed.FS

// extractDefaultLibrary will extract the embedded calibre to outer filesystem.
func extractDefaultLibrary(path string) error {
	dir, err := fs.Sub(library, "library")
	if err != nil {
		return err
	}

	// TODO How to extract all files?
	log.Println(dir, path)
	return nil
}
