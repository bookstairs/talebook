package model

const TableNameIdentifier = "identifiers"

// Identifier mapped from table <identifiers>
type Identifier struct {
	ID   int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Book int32  `gorm:"column:book;type:INTEGER" json:"book"`
	Type string `gorm:"column:type;type:TEXT" json:"type"`
	Val  string `gorm:"column:val;type:TEXT" json:"val"`
}

// TableName Identifier's table name
func (*Identifier) TableName() string {
	return TableNameIdentifier
}
