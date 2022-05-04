package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	location string
	title    string
	salary   string
	summary  string
}

var baseURL string = ""

func main() {
	var jobs []extractedJob
	totalPages := getPages()
	for i := 0; i < totalPages; i++ {
		extractedJob := getPage(i)
		jobs = append(jobs, extractedJob...)
	}

	fmt.Println(jobs)

}

func getPage(pagenumber int) []extractedJob {
	var jobs []extractedJob
	pageURL := baseURL + "&start=" + strconv.Itoa(pagenumber*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".tapItem")

	searchCards.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})

	return jobs

}

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find("h2>span").Text())
	location := cleanString(card.Find(".companyLocation").Text())
	salary := cleanString(card.Find(".salary-snippet").Text())
	summary := cleanString(card.Find(".job-snippet").Text())
	return extractedJob{id: id, title: title, location: location, salary: salary, summary: summary}
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination-list").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status", res.StatusCode)
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
