package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
)

func main() {
	resp, err := soup.Get("https://mathias-kettner.com/cms_check_plugins_list.html")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	fmt.Println("Official Check Plug-Ins listed on")
	fmt.Println("https://mathias-kettner.com/cms_check_plugins_list.html:")
	fmt.Println("")
	checks := doc.FindAllStrict("td", "class", "tt")
	links := doc.Find("tbody").FindAllStrict("a", "class", "quer")
	for i, _ := range checks {
		fmt.Println(checks[i].Text())
		fmt.Println("  " + links[i].Text())
		fmt.Println("  https://mathias-kettner.com/" + links[i].Attrs()["href"])
	}
}
