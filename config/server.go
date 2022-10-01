package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/bookstairs/talebook/calibre"
)

type ServerConfig struct {
	Port        int    // The binding port for backend server.
	WorkingPath string // The working directory for talebook.
	LibraryPath string // The calibre library directory.
	EncryptKey  string // This is used to encrypt the cookie.
	Limit       int    // Allowed request per seconds.
	CalibreDB   string // The executable file calibredb for adding books.
	Convert     string // The executable file ebook-convert for converting books.
	Debug       bool   // Enable debug log and metrics monitor and anything else.
}

func (c *ServerConfig) GetPath(paths ...string) string {
	return filepath.Join(c.WorkingPath, filepath.Join(paths...))
}

func DefaultSeverConfig() *ServerConfig {
	// Init the config variables with some default values.
	w, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return &ServerConfig{
		Port:        8000,
		WorkingPath: w,
		LibraryPath: filepath.Join(w, "library"),
		EncryptKey:  "this-is-an-encrypt-key",
		Limit:       100,
		CalibreDB:   calibre.DefaultCalibreDB,
		Convert:     calibre.DefaultConvert,
		Debug:       false,
	}
}
