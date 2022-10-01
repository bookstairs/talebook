package model

const TableNamePublisher = "publishers"

// Publisher mapped from table <publishers>
type Publisher struct {
	ID   int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Name string `gorm:"column:name;type:TEXT" json:"name"`
	Sort string `gorm:"column:sort;type:TEXT" json:"sort"`
}

// TableName Publisher's table name
func (*Publisher) TableName() string {
	return TableNamePublisher
}
