package scraper

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

const (
	ANCHOR_TAG     string        = "a"
	HREF_ATTRIBUTE string        = "href"
	HTTP_TIMEOUT   time.Duration = 5 * time.Second
)

func FetchUrl(ctx context.Context, link string) (*html.Tokenizer, error) {
	//! fetch
	client := &http.Client{
		Timeout: HTTP_TIMEOUT,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", link, nil)
	if err != nil {
		return nil, errors.New("Failed to create request: " + err.Error())
	}

	res, err := client.Do(req)

	if err != nil {
		if ctx.Err() != nil {
			return nil, errors.New(ctx.Err().Error())
		} else {
			return nil, errors.New("Request longed too much: " + HTTP_TIMEOUT.String())
		}
	}

	defer res.Body.Close()

	// i need to copy the body cuz newTokenizer reads directly from the res.Body
	// and when i pass it as a pointer, the body is already close cuz the defer
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return html.NewTokenizer(bytes.NewReader(bodyBytes)), nil
}
