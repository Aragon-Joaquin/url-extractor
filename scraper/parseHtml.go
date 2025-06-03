package scraper

import (
	u "url-extractor/utils"

	"golang.org/x/net/html"
)

func ParseHtml(z *html.Tokenizer, urlChan chan<- string, domain string) {
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
				anchor := getAttrHref(&element)
				urlChan <- u.RepairPath(domain, anchor)
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
