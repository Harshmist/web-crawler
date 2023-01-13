package main

import (
	"fmt"
	"github.com/harshmist/web-crawler/links"
	"net/http"
)

var url = "https://www.singlelink.co/"

func main() {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer resp.Body.Close()

	list, err := links.LinkFinder(resp)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Print(list)
}
