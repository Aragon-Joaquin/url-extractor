package scraper

import (
	"golang.org/x/net/html"
)

func ParseHtml(z *html.Tokenizer, urlChan chan<- string) {
	//! parse to html and find anchor

	for {
		token := (z).Next()

		if token == html.ErrorToken {
			urlChan <- ""
			break
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
