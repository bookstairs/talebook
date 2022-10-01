package calibre

const TableNameBooksPluginDatum = "books_plugin_data"

// BooksPluginDatum mapped from table <books_plugin_data>
type BooksPluginDatum struct {
	ID   int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Book int32  `gorm:"column:book;type:INTEGER" json:"book"`
	Name string `gorm:"column:name;type:TEXT" json:"name"`
	Val  string `gorm:"column:val;type:TEXT" json:"val"`
}

// TableName BooksPluginDatum's table name
func (*BooksPluginDatum) TableName() string {
	return TableNameBooksPluginDatum
}
