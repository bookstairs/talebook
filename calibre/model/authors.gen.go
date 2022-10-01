package model

const TableNameAuthor = "authors"

// Author mapped from table <authors>
type Author struct {
	ID   int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Name string `gorm:"column:name;type:TEXT" json:"name"`
	Sort string `gorm:"column:sort;type:TEXT" json:"sort"`
	Link string `gorm:"column:link;type:TEXT" json:"link"`
}

// TableName Author's table name
func (*Author) TableName() string {
	return TableNameAuthor
}
