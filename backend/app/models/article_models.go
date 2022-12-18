package models

type Article struct {
	Id           int    `db:"id" json:"id" `
	Created_date string `db:"created_date" json:"created_date"`
	Updated_date string `db:"updated_date" json:"updated_date"`
	Title        string `db:"title" json:"title" `
	Content      string `db:"content" json:"content" `
	Category     string `db:"category" json:"category" `
	Status       string `db:"status" json:"status"`
}
