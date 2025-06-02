package main

import (
	"strconv"
	s "url-extractor/scraper"
	u "url-extractor/utils"
)

const (
	LIMIT_URL_REQS int = 50
)

func main() {
	domain, err := u.PromptInput()

	if err != nil {
		panic(err)
	}

	//! constants
	MAX_THREADS := 1
	//sURLPaths := make(map[string]bool)
	//QueuePaths := make([]string)

	//! channels
	urlChan := make(chan string, MAX_THREADS)

	u.PrintColor(u.PURPLE, "RUNNING "+strconv.Itoa(MAX_THREADS)+" THREADS\n")

	for range MAX_THREADS {
		go func() {
			s.FetchUrl(domain, urlChan)
		}()
	}

	for range urlChan {
		url := <-urlChan
		u.PrintColor(u.RED, url+"\n")
		u.PrintColor(u.BLUE, u.RepairPath(domain, url)+"\n")
		// fixedUrl := u.RepairPath(domain, url)
		// if _, ok := URLPaths[fixedUrl]; !ok {
		// 	u.PrintColor(u.WHITE, fixedUrl+"\n")
		// 	URLPaths[fixedUrl] = true
		// }
	}

}
