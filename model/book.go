package model

type Book struct {
	ID            int64    `json:"id"`
	Title         string   `json:"title"`
	Rating        int      `json:"rating"`
	Timestamp     string   `json:"timestamp"`
	Pubdate       string   `json:"pubdate"`
	Author        string   `json:"author"`
	Authors       []string `json:"authors"`
	AuthorSort    string   `json:"author_sort"`
	Tag           string   `json:"tag"`
	Tags          []string `json:"tags"`
	Publisher     string   `json:"publisher"`
	Comments      string   `json:"comments"`
	Series        string   `json:"series"`
	Language      string   `json:"language"`
	Isbn          string   `json:"isbn"`
	Img           string   `json:"img"`
	Thumb         string   `json:"thumb"`
	Collector     string   `json:"collector"`
	CountVisit    int      `json:"count_visit"`
	CountDownload int      `json:"count_download"`
	AuthorURL     string   `json:"author_url"`
	PublisherURL  string   `json:"publisher_url"`
}
