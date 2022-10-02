package model

type SysInfo struct {
	Books      int          `json:"books"`
	Tags       int          `json:"tags"`
	Authors    int          `json:"authors"`
	Publishers int          `json:"publishers"`
	Series     int          `json:"series"`
	Mtime      string       `json:"mtime"`
	Users      int          `json:"users"`
	Active     int          `json:"active"`
	Version    string       `json:"version"`
	Title      string       `json:"title"`
	Friends    []FriendSite `json:"friends"`
	Footer     string       `json:"footer"`
	Allow      SysAllow     `json:"allow"`
}

type FriendSite struct {
	Name string `json:"text"`
	Href string `json:"href"`
}

type SysAllow struct {
	Register bool `json:"register"`
	Download bool `json:"download"`
	Push     bool `json:"push"`
	Read     bool `json:"read"`
}
