package model

const TableNameCustomColumn = "custom_columns"

// CustomColumn mapped from table <custom_columns>
type CustomColumn struct {
	ID            int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Label         string `gorm:"column:label;type:TEXT" json:"label"`
	Name          string `gorm:"column:name;type:TEXT" json:"name"`
	Datatype      string `gorm:"column:datatype;type:TEXT" json:"datatype"`
	MarkForDelete string `gorm:"column:mark_for_delete;type:BOOL" json:"mark_for_delete"`
	Editable      string `gorm:"column:editable;type:BOOL" json:"editable"`
	Display       string `gorm:"column:display;type:TEXT" json:"display"`
	IsMultiple    string `gorm:"column:is_multiple;type:BOOL" json:"is_multiple"`
	Normalized    string `gorm:"column:normalized;type:BOOL" json:"normalized"`
}

// TableName CustomColumn's table name
func (*CustomColumn) TableName() string {
	return TableNameCustomColumn
}
