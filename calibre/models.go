package calibre

import "time"

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

// AnnotationsDirtied mapped from table <annotations_dirtied>
type AnnotationsDirtied struct {
	ID   int32 `gorm:"column:id;type:INTEGER" json:"id"`
	Book int32 `gorm:"column:book;type:INTEGER" json:"book"`
}

// Author mapped from table <authors>
type Author struct {
	ID   int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Name string `gorm:"column:name;type:TEXT" json:"name"`
	Sort string `gorm:"column:sort;type:TEXT" json:"sort"`
	Link string `gorm:"column:link;type:TEXT" json:"link"`
}

// Book mapped from table <books>
type Book struct {
	ID           int32     `gorm:"column:id;type:INTEGER" json:"id"`
	Title        string    `gorm:"column:title;type:TEXT" json:"title"`
	Sort         string    `gorm:"column:sort;type:TEXT" json:"sort"`
	Timestamp    time.Time `gorm:"column:timestamp;type:TIMESTAMP" json:"timestamp"`
	Pubdate      time.Time `gorm:"column:pubdate;type:TIMESTAMP" json:"pubdate"`
	SeriesIndex  float64   `gorm:"column:series_index;type:REAL" json:"series_index"`
	AuthorSort   string    `gorm:"column:author_sort;type:TEXT" json:"author_sort"`
	Isbn         string    `gorm:"column:isbn;type:TEXT" json:"isbn"`
	Lccn         string    `gorm:"column:lccn;type:TEXT" json:"lccn"`
	Path         string    `gorm:"column:path;type:TEXT" json:"path"`
	Flags        int32     `gorm:"column:flags;type:INTEGER" json:"flags"`
	UUID         string    `gorm:"column:uuid;type:TEXT" json:"uuid"`
	HasCover     string    `gorm:"column:has_cover;type:BOOL" json:"has_cover"`
	LastModified time.Time `gorm:"column:last_modified;type:TIMESTAMP" json:"last_modified"`
}

// BooksAuthorsLink mapped from table <books_authors_link>
type BooksAuthorsLink struct {
	ID     int32 `gorm:"column:id;type:INTEGER" json:"id"`
	Book   int32 `gorm:"column:book;type:INTEGER" json:"book"`
	Author int32 `gorm:"column:author;type:INTEGER" json:"author"`
}

// BooksLanguagesLink mapped from table <books_languages_link>
type BooksLanguagesLink struct {
	ID        int32 `gorm:"column:id;type:INTEGER" json:"id"`
	Book      int32 `gorm:"column:book;type:INTEGER" json:"book"`
	LangCode  int32 `gorm:"column:lang_code;type:INTEGER" json:"lang_code"`
	ItemOrder int32 `gorm:"column:item_order;type:INTEGER" json:"item_order"`
}

// BooksPluginDatum mapped from table <books_plugin_data>
type BooksPluginDatum struct {
	ID   int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Book int32  `gorm:"column:book;type:INTEGER" json:"book"`
	Name string `gorm:"column:name;type:TEXT" json:"name"`
	Val  string `gorm:"column:val;type:TEXT" json:"val"`
}

// BooksPublishersLink mapped from table <books_publishers_link>
type BooksPublishersLink struct {
	ID        int32 `gorm:"column:id;type:INTEGER" json:"id"`
	Book      int32 `gorm:"column:book;type:INTEGER" json:"book"`
	Publisher int32 `gorm:"column:publisher;type:INTEGER" json:"publisher"`
}

// BooksRatingsLink mapped from table <books_ratings_link>
type BooksRatingsLink struct {
	ID     int32 `gorm:"column:id;type:INTEGER" json:"id"`
	Book   int32 `gorm:"column:book;type:INTEGER" json:"book"`
	Rating int32 `gorm:"column:rating;type:INTEGER" json:"rating"`
}

// BooksSeriesLink mapped from table <books_series_link>
type BooksSeriesLink struct {
	ID     int32 `gorm:"column:id;type:INTEGER" json:"id"`
	Book   int32 `gorm:"column:book;type:INTEGER" json:"book"`
	Series int32 `gorm:"column:series;type:INTEGER" json:"series"`
}

// BooksTagsLink mapped from table <books_tags_link>
type BooksTagsLink struct {
	ID   int32 `gorm:"column:id;type:INTEGER" json:"id"`
	Book int32 `gorm:"column:book;type:INTEGER" json:"book"`
	Tag  int32 `gorm:"column:tag;type:INTEGER" json:"tag"`
}

