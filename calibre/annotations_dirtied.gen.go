package calibre

const TableNameAnnotationsDirtied = "annotations_dirtied"

// AnnotationsDirtied mapped from table <annotations_dirtied>
type AnnotationsDirtied struct {
	ID   int32 `gorm:"column:id;type:INTEGER" json:"id"`
	Book int32 `gorm:"column:book;type:INTEGER" json:"book"`
}

// TableName AnnotationsDirtied's table name
func (*AnnotationsDirtied) TableName() string {
	return TableNameAnnotationsDirtied
}
