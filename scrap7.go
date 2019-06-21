package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	urls := "https://eiga.com/movie/88330/review/"
	doc, err := goquery.NewDocument(urls)
	if err != nil {
		fmt.Print(err)
	}
	comments := doc.Find("p").Text()

	file, err := os.OpenFile("comment_scrap.csv", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.UseCRLF = true

	writer.Write([]string{comments})
	writer.Flush()

}
