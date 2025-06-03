package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	s "url-extractor/scraper"
	u "url-extractor/utils"
)

func MakePetition(path string, resChan chan<- string) {
	ctx, cancel := context.WithTimeout(context.Background(), s.HTTP_TIMEOUT)
	defer cancel()

	html, err := s.FetchUrl(ctx, path)
	if err != nil {
		u.PrintColor(u.RED, "Error: "+err.Error()+"\n")
		resChan <- ""
		return
	}
	s.ParseHtml(html, resChan)

	select {
	case <-ctx.Done():
		fmt.Println("Context cancelled")
	default:
	}
}

func CheckURLState(domain string, url string, urlPaths *map[string]bool, queuePaths *[]string) {
	fixedUrl := u.RepairPath(domain, url)
	if _, ok := (*urlPaths)[fixedUrl]; !ok {
		u.PrintColor(u.WHITE, strconv.Itoa(len(*urlPaths)+1)+": "+fixedUrl+"\n")
		(*urlPaths)[fixedUrl] = true

		if strings.HasPrefix(fixedUrl, domain) {
			*queuePaths = append(*queuePaths, fixedUrl)
		}
	}
}
