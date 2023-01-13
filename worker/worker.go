package worker

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/harshmist/web-crawler/input"
	"github.com/harshmist/web-crawler/links"
)

func InitWorker() {
	for {
		input, err := input.InputReciever()
		if err != nil || strings.ToUpper(input) == "EXIT" {
			break
		}
		resp, err := http.Get(input)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		list, err := links.LinkFinder(resp)

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Print(list)
		resp.Body.Close()
	}
}
