package calibre

const TableNameBooksRatingsLink = "books_ratings_link"

// BooksRatingsLink mapped from table <books_ratings_link>
type BooksRatingsLink struct {
	ID     int32 `gorm:"column:id;type:INTEGER" json:"id"`
	Book   int32 `gorm:"column:book;type:INTEGER" json:"book"`
	Rating int32 `gorm:"column:rating;type:INTEGER" json:"rating"`
}

// TableName BooksRatingsLink's table name
func (*BooksRatingsLink) TableName() string {
	return TableNameBooksRatingsLink
}
