package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

type Check struct {
	Name        string
	Title       string
	URL         string
	Description string
	Item        string
	Discovery   string
}

func main() {
	checks := make([]Check, 0, 3000) // around 1500 Checks right now

	baseUrl := "mathias-kettner.com"
	url := "https://mathias-kettner.com/cms_check_plugins_list.html"

	collector := colly.NewCollector(
		colly.AllowedDomains(baseUrl),
		colly.CacheDir("/tmp/.cpf_cache"), // comment out during debugging
	)
	detailCollector := collector.Clone()

	counter := 0 // ugly but easiest way to implement the scraper

	// Callback triggers on every table line element on the check plugin overview list.
	// Check plugin specific sites don't contain the check name which
	// requires to use this parent collector in addition to the childdetailCollector.
	collector.OnHTML("tr", func(e *colly.HTMLElement) {
		e.ForEach("td", func(i int, el *colly.HTMLElement) {
			switch i {
			case 0:
				check := Check{
					Name: el.Text,
				}
				checks = append(checks, check)
			case 1:
				checks[counter].Title = el.Text
				el.ForEach("a[href]", func(_ int, ele *colly.HTMLElement) {
					checks[counter].URL = ele.Request.AbsoluteURL(ele.Attr("href"))
					// trigger child collector
					detailCollector.Visit(checks[counter].URL)
				})
				counter = counter + 1
			}
		})
	})

	// Callback triggers on every div element which contains check plugin metadata
	// on check plugin specific sites.
	detailCollector.OnHTML("div[class=\"nowiki\"]", func(e *colly.HTMLElement) {
		// TODO: filter to relevant info and assign correspondingly
		checks[counter].Description = e.Text
	})

	// debug callbacks
	collector.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error during scraping:", err)
	})

	// for future verbose output
	// collector.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Try to visit:", r.URL.String())
	// })

	// for future verbose output
	// collector.OnResponse(func(r *colly.Response) {
	// 	fmt.Println("Visited:", r.Request.URL)
	// })

	// start scraping
	collector.Visit(url)

	for _, check := range checks {
		fmt.Println("Name:         " + check.Name)
		fmt.Println(" Title:       " + check.Title)
		fmt.Println(" URL:         " + check.URL)
		fmt.Println(" Description: " + check.Description)
	}
}
