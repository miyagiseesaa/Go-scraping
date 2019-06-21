package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("https://chomanga.org/archives/126794.html")
	if err != nil {
		fmt.Print(err)
	}
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}
