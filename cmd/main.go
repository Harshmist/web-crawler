package main

import (
	"fmt"
	"github.com/harshmist/web-crawler/input"
	"github.com/harshmist/web-crawler/links"
	"net/http"
	"strings"
)

const (
	welcomeLogo    = "██╗    ██╗███████╗██████╗      ██████╗██████╗  █████╗ ██╗    ██╗██╗     ███████╗██████╗ \n██║    ██║██╔════╝██╔══██╗    ██╔════╝██╔══██╗██╔══██╗██║    ██║██║     ██╔════╝██╔══██╗\n██║ █╗ ██║█████╗  ██████╔╝    ██║     ██████╔╝███████║██║ █╗ ██║██║     █████╗  ██████╔╝\n██║███╗██║██╔══╝  ██╔══██╗    ██║     ██╔══██╗██╔══██║██║███╗██║██║     ██╔══╝  ██╔══██╗\n╚███╔███╔╝███████╗██████╔╝    ╚██████╗██║  ██║██║  ██║╚███╔███╔╝███████╗███████╗██║  ██║\n ╚══╝╚══╝ ╚══════╝╚═════╝      ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚══╝╚══╝ ╚══════╝╚══════╝╚═╝  ╚═╝\n                                                                                        "
	welcomeMessage = "Welcome to the web-crawler!\n\n* When prompted, enter a URL to find and display all the links from the page\n* All URL's must include the correct protocol\n* Type \"exit\" to close the application\n"
)

func main() {

	fmt.Println(welcomeLogo)
	fmt.Println(welcomeMessage)
	for {
		input, err := input.InputReciever()
		if err != nil || strings.ToUpper(input) == "EXIT" {
			break
		}
		resp, err := http.Get(input)

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

	fmt.Println("Goodbye!")

}
