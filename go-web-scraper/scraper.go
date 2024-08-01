package main

import (
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("ariyonaty.com", "www.ariyonaty.com"),
	)

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		h.Request.Visit(h.Attr("href"))
	})

	// called before an HTTP request is triggered
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting: ", r.URL)
	})

	// triggered when scraper hits error
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error: ", err)
	})

	// triggered when server responds
	c.OnResponse(func(r *colly.Response) {
		log.Println("Page visited: ", r.Request.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		log.Println("Scraped: ", r.Request.URL)
	})

	domains := []string{"https://ariyonaty.com/"}
	for _, domain := range domains {
		log.Println("Scraping: ", domain)
		c.Visit(domain)
	}
}
