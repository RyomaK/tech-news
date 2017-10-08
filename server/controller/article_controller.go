package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/RyomaK/tech_news/server/model"
	"github.com/RyomaK/tech_news/server/scrape"
)

//jsonで記事データ返す
func JsonGetArticles(w http.ResponseWriter, r *http.Request) {
	db := model.NewDBConn()
	defer db.Close()
	queryParam := r.URL.Query()
	company := queryParam.Get("company")
	articles := []model.Article{}
	if company != "" {
		articles = model.CompanyArticles(db, company)
	} else {
		articles = model.AllArticles(db)
	}
	a, err := json.Marshal(articles)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(a)
}

func JsonGetCompanies(w http.ResponseWriter, r *http.Request) {
	db := model.NewDBConn()
	defer db.Close()
	companies := []model.Company{}
	companies = model.GetCompanies(db)
	a, err := json.Marshal(companies)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(a)
}

func UpdateDB() {
	db := model.NewDBConn()
	defer db.Close()
	for {
		//二日おきに更新
		model.UpdateArticles(db, scrape.NewArticles())
		time.Sleep(48 * time.Hour)
	}
}
