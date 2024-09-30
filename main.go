package main

import (
	"fmt"
	"net/http"
	"workspace/internal/app"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	app.Init() // init config and logs

	// 1. Core modules: app, asynq
	url := "https://www.포레나당산.com/board/bbs/board.php?bo_table=notice"
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to send GET request: %v", err)
		return
	}

	//_ = res.Body.Close()
	defer res.Body.Close()

	// goquery를 사용하여 HTML 파싱
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		//panic(err)
		fmt.Printf("Failed to send GET request: %v", err)
	}

	// Find the review items
	doc.Find(".tbl_head01.tbl_wrap > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title, _ := s.Html()
		fmt.Printf("Review %d: %s\n", i, title)
		//title := s.Find("td > a").Text()
		//fmt.Printf("Review %d: %s\n", i, title)
	})
}
