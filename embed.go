package main

import (
	"embed"
	"os"
	"path/filepath"
)

// This should work after the `npm run build`
//
//go:embed app/dist/*
var frontend embed.FS

// This would be the default calibre library to use.
//
//go:embed library/*
var library embed.FS

// extractDefaultLibrary will extract the embedded calibre library to outer filesystem.
func extractDefaultLibrary(path string) error {
	return extractFS("library", path)
}

// extractFS will reclusive extract the file and directories from the fs.FS to system directory filesystem.
func extractFS(source, target string) error {
	files, err := library.ReadDir(source)
	if err != nil {
		return err
	}

	// Create all the files.
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		read, err := library.ReadFile(source + "/" + name)
		if err != nil {
			return err
		}

		// Create the file.
		fo, err := os.Create(filepath.Join(target, name))
		if err != nil {
			return err
		}

		if _, err := fo.Write(read); err != nil {
			return err
		}

		// Close fo and check for its returned error
		if err := fo.Close(); err != nil {
			return err
		}
	}

	for _, dir := range files {
		if !dir.IsDir() {
			continue
		}

		name := dir.Name()
		nextSource := source + "/" + name
		nextTarget := filepath.Join(target, name)

		// Create directory.
		if err := os.MkdirAll(nextTarget, os.ModePerm); err != nil {
			return err
		}

		if err := extractFS(nextSource, nextTarget); err != nil {
			return err
		}
	}

	return nil
}
