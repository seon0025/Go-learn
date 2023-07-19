package main

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
	date     string
	company  string
}

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=python"

func main() {
	var jobs []extractedJob
	totalPages := getPages()
	for i := 0; i < totalPages; i++ {
		extractedJob := getPage(i)
		jobs = append(jobs, extractedJob...)
	}
	writeJobs(jobs)
	fmt.Println("Done, extracted : ", len(jobs))
}

func getPage(page int) []extractedJob {
	var jobs []extractedJob
	pageURL := baseURL + "&recruitPage=" + strconv.Itoa(page) + "&recruitSort=relation&recruitPageCount=40&inner_com_type=&company_cd=0%2C1%2C2%2C3%2C4%2C5%2C6%2C7%2C9%2C10&show_applied=&quick_apply=&except_read=&ai_head_hunting=&mainSearch=n"
	fmt.Println("requesting: ", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	searchCards := doc.Find(".item_recruit")
	searchCards.Each(func(i int, s *goquery.Selection) {
		job := extracJob(s)
		jobs = append(jobs, job)
	})
	return jobs
}

func extracJob(s *goquery.Selection) extractedJob {
	id, _ := s.Attr("value")
	title := cleanString(s.Find(".job_tit>a").Text())
	location := cleanString(s.Find(".job_condition>span>a").Text())
	company := cleanString(s.Find(".corp_name>a").Text())
	date := cleanString(s.Find(".job_date>span").Text())
	return extractedJob{id: id, title: title, location: location, company: company, date: date}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w:= csv.NewWriter(file)
	defer w.Flush()

	headers := []string {"LINK", "TITLE", "LOCATION", "DATE", "COMPANY"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs{
		jobSlice := []string{"https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx=" + job.id, job.title, job.location, job.date, job.company}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
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
		log.Fatalln("Request failed with status : ", res.StatusCode)
	}
}
