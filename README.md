# tech_news
___techblog__のスマートニュース版みたいの作る___
## language 
### server 
Golang
### front 
swift

## ToDo
- [x] tech blog をまとめる
- [x] tech blog スクレイピング
- [x] 取得した記事をDB格納
- [x] API設計
- [ ] dbconfig.yml

## API
/api/articles/
?q=company

```json
[
  "id":1,
  "title":"aa",
  "poseted_at":,
  "url",
  "text",
  "company"
]
```

### techblog
- [COOK PAD](http://techlife.cookpad.com/)
- [eureka](https://developers.eure.jp/)
- [VOYAGE](http://techlog.voyagegroup.com/)
- [ajito in vouyage](https://ajito.fm/)
- [smart news](https://developer.smartnews.com/blog/)
- [wantedly](https://www.wantedly.com/feed/s/wantedly_engineers)
- [サイバーエージェント](https://developers.cyberagent.co.jp/blog/)

### 起動
```bash
//db起動
systemctl start mariadb.service
//iptables　設定
sudo iptables -A INPUT -p tcp --dport (port_mum) -j ACCEPT
sudo systemctl restart iptables
//起動
nohup ./tech &
