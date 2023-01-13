package main

import (
	"fmt"

	"github.com/harshmist/web-crawler/worker"
)

const (
	welcomeLogo    = "██╗    ██╗███████╗██████╗      ██████╗██████╗  █████╗ ██╗    ██╗██╗     ███████╗██████╗ \n██║    ██║██╔════╝██╔══██╗    ██╔════╝██╔══██╗██╔══██╗██║    ██║██║     ██╔════╝██╔══██╗\n██║ █╗ ██║█████╗  ██████╔╝    ██║     ██████╔╝███████║██║ █╗ ██║██║     █████╗  ██████╔╝\n██║███╗██║██╔══╝  ██╔══██╗    ██║     ██╔══██╗██╔══██║██║███╗██║██║     ██╔══╝  ██╔══██╗\n╚███╔███╔╝███████╗██████╔╝    ╚██████╗██║  ██║██║  ██║╚███╔███╔╝███████╗███████╗██║  ██║\n ╚══╝╚══╝ ╚══════╝╚═════╝      ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚══╝╚══╝ ╚══════╝╚══════╝╚═╝  ╚═╝\n                                                                                        "
	welcomeMessage = "Welcome to the web-crawler!\n\n* When prompted, enter a URL to find and display all the links from the page\n* All URL's must include the correct protocol\n* Type \"exit\" to close the application\n"
	exitMessage    = "Goodbye"
)

func main() {
	startUpScreen()

	worker.InitWorker()

	fmt.Println(exitMessage)
}

func startUpScreen() {
	fmt.Println(welcomeLogo)
	fmt.Println(welcomeMessage)
}
