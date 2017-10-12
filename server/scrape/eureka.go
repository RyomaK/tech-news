package scrape

import (
	"fmt"
	"time"

	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/RyomaK/tech_news/server/model"
)

func Eureka() []model.Article {
	doc, err := goquery.NewDocument("https://developers.eure.jp/")
	if err != nil {
		fmt.Print("eureka scarapping failed")
	}
	articles := []model.Article{}
	doc.Find(".articleWrap").Each(func(_ int, s *goquery.Selection) {
		article := model.Article{Company: "eureka"}
		article.Title = strings.Replace(s.Find(".articleTitle > a ").Text(), "/n", "", -1)
		if err != nil {
			fmt.Print("title err")
		}
		s.Find(".articleInfoList >li ").Each(func(i int, se *goquery.Selection) {
			if i == 0 {
				t, _ := time.Parse("2006.01.02", se.Text())
				article.PostedAt = t
			}
		})
		article.URL, _ = s.Find(".articleText > a").Attr("href")
		article.Text = s.Find(".articleText > a").Text()
		articles = append(articles, article)
	})

	return articles
}

func NewEureka() model.Article {
	doc, err := goquery.NewDocument("https://developers.eure.jp/")
	if err != nil {
		fmt.Print("eureka scarapping failed")
	}
	article := model.Article{Company: "eureka"}
	doc.Find(".articleWrap").EachWithBreak(func(_ int, s *goquery.Selection) bool {
		article.Title = strings.Replace(s.Find(".articleTitle > a ").Text(), "/n", "", -1)
		if err != nil {
			fmt.Print("title err")
		}
		s.Find(".articleInfoList >li ").Each(func(i int, se *goquery.Selection) {
			if i == 0 {
				t, _ := time.Parse("2006.01.02", se.Text())
				article.PostedAt = t
			}
		})
		article.URL, _ = s.Find(".articleText > a").Attr("href")
		article.Text = s.Find(".articleText > a").Text()
		return false
	})

	return article
}
