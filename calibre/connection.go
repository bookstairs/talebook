package calibre

import (
	"log"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// database is the file for storing the calibre data
const database = "metadata.db"

// GetDatabase will return the full path of calibre db.
func GetDatabase(librayPath string) string {
	return filepath.Join(librayPath, database)
}

// Connect to calibre library database for querying.
func Connect(libraryPath string) *gorm.DB {
	dbPath := GetDatabase(libraryPath)
	db, err := gorm.Open(sqlite.Open(dbPath))
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
