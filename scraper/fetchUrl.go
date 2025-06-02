package scraper

import (
	"log"
	"net/http"

	"golang.org/x/net/html"
)

const (
	ANCHOR_TAG     string = "a"
	HREF_ATTRIBUTE string = "href"
)

func FetchUrl(link string, urlChan chan string) {
	//! fetch
	res, err := http.Get(link)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	//! parse to html and find anchor
	z := html.NewTokenizer(res.Body)
	for {
		token := z.Next()

		if token == html.ErrorToken {
			return
		}

		if token == html.StartTagToken {
			element := z.Token()
			if element.Data == ANCHOR_TAG {
				urlChan <- getAttrHref(&element)
				continue
			}
		}
	}
}

func getAttrHref(t *html.Token) string {
	for _, val := range t.Attr {
		if val.Key == HREF_ATTRIBUTE {
			return val.Val
		}
	}
	return ""
}
