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

	// check plugin list website:
	// callback triggers on every table line element to get check plugin,
	// gets the check name from the first td element and the title from the second element.
	counter := 0
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
				// TODO: trigger child collector
				counter = counter + 1
			}
		})
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
		fmt.Println("Name:   " + check.Name)
		fmt.Println(" Title: " + check.Title)
	}
}
