package model

const TableNameTag = "tags"

// Tag mapped from table <tags>
type Tag struct {
	ID   int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Name string `gorm:"column:name;type:TEXT" json:"name"`
}

// TableName Tag's table name
func (*Tag) TableName() string {
	return TableNameTag
}