// Comment mapped from table <comments>
type Comment struct {
	ID   int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Book int32  `gorm:"column:book;type:INTEGER" json:"book"`
	Text string `gorm:"column:text;type:TEXT" json:"text"`
}

// ConversionOption mapped from table <conversion_options>
type ConversionOption struct {
	ID     int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Format string `gorm:"column:format;type:TEXT" json:"format"`
	Book   int32  `gorm:"column:book;type:INTEGER" json:"book"`
	Data   []byte `gorm:"column:data;type:BLOB" json:"data"`
}

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

// Datum mapped from table <data>
type Datum struct {
	ID               int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Book             int32  `gorm:"column:book;type:INTEGER" json:"book"`
	Format           string `gorm:"column:format;type:TEXT" json:"format"`
	UncompressedSize int32  `gorm:"column:uncompressed_size;type:INTEGER" json:"uncompressed_size"`
	Name             string `gorm:"column:name;type:TEXT" json:"name"`
}

// Feed mapped from table <feeds>
type Feed struct {
	ID     int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Title  string `gorm:"column:title;type:TEXT" json:"title"`
	Script string `gorm:"column:script;type:TEXT" json:"script"`
}

// Identifier mapped from table <identifiers>
type Identifier struct {
	ID   int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Book int32  `gorm:"column:book;type:INTEGER" json:"book"`
	Type string `gorm:"column:type;type:TEXT" json:"type"`
	Val  string `gorm:"column:val;type:TEXT" json:"val"`
}

// Language mapped from table <languages>
type Language struct {
	ID       int32  `gorm:"column:id;type:INTEGER" json:"id"`
	LangCode string `gorm:"column:lang_code;type:TEXT" json:"lang_code"`
}

// LastReadPosition mapped from table <last_read_positions>
type LastReadPosition struct {
	ID      int32   `gorm:"column:id;type:INTEGER" json:"id"`
	Book    int32   `gorm:"column:book;type:INTEGER" json:"book"`
	Format  string  `gorm:"column:format;type:TEXT" json:"format"`
	User    string  `gorm:"column:user;type:TEXT" json:"user"`
	Device  string  `gorm:"column:device;type:TEXT" json:"device"`
	Cfi     string  `gorm:"column:cfi;type:TEXT" json:"cfi"`
	Epoch   float64 `gorm:"column:epoch;type:REAL" json:"epoch"`
	PosFrac float64 `gorm:"column:pos_frac;type:REAL" json:"pos_frac"`
}

// LibraryID mapped from table <library_id>
type LibraryID struct {
	ID   int32  `gorm:"column:id;type:INTEGER" json:"id"`
	UUID string `gorm:"column:uuid;type:TEXT" json:"uuid"`
}

// MetadataDirtied mapped from table <metadata_dirtied>
type MetadataDirtied struct {
	ID   int32 `gorm:"column:id;type:INTEGER" json:"id"`
	Book int32 `gorm:"column:book;type:INTEGER" json:"book"`
}

// Preference mapped from table <preferences>
type Preference struct {
	ID  int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Key string `gorm:"column:key;type:TEXT" json:"key"`
	Val string `gorm:"column:val;type:TEXT" json:"val"`
}

// Publisher mapped from table <publishers>
type Publisher struct {
	ID   int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Name string `gorm:"column:name;type:TEXT" json:"name"`
	Sort string `gorm:"column:sort;type:TEXT" json:"sort"`
}

// Rating mapped from table <ratings>
type Rating struct {
	ID     int32 `gorm:"column:id;type:INTEGER" json:"id"`
	Rating int32 `gorm:"column:rating;type:INTEGER" json:"rating"`
}

// Series mapped from table <series>
type Series struct {
	ID   int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Name string `gorm:"column:name;type:TEXT" json:"name"`
	Sort string `gorm:"column:sort;type:TEXT" json:"sort"`
}

// SqliteSequence mapped from table <sqlite_sequence>
type SqliteSequence struct {
	Name string `gorm:"column:name;type:" json:"name"`
	Seq  string `gorm:"column:seq;type:" json:"seq"`
}

// Tag mapped from table <tags>
type Tag struct {
	ID   int32  `gorm:"column:id;type:INTEGER" json:"id"`
	Name string `gorm:"column:name;type:TEXT" json:"name"`
}
