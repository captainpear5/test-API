package models

type Post struct {
	Id     int    `json:"id" gorm:"primary_key"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
