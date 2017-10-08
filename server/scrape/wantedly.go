package scrape

import (
	"fmt"

	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/RyomaK/tech_news/server/model"
)

func Wantedly() []model.Article {
	url := "https://www.wantedly.com/feed/s/wantedly_engineers"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Print("wantedly scarapping failed")
	}
	articles := []model.Article{}
	doc.Find(".post-space-item").Each(func(_ int, s *goquery.Selection) {
		article := model.Article{Company: "wantedly"}
		article.Title = s.Find(".post-title").Text()
		article.PostedAt = time.Now()
		comURL, _ := s.Find(".post-content > a ").Attr("href")
		article.URL = "https://www.wantedly.com" + comURL
		article.Text = s.Find(".post-description").Text()
		articles = append(articles, article)
	})

	return articles
}

func NewWantedly() model.Article {
	url := "https://www.wantedly.com/feed/s/wantedly_engineers"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Print("wantedly scarapping failed")
	}
	article := model.Article{Company: "wantedly"}
	doc.Find(".post-space-item").Each(func(_ int, s *goquery.Selection) {
		article.Title = s.Find(".post-title").Text()
		article.PostedAt = time.Now()
		comURL, _ := s.Find(".post-content > a ").Attr("href")
		article.URL = "https://www.wantedly.com" + comURL
		article.Text = s.Find(".post-description").Text()
	})

	return article
}
