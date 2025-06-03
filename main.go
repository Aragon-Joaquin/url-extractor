package main

import (
	"runtime"
	"strconv"
	"strings"
	"sync"
	u "url-extractor/utils"
)

const (
	LIMIT_URL_REQS int = 20
)

type GlobalVariables struct {
	UrlPaths   map[string]bool
	QueuePaths []string
	UrlCrawled int
}

func main() {
	domain, err := u.PromptInput()

	if err != nil {
		panic(err)
	}

	//! constants
	MAX_THREADS := runtime.NumCPU()

	globalVar := &GlobalVariables{
		UrlPaths:   make(map[string]bool),
		QueuePaths: []string{},
		UrlCrawled: 0,
	}

	//! do the first cycle
	urlChan := make(chan string)

	u.PrintColor(u.PURPLE, "using one thread for the first loop...\n")

	go func() {
		MakePetition(domain, urlChan, globalVar)
		defer close(urlChan)
	}()

	for url := range urlChan {
		if url == "" {
			continue
		}

		CheckURLState(url, globalVar)

		if strings.HasPrefix(url, domain) {
			globalVar.QueuePaths = append(globalVar.QueuePaths, url)
		}
	}

	/*
	* program second stage. bulk requesting... i don´t know if this term exists but whatever
	 */

	keepContinuing := u.PromptConfirm(1, u.CONFIRM_BULK_REQUEST)
	u.PrintColor(u.PURPLE, "RUNNING "+strconv.Itoa(MAX_THREADS)+" THREADS\n")

	//! loop
	for keepContinuing && len(globalVar.QueuePaths) > 0 {
		results := make(chan string, MAX_THREADS)
		wg := &sync.WaitGroup{}

		for range MAX_THREADS {
			if len(globalVar.QueuePaths) == 0 {
				break
			}
			wg.Add(1)
			currentPath := globalVar.QueuePaths[0]
			globalVar.QueuePaths = globalVar.QueuePaths[1:]

			go func() {
				defer wg.Done()
				MakePetition(currentPath, results, globalVar)
			}()
		}

		go func() {
			wg.Wait()
			close(results)
		}()

		for url := range results {
			if url == "" {
				continue
			}
			CheckURLState(url, globalVar)

			if strings.HasPrefix(url, domain) {
				globalVar.QueuePaths = append(globalVar.QueuePaths, url)
			}
		}

		if globalVar.UrlCrawled >= LIMIT_URL_REQS {
			keepContinuing = u.PromptConfirm(LIMIT_URL_REQS, u.MESSAGE_PER_REQUESTS)
		}
	}

	//! goodbye
	if keepContinuing {
		u.PrintColor(u.GREEN, "Cancelled execution.\n")
	} else {
		u.PrintColor(u.GREEN, "No more urls left to search.\n")
	}

}
