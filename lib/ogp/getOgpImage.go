package ogp

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/CryptocurrencyApp/CoinpocketServer/models"
)

func GetOgpImage(article *models.Article) (err error) {
	doc, _ := goquery.NewDocument(article.Url)
	doc.Find("meta[property='og:title']").Each(func(_ int, s *goquery.Selection) {
		property, _ := s.Attr("content")
		article.SiteTitle = property
	})
	doc.Find("meta[property='og:image']").Each(func(_ int, s *goquery.Selection) {
		property, _ := s.Attr("content")
		article.Image = property
	})
	doc.Find("meta[property='og:site_name']").Each(func(_ int, s *goquery.Selection) {
		property, _ := s.Attr("content")
		article.SiteName = property
	})
	return
}
