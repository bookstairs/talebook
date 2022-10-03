package config

import (
	"log"
	"os"
	"path/filepath"
)

const DefaultCoverPath = "/cover/default.jpg"

// ServerConfig is used to start the fiber server.
type ServerConfig struct {
	Port        int    // The binding port for backend server.
	WorkingPath string // The working directory for talebook.
	LibraryPath string // The calibre library directory.
	EncryptKey  string // This is used to encrypt the cookie.
	Limit       int    // Allowed request per seconds.
	CalibreDB   string // The executable file calibredb for adding books.
	Convert     string // The executable file ebook-convert for converting books.
	CoverCache  int    // The cache in memory for storing frequently accessed files, such as the book cover.
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
	w = filepath.Join(w, "repository")

	return &ServerConfig{
		Port:        8000,
		WorkingPath: w,
		LibraryPath: DefaultLibraryPath(w),
		EncryptKey:  "this-is-an-encrypt-key",
		Limit:       100,
		CalibreDB:   defaultCalibreDB,
		Convert:     defaultConvert,
		CoverCache:  0,
	}
}

func DefaultLibraryPath(workingPath string) string {
	return filepath.Join(workingPath, "library")
}
