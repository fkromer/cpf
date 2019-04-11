package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
)

type Check struct {
	Name  string
	Title string
	URL   string
}

func main() {
	checks := make([]Check, 0, 3000) // around 1500 Checks right now
	resp, err := soup.Get("https://mathias-kettner.com/cms_check_plugins_list.html")
	if err != nil {
		os.Exit(1) // bad cases: no internet connection, etc.
	}
	doc := soup.HTMLParse(resp)
	fmt.Println("Official Check Plug-Ins listed on")
	fmt.Println("https://mathias-kettner.com/cms_check_plugins_list.html:")
	fmt.Println("")
	checksDoc := doc.FindAllStrict("td", "class", "tt")
	links := doc.Find("tbody").FindAllStrict("a", "class", "quer")
	for i, _ := range checksDoc {
		name := checksDoc[i].Text()
		title := links[i].Text()
		url := "https://mathias-kettner.com/" + links[i].Attrs()["href"]
		check := Check{
			Name:  name,
			Title: title,
			URL:   url,
		}
		checks = append(checks, check)
	}
	for _, check := range checks {
		fmt.Println(check.Name)
		fmt.Println("  " + check.Title)
		fmt.Println("  " + check.URL)
	}
}
