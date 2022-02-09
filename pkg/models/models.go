package models

type Request struct {
	ID       int    `json:"id"`
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	Respons  string `json:"respons"`
}
