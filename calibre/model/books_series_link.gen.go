package model

const TableNameBooksSeriesLink = "books_series_link"

// BooksSeriesLink mapped from table <books_series_link>
type BooksSeriesLink struct {
	ID     int32 `gorm:"column:id;type:INTEGER" json:"id"`
	Book   int32 `gorm:"column:book;type:INTEGER" json:"book"`
	Series int32 `gorm:"column:series;type:INTEGER" json:"series"`
}

// TableName BooksSeriesLink's table name
func (*BooksSeriesLink) TableName() string {
	return TableNameBooksSeriesLink
}
