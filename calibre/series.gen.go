package calibre

const TableNameSeries = "series"

// Series mapped from table <series>
type Series struct {
	ID   int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Name string `gorm:"column:name;type:TEXT" json:"name"`
	Sort string `gorm:"column:sort;type:TEXT" json:"sort"`
}

// TableName Series's table name
func (*Series) TableName() string {
	return TableNameSeries
}
