package main

import (
	"flag"
	"log"
	"net"
	"net/url"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

var defaultSite string = "https://www.scrapethissite.com/pages/simple/"

func main() {
	// setup CLI flags
	site := flag.String("site", "", "site to scrape")
	flag.Parse()

	// check if URL passed
	if *site == "" {
		log.Fatalln("Usage: ", os.Args[0], "-site <SITE>")
	}

	// Ensure input formatted correctly
	domain := *site
	if !strings.HasPrefix(*site, "https://") {
		domain = "https://" + *site
	}

	u, err := url.Parse(domain)
	if err != nil {
		log.Fatalln(err)
	}
	// check if domain resolves
	_, err = net.LookupIP(u.Host)
	if err != nil {
		log.Fatalln("Invalid URL: ", u.Host)
	}

	// Colly configuration
	c := colly.NewCollector(
		colly.AllowedDomains(u.Host),
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

	c.Visit(domain)
}
