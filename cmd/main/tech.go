package main

import (
	"net/http"

	"log"

	"github.com/RyomaK/tech_news/server/controller"
)

func main() {
	// ログ出力
	log.Printf("Start Go HTTP Server")

	//ネットつながナイトエラー出る
	go controller.UpdateDB()

	http.HandleFunc("/api/articles/", controller.JsonGetArticles)

	http.HandleFunc("/api/companies/", controller.JsonGetCompanies)

	// 8080ポートで起動
	http.ListenAndServe(":8080", nil)
}
