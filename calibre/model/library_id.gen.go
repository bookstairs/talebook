package model

const TableNameLibraryID = "library_id"

// LibraryID mapped from table <library_id>
type LibraryID struct {
	ID   int32  `gorm:"column:id;type:INTEGER" json:"id"`
	UUID string `gorm:"column:uuid;type:TEXT" json:"uuid"`
}

// TableName LibraryID's table name
func (*LibraryID) TableName() string {
	return TableNameLibraryID
}
