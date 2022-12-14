package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
}

// scrape indeed by a term
func Scrape(term string) {
	var baseURL string = "https://www.jobkorea.co.kr/Search/?stext=" + term + "&tabType=recruit"
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages(baseURL)
	fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ {
		go getPage(i, baseURL, c)

	}
	for i := 0; i < totalPages; i++ {
		extractJobs := <-c
		jobs = append(jobs, extractJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted= ", len(jobs))
}
func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "Title", "Location"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://www.jobkorea.co.kr/Recruit/GI_Read/" + job.id, job.title, job.location}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

// https://www.jobkorea.co.kr/Search/?stext=python&tabType=recruit&Page_No=1
func getPage(page int, url string, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := url + "&Page_No=" + strconv.Itoa(page)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".list-post")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)

	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-gno")
	title := CleanString(card.Find(".title").Text())
	location := CleanString(card.Find(".loc.long").Text())
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location}

}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".tplPagination.newVer.wide").Each(func(i int, s *goquery.Selection) {
		// fmt.Println(s.Html())
		// fmt.Println(s.Find("a").Length())
		pages = s.Find("a").Length()
		fmt.Println("page count done")
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
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}
