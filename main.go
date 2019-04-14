package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"gopkg.in/urfave/cli.v1"
	"os"
	"time"
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

	app := cli.NewApp()
	app.Name = "cpf"
	app.Version = "0.1.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Florian Kromer",
			Email: "florian.kromer@mailbox.org",
		},
	}
	app.Copyright = "(c) 2019 Florian Kromer"

	app.Usage = `Check Plugin Find - Get metadata from official Check_MK check plugins.

When executing the app the first time the execution time is in the range of
approx. 30 seconds. The app has the side effect of caching data in
/tmp/.cpf_cache for providing the metadata more quickly after the first execution.
To force getting the data from the website remove this cache file.`

	// main command executed if no -h and -v provided
	app.Action = func(c *cli.Context) error {

		checks := make([]Check, 0, 1500)

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
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
