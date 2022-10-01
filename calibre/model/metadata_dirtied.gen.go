package model

const TableNameMetadataDirtied = "metadata_dirtied"

// MetadataDirtied mapped from table <metadata_dirtied>
type MetadataDirtied struct {
	ID   int32 `gorm:"column:id;type:INTEGER" json:"id"`
	Book int32 `gorm:"column:book;type:INTEGER" json:"book"`
}

// TableName MetadataDirtied's table name
func (*MetadataDirtied) TableName() string {
	return TableNameMetadataDirtied
}
