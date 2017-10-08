package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewDBConn() *sql.DB {
	//db time 型合わせる
	db, err := sql.Open("mysql", "root:@/tech_news?parseTime=true")
	if err != nil {
		fmt.Errorf("db err : %v", err)
	}
	return db
}

func AllArticles(db *sql.DB) []Article {
	rows, _ := db.Query("SELECT * from articles")
	articles, err := ScanArticles(rows)
	if err != nil {
		fmt.Errorf("err article all : %v", err)
	}
	return articles
}

func CompanyArticles(db *sql.DB, company string) []Article {
	rows, _ := db.Query(`SELECT * from articles where company = ? `, company)
	articles, err := ScanArticles(rows)
	if err != nil {
		fmt.Errorf("err companey articles  : %v", err)
	}
	return articles
}

func isExsit(db *sql.DB, article Article) bool {
	row := db.QueryRow(`SELECT * from articles where company = ? AND title = ?`, article.Company, article.Title)
	match, _ := ScanArticle(row)
	if match == (Article{}) {
		return false
	}
	return true
}

func UpdateArticle(db *sql.DB, article Article) {
	if !isExsit(db, article) {
		stmt, err := db.Prepare("INSERT articles SET title=?,posted_at=?,url=?,text=?,company=?")
		if err != nil {
			panic(err)
		}
		_, err = stmt.Exec(article.Title, article.PostedAt, article.URL, article.Text, article.Company)
		if err != nil {
			fmt.Errorf("%v:%v", err, article.Title)
		}
		stmt.Close()
	}
}

func UpdateArticles(db *sql.DB, articles []Article) {
	for _, article := range articles {
		if !isExsit(db, article) {
			stmt, err := db.Prepare("INSERT articles SET title=?,posted_at=?,url=?,text=?,company=?")
			if err != nil {
				panic(err)
			}
			_, err = stmt.Exec(article.Title, article.PostedAt, article.URL, article.Text, article.Company)
			if err != nil {
				fmt.Errorf("%v:%v", err, article.Title)
			}
			stmt.Close()
		}
	}
}

func GetCompanies(db *sql.DB) []Company {
	rows, _ := db.Query(`SELECT * FROM companies`)
	companies, err := ScanCompanies(rows)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	return companies

}
