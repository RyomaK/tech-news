package scrape

import (
	"fmt"

	"time"

	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/RyomaK/tech_news/server/model"
)

func CyberAgent() []model.Article {
	url := "https://developers.cyberagent.co.jp/blog/archives/category/engineer/"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Print("cyberagent scarapping failed")
	}
	articles := []model.Article{}
	doc.Find("div[class='col-md-12 col-xs-6']  ").Each(func(_ int, s *goquery.Selection) {
		article := model.Article{Company: "cyber agent"}
		article.Title = strings.Replace(s.Find(".card-caeng__title > a ").Text(), "/n", "", -1)
		t, _ := time.Parse("2006/01/02", s.Find("div[class='card-caeng__meta hidden-sm-down'] > time").Text())
		article.PostedAt = t
		article.URL, _ = s.Find(".card-caeng__title > a").Attr("href")
		article.Text = s.Find(".card-caeng__text").Text()
		articles = append(articles, article)
	})

	return articles
}

func NewCyberAgent() model.Article {
	url := "https://developers.cyberagent.co.jp/blog/archives/category/engineer/"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Print("cyberagent scarapping failed")
	}
	article := model.Article{Company: "cyber agent"}
	doc.Find("div[class='col-md-12 col-xs-6']  ").EachWithBreak(func(_ int, s *goquery.Selection) bool {
		article.Title = strings.Replace(s.Find(".card-caeng__title > a ").Text(), "/n", "", -1)
		t, _ := time.Parse("2006/01/02", s.Find("div[class='card-caeng__meta hidden-sm-down'] > time").Text())
		article.PostedAt = t
		article.URL, _ = s.Find(".card-caeng__title > a").Attr("href")
		article.Text = s.Find(".card-caeng__text").Text()
		return false
	})

	return article
}
