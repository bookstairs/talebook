package model

const TableNameFeed = "feeds"

// Feed mapped from table <feeds>
type Feed struct {
	ID     int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Title  string `gorm:"column:title;type:TEXT" json:"title"`
	Script string `gorm:"column:script;type:TEXT" json:"script"`
}

// TableName Feed's table name
func (*Feed) TableName() string {
	return TableNameFeed
}
