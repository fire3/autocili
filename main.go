package main

import (
    "net/url"
	"fmt"
	"log"
    "os"
    "strings"
	"github.com/PuerkitoBio/goquery"
)

var CILI001 string = "http://cili001.com/"
var REALURL string

func GetBaseUrl() string {

    var newurl string
    doc, err := goquery.NewDocument(CILI001)

    if err != nil {
        log.Fatal(err)
    }

    doc.Find("meta[http-equiv=\"refresh\"]").Each(func(i int, s *goquery.Selection) {
        value, _ := s.Attr("content")
        newurl = strings.Split(value,"=")[1]
    })
    return newurl
}

func GetSearchUrl(title string) string {
    u,err := url.Parse(GetBaseUrl()+"index")
    if err != nil {
        log.Fatal(err)
    }
    q := u.Query()
    q.Add("topic_title3",title)
    u.RawQuery= q.Encode()
    fmt.Println(u)
    return u.String()
}


func ExampleScrape() {

    u := GetSearchUrl("指定幸存者 S02")
	doc, err := goquery.NewDocument(u)
	if err != nil {
		log.Fatal(err)
	}
    // Get the total pages
    pages := doc.Find("div .pages a").Length()
    if pages >= 1 {
        pages = pages - 1
    }
    fmt.Printf("pages: %d\n",pages)
	// Find the review items
	doc.Find(".list-item dd").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		title := s.Find(".b a").Text()
		fmt.Printf("item %d: %s \n", i, title)
		//value, _ := s.Attr("magnet")
		//fmt.Printf("item %d: %s - %s\n", i, title, value)
	})
}

func main() {
    REALURL = GetBaseUrl()
	ExampleScrape()
}
