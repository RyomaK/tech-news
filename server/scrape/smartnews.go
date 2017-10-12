package scrape

import (
	"fmt"

	"time"

	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/RyomaK/tech_news/server/model"
)

func SmartNews() []model.Article {
	url := "https://developer.smartnews.com/blog/"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Print("smart news scarapping failed")
	}
	articles := []model.Article{}
	doc.Find(".li").Each(func(_ int, s *goquery.Selection) {
		article := model.Article{Company: "smartnews"}
		article.Title = strings.Replace(s.Find(".title > a").Text(), "/n", "", -1)
		article.PostedAt = getDateForSmart(s.Find("time").Text())
		article.URL, _ = s.Find(".title > a").Attr("href")
		article.Text = s.Find(".summary").Text()
		articles = append(articles, article)
	})

	return articles
}

func getDateForSmart(str string) time.Time {
	s := strings.Split(str, ",")
	t, _ := time.Parse("2006 Jan 2", strings.TrimSpace(s[1])+" "+s[0])
	return t
}

func NewSmartNews() model.Article {
	url := "https://developer.smartnews.com/blog/"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Print("smart news scarapping failed")
	}
	article := model.Article{Company: "smartnews"}
	doc.Find(".li").EachWithBreak(func(_ int, s *goquery.Selection) bool {
		article.Title = strings.Replace(s.Find(".title > a").Text(), "/n", "", -1)
		article.PostedAt = getDateForSmart(s.Find("time").Text())
		article.URL, _ = s.Find(".title > a").Attr("href")
		article.Text = s.Find(".summary").Text()
		return false
	})

	return article
}
