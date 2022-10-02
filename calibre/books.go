package calibre

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/golang-module/carbon/v2"
	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"

	"github.com/bookstairs/talebook/model"
)

const (
	// sqlite3 time format. 2014-08-24 10:10:26.734752+00:00
	sqlite3TimeLayout = "2006-01-02 15:04:05.999999-07:00"

	// This is a common query prefix for select the most information from the calibre sqlite by using left join.
	bookDetailQueryTmpl = `SELECT b.id,
       b.title,
       r.rating,
       b.timestamp,
       b.pubdate,
       b.author_sort,
       p.name      as publisher,
       c.text      as comments,
       s.name      as series,
       l.lang_code as language,
       b.isbn
FROM books b
         LEFT JOIN books_ratings_link brl ON b.id = brl.book
         LEFT JOIN ratings r ON brl.rating = r.id
         LEFT JOIN books_publishers_link bpl ON b.id = bpl.book
         LEFT JOIN publishers p ON bpl.publisher = p.id
         LEFT JOIN comments c ON b.id = c.book
         LEFT JOIN books_series_link bsl ON b.id = bsl.book
         LEFT JOIN series s ON bsl.series = s.id
         LEFT JOIN books_languages_link bll ON b.id = bll.book
         LEFT JOIN languages l ON bll.lang_code = l.id`

	// Query all the authors for given books.
	bookAuthorQueryTmpl = `SELECT bal.book AS book_id, a.name AS author
FROM books_authors_link bal
         LEFT JOIN authors a ON bal.author = a.id
WHERE bal.book IN (%s);`

	// Query all the tags for given books.
	bookTagQueryTmpl = `SELECT btl.book AS book_id, t.name AS tag
FROM books_tags_link btl
         LEFT JOIN tags t ON btl.tag = t.id
WHERE btl.book IN (%s);`

	// Query all the formats for the given books.
	bookFormatsQueryImpl = `SELECT b.id AS id,
       b.title AS title,
       b.path || '/' || d.name || '.' || LOWER(d.format) AS path,
       d.format AS format
FROM books b
         LEFT JOIN data d ON b.id = d.book
WHERE b.id IN (%s);`
)

// QueryRandomBookIDs will return random book ids from calibre.
func QueryRandomBookIDs(ctx context.Context, size int) (ids []string, err error) {
	ids = make([]string, 0, size)
	err = Execute(ctx, "SELECT id FROM books ORDER BY RANDOM() LIMIT ?;", &sqlitex.ExecOptions{
		Args: []any{size},
		ResultFunc: func(stmt *sqlite.Stmt) error {
			ids = append(ids, stmt.GetText("id"))
			return nil
		},
	})

	return
}

// QueryBookCount will return the size of books.
func QueryBookCount(ctx context.Context) (result int64, err error) {
	err = Execute(ctx, "SELECT COUNT(1) AS counts FROM books;", &sqlitex.ExecOptions{
		ResultFunc: func(stmt *sqlite.Stmt) error {
			result = stmt.GetInt64("counts")
			return nil
		},
	})

	return
}

// QueryBooksByIDs query the books by given ids.
func QueryBooksByIDs(ctx context.Context, ids []string) (books []model.Book, err error) {
	if len(ids) == 0 {
		return []model.Book{}, nil
	}

	books = make([]model.Book, 0, len(ids))
	query := bookDetailQueryTmpl + " WHERE b.id in (" + strings.Join(ids, ", ") + ");"
	err = Execute(ctx, query, &sqlitex.ExecOptions{
		ResultFunc: func(stmt *sqlite.Stmt) error {
			books = append(books, convertBookDetailQuery(stmt))
			return nil
		},
	})
	if err == nil {
		books, err = bookMetadataQuery(ctx, books)
	}

	return
}

// QueryBooks is used for the page query.
func QueryBooks(ctx context.Context, index, size int) (books []model.Book, err error) {
	if index < 1 || size < 1 {
		err = fmt.Errorf("invalid index %d or size %d, they should exceed 0", index, size)
		return
	}

	books = make([]model.Book, 0, size)
	query := bookDetailQueryTmpl + ` ORDER BY b.id DESC LIMIT ? OFFSET ?;`
	err = Execute(ctx, query, &sqlitex.ExecOptions{
		Args: []any{size, (index - 1) * size},
		ResultFunc: func(stmt *sqlite.Stmt) error {
			books = append(books, convertBookDetailQuery(stmt))
			return nil
		},
	})
	if err == nil {
		books, err = bookMetadataQuery(ctx, books)
	}

	return
}

