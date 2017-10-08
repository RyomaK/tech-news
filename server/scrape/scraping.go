package scrape

import "github.com/RyomaK/tech_news/server/model"

func NewArticles() []model.Article {
	articles := []model.Article{}
	articles = append(articles, NewAjito())
	articles = append(articles, NewCookPad(2017))
	articles = append(articles, NewCyberAgent())
	articles = append(articles, NewEureka())
	articles = append(articles, NewSmartNews())
	articles = append(articles, NewVoyage(2017))
	articles = append(articles, NewWantedly())
	return articles
}
