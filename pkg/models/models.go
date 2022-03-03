package models

type Request struct {
	ID       int    `json:"id,omitempty"`
	UserID   int64  `json:"user_id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Respons  string `json:"respons,omitempty"`
	ResTime  string `json:"resTime,omitempty"`
}
