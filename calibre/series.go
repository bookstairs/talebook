package calibre

import (
	"context"

	"crawshaw.io/sqlite"
)

// QuerySeriesCount will return the size of tags.
func QuerySeriesCount(ctx context.Context) (result int64, err error) {
	err = Execute(ctx, "SELECT COUNT(1) AS counts FROM series;", &ExecOptions{
		ResultFunc: func(stmt *sqlite.Stmt) error {
			result = stmt.GetInt64("counts")
			return nil
		},
	})

	return
}
