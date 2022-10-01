package model

const TableNameBooksPublishersLink = "books_publishers_link"

// BooksPublishersLink mapped from table <books_publishers_link>
type BooksPublishersLink struct {
	ID        int32 `gorm:"column:id;type:INTEGER" json:"id"`
	Book      int32 `gorm:"column:book;type:INTEGER" json:"book"`
	Publisher int32 `gorm:"column:publisher;type:INTEGER" json:"publisher"`
}

// TableName BooksPublishersLink's table name
func (*BooksPublishersLink) TableName() string {
	return TableNameBooksPublishersLink
}
