package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Make HTTP request
	res, err := http.Get("https://www.iana.org/help/example-domains")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	// Parse HTML
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find and print article titles
	doc.Find(".help-article").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		fmt.Printf("Article %d: %s\n", i+1, title)
	})
}
