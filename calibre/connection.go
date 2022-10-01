package calibre

import (
	"log"
	"path/filepath"

	"zombiezen.com/go/sqlite"
)

var (
	// The calibre database path.
	currentPath string
	// We will hold this instance for caching the connection pool.
	db *sqlite.Conn
)

// database is the file for storing the calibre data
const database = "metadata.db"

// GetDatabase will return the full path of calibre db.
func GetDatabase(librayPath string) string {
	return filepath.Join(librayPath, database)
}

// DB will return a calibre connection pool.
// We use a pure-go implementation of the https://github.com/crawshaw/sqlite for better performance.
func DB() *sqlite.Conn {
	if db == nil {
		log.Fatalln("Wrong calibre initialization.")
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

	// Close old connection if exists.
	if db != nil {
		if err := db.Close(); err != nil {
			return err
		}
	}

	// Create new connection.
	conn, err := sqlite.OpenConn(currentPath, sqlite.OpenReadOnly, sqlite.OpenSharedCache)
	if err != nil {
		return err
	}
	db = conn

	return err
}
