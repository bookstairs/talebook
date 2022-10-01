package main

import (
	"errors"
	"log"
	"os"

	"github.com/bookstairs/talebook/calibre"
	"github.com/bookstairs/talebook/config"
)

// initRuntime will create all the required working directory.
func initRuntime(c *config.ServerConfig) {
	// Internal method for making all the directories if it's not existed.
	createPath := func(path string) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
	// Internal method for checking if a file exists.
	fileExist := func(path string) bool {
		_, err := os.Stat(path)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			// We can't access this file because it's unreachable.
			log.Fatal(err)
		}
		return err == nil
	}

	// Create all the working directories if they are not existed.
	createPath(c.GetPath("statics"))  // Frontend files.
	createPath(c.GetPath("settings")) // The settings for talebook in sqlite3.
	createPath(c.GetPath("converts")) // The temporary directory for storing the convert files.
	createPath(c.GetPath("imports"))  // The directory for scanning and batch import books to your calibre library.
	createPath(c.GetPath("uploads"))  // The temporary directory for storing the uploaded books.
	createPath(c.LibraryPath)         // The calibre library directory, it may not exist for a new machine.

	// Check the calibredb
	if !fileExist(c.CalibreDB) {
		log.Fatal("No calibredb could be found on path:", c.CalibreDB, "check or provide a new path through cmd line.")
	}

	// Check the ebook-convert
	if !fileExist(c.Convert) {
		log.Fatal("No ebook-convert could be found on path:", c.Convert, "check or provide a new path through cmd line.")
	}

	// Extract the default calibre library in case of failure.
	if !fileExist(calibre.GetDatabase(c.LibraryPath)) {
		log.Print("No calibre library found extract the default library.")
		err := extractDefaultLibrary(c.LibraryPath)
		if err != nil {
			// Use log.Fatal may not execute the defer method.
			log.Fatal(err)
		}
	}
}
