package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Create a new collector
	c := colly.NewCollector(
		// Visit only domains: example.com, www.example.com
		colly.AllowedDomains("example.com"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println("Link found:", link)
		// Visit link found on page
		e.Request.Visit(link)
	})

	// On every page scraped, check for title and print it
	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println("Page Title:", e.Text)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://example.com
	c.Visit("https://example.com")
}
