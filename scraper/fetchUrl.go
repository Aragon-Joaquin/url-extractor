package scraper

import (
	"net/http"
	"time"

	u "url-extractor/utils"

	"golang.org/x/net/html"
)

const (
	ANCHOR_TAG     string        = "a"
	HREF_ATTRIBUTE string        = "href"
	HTTP_TIMEOUT   time.Duration = 5 * time.Second
)

func FetchUrl(link string, urlChan chan<- string) {
	//! fetch
	c := &http.Client{
		Timeout: HTTP_TIMEOUT,
	}

	res, err := c.Get(link)

	if err != nil {
		u.PrintColor(u.RED, "Request longed too much: "+HTTP_TIMEOUT.String()+"\n")
		urlChan <- ""
		return
	}

	defer res.Body.Close()

	//! parse to html and find anchor
	z := html.NewTokenizer(res.Body)

	for {
		token := z.Next()

		if token == html.ErrorToken {
			urlChan <- ""
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
