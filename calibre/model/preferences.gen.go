package model

const TableNamePreference = "preferences"

// Preference mapped from table <preferences>
type Preference struct {
	ID  int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Key string `gorm:"column:key;type:TEXT" json:"key"`
	Val string `gorm:"column:val;type:TEXT" json:"val"`
}

// TableName Preference's table name
func (*Preference) TableName() string {
	return TableNamePreference
}
