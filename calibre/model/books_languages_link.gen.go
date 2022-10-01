package model

const TableNameBooksLanguagesLink = "books_languages_link"

// BooksLanguagesLink mapped from table <books_languages_link>
type BooksLanguagesLink struct {
	ID        int32 `gorm:"column:id;type:INTEGER" json:"id"`
	Book      int32 `gorm:"column:book;type:INTEGER" json:"book"`
	LangCode  int32 `gorm:"column:lang_code;type:INTEGER" json:"lang_code"`
	ItemOrder int32 `gorm:"column:item_order;type:INTEGER" json:"item_order"`
}

// TableName BooksLanguagesLink's table name
func (*BooksLanguagesLink) TableName() string {
	return TableNameBooksLanguagesLink
}
