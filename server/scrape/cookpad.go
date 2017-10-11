package scrape

import (
	"fmt"
	"strconv"

	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/RyomaK/tech_news/server/model"
)

func CookPad(year int) []model.Article {
	url := "http://techlife.cookpad.com/archive/" + strconv.Itoa(year)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Print("cookpad scarapping failed")
	}
	articles := []model.Article{}
	doc.Find("section").Each(func(_ int, s *goquery.Selection) {
		article := model.Article{Company: "cookpad"}
		article.Title = s.Find(".entry-title-link").Text()
		date, _ := s.Find("time").Attr("datetime")
		t, _ := time.Parse("2006-01-02", date)
		article.PostedAt = t
		article.URL, _ = s.Find(".entry-title-link").Attr("href")
		article.Text = s.Find(".entry-description").Text()
		articles = append(articles, article)
	})

	return articles
}

func NewCookPad(year int) model.Article {
	url := "http://techlife.cookpad.com/archive/" + strconv.Itoa(year)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Print("cookpad scarapping failed")
	}
	article := model.Article{Company: "cookpad"}
	doc.Find("section").EachWithBreak(func(_ int, s *goquery.Selection) bool {
		article.Title = s.Find(".entry-title-link").Text()
		date, _ := s.Find("time").Attr("datetime")
		t, _ := time.Parse("2006-01-02", date)
		article.PostedAt = t
		article.URL, _ = s.Find(".entry-title-link").Attr("href")
		article.Text = s.Find(".entry-description").Text()
		return false
	})

	return article
}
