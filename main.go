package main

import (
	"runtime"
	"strconv"
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
	var queuePaths []string
	urlPaths := make(map[string]bool)

	//! do the first cycle
	urlChan := make(chan string)

	u.PrintColor(u.PURPLE, "using one thread for the first loop...\n")

	go func() {
		MakePetition(domain, urlChan)
		defer close(urlChan)
	}()

	for url := range urlChan {
		if url == "" {
			continue
		}

		CheckURLState(domain, url, &urlPaths, &queuePaths)
	}

	u.PrintColor(u.PURPLE, "RUNNING "+strconv.Itoa(MAX_THREADS)+" THREADS\n")

	//! loop

}
