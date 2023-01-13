package links

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type LinkList []string

func LinkFinder(response *http.Response) (string, error) {
	var links LinkList
	var err error

	tokens := html.NewTokenizer(response.Body)
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

	if len(links) < 1 {
		err = fmt.Errorf("no links found from provided url")
		return "", err
	}

	listString := listFormatter(links)

	return listString, nil
}

func listFormatter(links LinkList) string {
	listString := "\nHere are your links:\n"

	for _, v := range links {
		if !strings.Contains(v, "://") {
			continue
		}
		listString = fmt.Sprintf("%s%s\n", listString, v)
	}

	return listString + "\n"
}
