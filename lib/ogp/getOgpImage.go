package ogp

import (
	"github.com/CryptocurrencyApp/CoinpocketServer/models"
	"github.com/PuerkitoBio/goquery"
)

func GetOgpImage(article *models.Article) (err error) {
	doc, _ := goquery.NewDocument(article.Url)

	// ogpがない場合の初期値
	// 画像はionicにある
	article.Image = "../../assets/imgs/noImage.png"
	article.SiteTitle = doc.Find("title").Text()
	article.SiteName  = doc.Url.Host

	doc.Find("meta[property='og:image']").Each(func(_ int, s *goquery.Selection) {
		property, _ := s.Attr("content")
		article.Image = property
	})
	doc.Find("meta[property='og:title']").Each(func(_ int, s *goquery.Selection) {
		property, _ := s.Attr("content")
		article.SiteTitle = property
	})
	doc.Find("meta[property='og:site_name']").Each(func(_ int, s *goquery.Selection) {
		property, _ := s.Attr("content")
		article.SiteName = property
	})
	return
}
