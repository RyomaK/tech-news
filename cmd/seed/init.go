package main

import (
	"database/sql"
	"fmt"

	"github.com/RyomaK/tech_news/server/model"
	"github.com/RyomaK/tech_news/server/scrape"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := model.NewDBConn()
	defer db.Close()
	//記事
	insertData(db, scrape.Ajito())
	insertData(db, scrape.CookPad(2017))
	insertData(db, scrape.CookPad(2016))
	insertData(db, scrape.CyberAgent())
	insertData(db, scrape.Eureka())
	insertData(db, scrape.SmartNews())
	insertData(db, scrape.Voyage(2017))
	insertData(db, scrape.Voyage(2016))
	insertData(db, scrape.Wantedly())

	//会社
	insertCompany(db, model.Company{Company: "ajito"})
	insertCompany(db, model.Company{Company: "cookpad"})
	insertCompany(db, model.Company{Company: "cyberagent"})
	insertCompany(db, model.Company{Company: "eureka"})
	insertCompany(db, model.Company{Company: "smartnews"})
	insertCompany(db, model.Company{Company: "voyage"})
	insertCompany(db, model.Company{Company: "wantedly"})

}

func insertData(d *sql.DB, articles []model.Article) {
	for _, article := range articles {
		stmt, err := d.Prepare("INSERT articles SET title=?,posted_at=?,url=?,text=?,company=?")
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

func insertCompany(d *sql.DB, company model.Company) {
	stmt, err := d.Prepare("INSERT companies SET company=?")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(company.Company)
	if err != nil {
		fmt.Errorf("%v:%v", err, company)
	}
	stmt.Close()
}
