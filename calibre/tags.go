package calibre

import (
	"context"

	"crawshaw.io/sqlite"
)

// QueryTagCount will return the size of tags.
func QueryTagCount(ctx context.Context) (result int64, err error) {
	err = Execute(ctx, "SELECT COUNT(1) AS counts FROM tags;", &ExecOptions{
		ResultFunc: func(stmt *sqlite.Stmt) error {
			result = stmt.GetInt64("counts")
			return nil
		},
	})

	return
}
