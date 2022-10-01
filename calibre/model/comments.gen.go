package model

const TableNameComment = "comments"

// Comment mapped from table <comments>
type Comment struct {
	ID   int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Book int32  `gorm:"column:book;type:INTEGER" json:"book"`
	Text string `gorm:"column:text;type:TEXT" json:"text"`
}

// TableName Comment's table name
func (*Comment) TableName() string {
	return TableNameComment
}
