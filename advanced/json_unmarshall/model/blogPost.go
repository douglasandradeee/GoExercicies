package model

type BlogPost struct {
	UserIdD int    `json:"userId"`
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
}
