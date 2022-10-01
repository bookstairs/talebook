package calibre

const TableNameBooksAuthorsLink = "books_authors_link"

// BooksAuthorsLink mapped from table <books_authors_link>
type BooksAuthorsLink struct {
	ID     int32 `gorm:"column:id;type:INTEGER" json:"id"`
	Book   int32 `gorm:"column:book;type:INTEGER" json:"book"`
	Author int32 `gorm:"column:author;type:INTEGER" json:"author"`
}

// TableName BooksAuthorsLink's table name
func (*BooksAuthorsLink) TableName() string {
	return TableNameBooksAuthorsLink
}
