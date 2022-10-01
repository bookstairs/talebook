package calibre

const TableNameBooksTagsLink = "books_tags_link"

// BooksTagsLink mapped from table <books_tags_link>
type BooksTagsLink struct {
	ID   int32 `gorm:"column:id;type:INTEGER" json:"id"`
	Book int32 `gorm:"column:book;type:INTEGER" json:"book"`
	Tag  int32 `gorm:"column:tag;type:INTEGER" json:"tag"`
}

// TableName BooksTagsLink's table name
func (*BooksTagsLink) TableName() string {
	return TableNameBooksTagsLink
}
