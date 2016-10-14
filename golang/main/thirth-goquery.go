package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	doc, err := goquery.NewDocument("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
}

func ExampleJialinWu() {
	doc, err := goquery.NewDocument("https://jialinwu.com/about-me/")
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".entry p").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		// band := s.Find("a").Text()
		// title := s.Find("p").Text()
		// title, err := s.Find("a").Html()
		title, err := s.Html()
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%d p: %+v\n", i, title)

		band, exit := s.Find("a").Attr("href")
		if exit {
			fmt.Printf("band: %+v\n", band)
		}
		// fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
}

func main() {
	// ExampleScrape()
	ExampleJialinWu()
}
