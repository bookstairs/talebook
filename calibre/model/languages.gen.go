package model

const TableNameLanguage = "languages"

// Language mapped from table <languages>
type Language struct {
	ID       int32  `gorm:"column:id;type:INTEGER" json:"id"`
	LangCode string `gorm:"column:lang_code;type:TEXT" json:"lang_code"`
}

// TableName Language's table name
func (*Language) TableName() string {
	return TableNameLanguage
}
