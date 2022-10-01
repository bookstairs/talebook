package calibre

import "path/filepath"

// database is the file for storing the calibre data
const database = "metadata.db"

// GetDatabase will return the full path of calibre db.
func GetDatabase(librayPath string) string {
	return filepath.Join(librayPath, database)
}
