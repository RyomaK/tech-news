DBCONF := dbconfig.yml

run: ## Run app
	go run ./cmd/main/tech.go

db/init: 
	go run ./cmd/seed/init.go

install:
	go get github.com/go-sql-driver/mysql
	go get github.com/PuerkitoBio/goquery
