package main

import (
	"context"
	"fmt"
	"strconv"

	s "url-extractor/scraper"
	u "url-extractor/utils"
)

func MakePetition(path string, resChan chan<- string, globalVar *GlobalVariables) {
	ctx, cancel := context.WithTimeout(context.Background(), s.HTTP_TIMEOUT)
	defer cancel()

	html, err := s.FetchUrl(ctx, path)
	globalVar.UrlCrawled += 1

	if err != nil {
		u.PrintColor(u.RED, "Error: "+err.Error()+"\n")
		resChan <- ""
		return
	}
	s.ParseHtml(html, resChan, path)

	select {
	case <-ctx.Done():
		fmt.Println("Context cancelled")
	default:
	}
}

func CheckURLState(fixedUrl string, globalVar *GlobalVariables) {
	if _, ok := globalVar.UrlPaths[fixedUrl]; !ok {
		u.PrintColor(u.WHITE, strconv.Itoa(len(globalVar.UrlPaths)+1)+": "+fixedUrl+"\n")
		globalVar.UrlPaths[fixedUrl] = true
	}
}
