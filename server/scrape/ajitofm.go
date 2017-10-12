package scrape

import (
	"fmt"

	"time"

	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/RyomaK/tech_news/server/model"
)

func Ajito() []model.Article {
	url := "https://ajito.fm/"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Print("ajito scarapping failed")
	}
	articles := []model.Article{}
	doc.Find(".post-preview").Each(func(_ int, s *goquery.Selection) {
		article := model.Article{Company: "ajito"}
		article.Title = strings.Replace(s.Find(".post-title").Text(), "/n", "", -1)
		t, _ := time.Parse(" 2006 Jan 2", getDate(strings.TrimSpace(s.Find(".post-meta").Text())))
		article.PostedAt = t
		article.URL, _ = s.Find(".post-preview > a").Attr("href")
		article.Text = ""
		articles = append(articles, article)
	})

	return articles
}

func getDate(str string) string {
	s := strings.Split(str, ",")
	return s[2] + s[1]
}

func NewAjito() model.Article {
	url := "https://ajito.fm/"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Print("ajito scarapping failed")
	}
	article := model.Article{Company: "ajito"}
	doc.Find(".post-preview").EachWithBreak(func(_ int, s *goquery.Selection) bool {
		article.Title = strings.Replace(s.Find(".post-title").Text(), "/n", "", -1)
		t, _ := time.Parse(" 2006 Jan 2", getDate(strings.TrimSpace(s.Find(".post-meta").Text())))
		article.PostedAt = t
		article.URL, _ = s.Find(".post-preview > a").Attr("href")
		article.Text = ""
		return false
	})
	return article
}
