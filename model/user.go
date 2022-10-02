package model

type User struct {
	Avatar      string         `json:"avatar"`
	IsLogin     bool           `json:"is_login"`
	IsAdmin     bool           `json:"is_admin"`
	Nickname    string         `json:"nickname"`
	Email       string         `json:"email"`
	KindleEmail string         `json:"kindle_email"`
	Extra       map[string]any `json:"extra"`
}
