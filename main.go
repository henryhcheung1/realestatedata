package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

type Property struct {
	Title       string
	Description string
	Creator     string
	Level       string
	URL         string
	Language    string
	Commitment  string
	HowToPass   string
	Rating      string
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("zillow.com", "www.zillow.com"),
		colly.AllowURLRevisit(),
		colly.CacheDir("./zillow_cache"),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
	})

	c.OnHTML("html", func(e *colly.HTMLElement) {
		htmlContent := e.Text
		fmt.Println(htmlContent)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit("https://www.zillow.com/irvine-ca/?searchQueryState=%7B%22usersSearchTerm%22%3A%22Los%20Angeles%2C%20CA%22%2C%22mapBounds%22%3A%7B%22north%22%3A33.82412885475799%2C%22east%22%3A-117.67383009448243%2C%22south%22%3A33.54974119236114%2C%22west%22%3A-117.87467390551758%7D%2C%22isMapVisible%22%3Atrue%2C%22filterState%22%3A%7B%22sort%22%3A%7B%22value%22%3A%22days%22%7D%2C%22ah%22%3A%7B%22value%22%3Atrue%7D%2C%22sche%22%3A%7B%22value%22%3Afalse%7D%2C%22schm%22%3A%7B%22value%22%3Afalse%7D%2C%22schh%22%3A%7B%22value%22%3Afalse%7D%2C%22schp%22%3A%7B%22value%22%3Afalse%7D%2C%22schr%22%3A%7B%22value%22%3Afalse%7D%2C%22schc%22%3A%7B%22value%22%3Afalse%7D%2C%22schu%22%3A%7B%22value%22%3Afalse%7D%2C%22mf%22%3A%7B%22value%22%3Afalse%7D%2C%22land%22%3A%7B%22value%22%3Afalse%7D%2C%22apa%22%3A%7B%22value%22%3Afalse%7D%2C%22manu%22%3A%7B%22value%22%3Afalse%7D%7D%2C%22isListVisible%22%3Atrue%2C%22mapZoom%22%3A13%2C%22regionSelection%22%3A%5B%7B%22regionId%22%3A52650%2C%22regionType%22%3A6%7D%5D%2C%22pagination%22%3A%7B%7D%7D")
	if err != nil {
		log.Fatal("Failed to load HTML:", err)
	}
}
