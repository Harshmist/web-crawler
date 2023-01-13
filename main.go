package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

var url = ""

func main() {
	resp, _ := http.Get(url)

	defer resp.Body.Close()

	var links []string
	tokens := html.NewTokenizer(resp.Body)
	for {
		tokenType := tokens.Next()

		switch tokenType {
		case html.ErrorToken:
			break
		case html.StartTagToken, html.EndTagToken:
			token := tokens.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}

				}
			}

		}
		if tokens.Err() != nil {
			break
		}
	}

	fmt.Print(links)
}
