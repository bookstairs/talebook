package model

const TableNameAnnotation = "annotations"

// Annotation mapped from table <annotations>
type Annotation struct {
	ID             int32   `gorm:"column:id;type:INTEGER" json:"id"`
	Book           int32   `gorm:"column:book;type:INTEGER" json:"book"`
	Format         string  `gorm:"column:format;type:TEXT" json:"format"`
	UserType       string  `gorm:"column:user_type;type:TEXT" json:"user_type"`
	User           string  `gorm:"column:user;type:TEXT" json:"user"`
	Timestamp      float64 `gorm:"column:timestamp;type:REAL" json:"timestamp"`
	AnnotID        string  `gorm:"column:annot_id;type:TEXT" json:"annot_id"`
	AnnotType      string  `gorm:"column:annot_type;type:TEXT" json:"annot_type"`
	AnnotData      string  `gorm:"column:annot_data;type:TEXT" json:"annot_data"`
	SearchableText string  `gorm:"column:searchable_text;type:TEXT" json:"searchable_text"`
}

// TableName Annotation's table name
func (*Annotation) TableName() string {
	return TableNameAnnotation
}
