package calibre

import (
	"log"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var (
	// The calibre database path.
	currentPath string
	// We will hold this instance for caching the connection pool.
	db *gorm.DB
)

// database is the file for storing the calibre data
const database = "metadata.db"

// GetDatabase will return the full path of calibre db.
func GetDatabase(librayPath string) string {
	return filepath.Join(librayPath, database)
}

// DB will return a calibre connection pool.
func DB() *gorm.DB {
	if db == nil {
		log.Fatalln("Wrong calibre initialization")
	}

	return db
}

func Reconnect(libraryPath string) error {
	dbPath := GetDatabase(libraryPath)
	if dbPath == currentPath {
		// No need to reconnect the calibre.
		return nil
	} else {
		currentPath = dbPath
	}

	// Close old connection.
	connect, err := db.DB()
	if err != nil {
		return err
	}
	err = connect.Close()
	if err != nil {
		return err
	}

	// Create new connection.
	db, err = gorm.Open(sqlite.Open(currentPath))
	return err
}
