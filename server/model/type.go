package model

import "time"

type Article struct {
	ID       uint      `db:"id" json:"id"`
	Title    string    `db:"title" json:"title"`
	PostedAt time.Time `db:"posted_at" json:"posted_at"`
	URL      string    `db:"url" json:"url"`
	Text     string    `db:"text" json:"text"`
	Company  string    `db:"company" json:"company"`
}

type Company struct {
	ID      uint   `db:"id" json:"id"`
	Company string `db:"company" json:"company"`
}
