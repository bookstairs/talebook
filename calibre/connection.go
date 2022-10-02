package calibre

import (
	"context"
	"log"
	"path/filepath"

	"crawshaw.io/sqlite"
	"crawshaw.io/sqlite/sqlitex"
)

var (
	// The calibre library path.
	currentPath string
	// We will hold this instance for caching the connection pool.
	pool *sqlitex.Pool
)

// database is the file for storing the calibre data.
const database = "metadata.db"

// poolSize used to cache the connection.
const poolSize = 30

type ExecOptions struct {
	Args       []any
	ResultFunc func(stmt *sqlite.Stmt) error
}

// GetDatabase will return the full path of calibre db.
func GetDatabase(librayPath string) string {
	return filepath.Join(librayPath, database)
}

// BorrowDB will return a calibre connection pool. Remember to return it after using it.
// We use a pure-go implementation of the https://github.com/crawshaw/sqlite for better performance.
func BorrowDB(ctx context.Context) *sqlite.Conn {
	if pool == nil {
		log.Fatalln("Wrong calibre initialization.")
	}

	return pool.Get(ctx)
}

// ReturnDB will put the connection back to the pool.
func ReturnDB(conn *sqlite.Conn) {
	pool.Put(conn)
}

// Reconnect will change the calibre connection to a new space.
func Reconnect(libraryPath string) error {
	if libraryPath == currentPath {
		// No need to reconnect the calibre.
		return nil
	} else {
		currentPath = libraryPath
	}

	// Close old connection if exists.
	if pool != nil {
		if err := pool.Close(); err != nil {
			return err
		}
	}

	// Create new connection.
	db := GetDatabase(libraryPath)
	flags := sqlite.SQLITE_OPEN_READONLY | sqlite.SQLITE_OPEN_SHAREDCACHE
	conn, err := sqlitex.Open(db, flags, poolSize)
	if err != nil {
		return err
	}
	pool = conn

	return err
}

// Execute will execute the sqlite query under the connection pool.
func Execute(ctx context.Context, query string, opts *ExecOptions) error {
	conn := BorrowDB(ctx)
	defer ReturnDB(conn)

	return sqlitex.Exec(conn, query, opts.ResultFunc, opts.Args...)
}