// QueryBookFormats will query the formats for the bookFormatsQueryImpl.
func QueryBookFormats(ctx context.Context, bookID int64) (result []model.BookFormat, err error) {
	// Query book formats.
	formatsQuery := fmt.Sprintf(bookFormatsQueryImpl, strconv.FormatInt(bookID, 10))
	err = Execute(ctx, formatsQuery, &sqlitex.ExecOptions{
		ResultFunc: func(stmt *sqlite.Stmt) error {
			result = append(result, model.BookFormat{
				ID:     stmt.GetInt64("id"),
				Title:  stmt.GetText("title"),
				Format: stmt.GetText("format"),
				Path:   stmt.GetText("path"),
			})
			return nil
		},
	})

	return
}

// convertBookDetailQuery will convert the bookDetailQueryTmpl into a book model.
func convertBookDetailQuery(stmt *sqlite.Stmt) model.Book {
	return model.Book{
		ID:         stmt.GetInt64("id"),
		Title:      stmt.GetText("title"),
		Rating:     int(stmt.GetInt64("rating")),
		Timestamp:  carbon.ParseByLayout(stmt.GetText("timestamp"), sqlite3TimeLayout).ToDateString(),
		Pubdate:    carbon.ParseByLayout(stmt.GetText("pubdate"), sqlite3TimeLayout).ToDateString(),
		AuthorSort: stmt.GetText("author_sort"),
		Publisher:  stmt.GetText("publisher"),
		Comments:   stmt.GetText("comments"),
		Series:     stmt.GetText("series"),
		Language:   stmt.GetText("language"),
		Isbn:       stmt.GetText("isbn"),
	}
}

// bookMetadataQuery will query the metadata for the bookDetailQueryTmpl.
// We will add authors, tags, img, thumb, author_url and publisher_url.
// TODO collector, count_visit and count_download will be added later. We will set default value currently.
func bookMetadataQuery(ctx context.Context, books []model.Book) ([]model.Book, error) {
	ids := make([]string, 0, len(books))
	bookIdx := make(map[string]*model.Book, len(books))
	for i := range books {
		ref := &books[i]
		id := strconv.FormatInt(ref.ID, 10)
		ids = append(ids, id)
		bookIdx[id] = ref
	}
	idStr := strings.Join(ids, ", ")

	// Query authors.
	authorQuery := fmt.Sprintf(bookAuthorQueryTmpl, idStr)
	err := Execute(ctx, authorQuery, &sqlitex.ExecOptions{
		ResultFunc: func(stmt *sqlite.Stmt) error {
			id := stmt.GetText("book_id")
			bookIdx[id].Authors = append(bookIdx[id].Authors, stmt.GetText("author"))
			return nil
		},
	})
	if err != nil {
		return nil, err
	}

	// Query tags.
	tagQuery := fmt.Sprintf(bookTagQueryTmpl, idStr)
	err = Execute(ctx, tagQuery, &sqlitex.ExecOptions{
		ResultFunc: func(stmt *sqlite.Stmt) error {
			id := stmt.GetText("book_id")
			bookIdx[id].Tags = append(bookIdx[id].Tags, stmt.GetText("tag"))
			return nil
		},
	})
	if err != nil {
		return nil, err
	}

	// Setting the fixed books metadata
	for i := range books {
		id := strconv.FormatInt(books[i].ID, 10)
		ts := strconv.FormatInt(carbon.Now().Timestamp(), 10)

		books[i].Author = strings.Join(books[i].Authors, " / ")
		books[i].Tag = strings.Join(books[i].Tags, ", ")
		books[i].Img = "/get/cover/" + id + ".jpg?t=" + ts
		books[i].Thumb = "/get/thumb_60x80/" + id + ".jpg?t=" + ts
		if len(books[i].Authors) > 0 {
			books[i].AuthorURL = "/author/" + books[i].Authors[0]
		}
		if books[i].Publisher != "" {
			books[i].PublisherURL = "/publisher/" + books[i].Publisher
		}

		books[i].Collector = "admin"
	}

	return books, nil
}
