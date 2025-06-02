package main

import (
	"runtime"
	"slices"
	"strconv"
	"sync"
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
	MAX_THREADS := runtime.NumCPU()

	//! helpers
	var queuePaths []string
	urlPaths := make(map[string]bool)

	//! do the first cycle
	urlChan := make(chan string)
	u.PrintColor(u.PURPLE, "using one thread for the first loop...\n")
	go func() {
		s.FetchUrl(domain, urlChan)
		defer close(urlChan)
	}()

	for url := range urlChan {

		if url == "" {
			continue
		}

		fixedUrl := u.RepairPath(domain, url)
		if _, ok := urlPaths[fixedUrl]; !ok {
			u.PrintColor(u.WHITE, fixedUrl+"\n")
			urlPaths[fixedUrl] = true
			queuePaths = append(queuePaths, fixedUrl)
		}
	}

	//! loop
	u.PrintColor(u.PURPLE, "RUNNING "+strconv.Itoa(MAX_THREADS)+" THREADS\n")
	var keepContinuing bool = true

	for keepContinuing {
		results := make(chan string, MAX_THREADS)
		var wg sync.WaitGroup
		for idx := range MAX_THREADS {
			if idx > len(queuePaths) {
				return
			}
			wg.Add(1)
			go func() {
				queuePaths = slices.Delete(queuePaths, idx, idx+1)
				s.FetchUrl(queuePaths[idx], results)
				defer wg.Done()
			}()
		}

		go func() {
			wg.Wait()
			defer close(results)
		}()

		for url := range urlChan {

			if url == "" {
				continue
			}

			fixedUrl := u.RepairPath(domain, url)
			if _, ok := urlPaths[fixedUrl]; !ok {
				u.PrintColor(u.WHITE, fixedUrl+"\n")
				urlPaths[fixedUrl] = true
				queuePaths = append(queuePaths, fixedUrl)
			}
		}

		if len(urlPaths) >= LIMIT_URL_REQS {
			keepContinuing = u.PromptConfirm(LIMIT_URL_REQS)
		}
	}

}
