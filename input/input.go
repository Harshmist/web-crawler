package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InputReciever() (string, error) {
	fmt.Print("Enter a URL: ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occurred while reading input.", err)
		return "", err
	}

	// remove the delimiter
	input = strings.TrimSuffix(input, "\n")

	return input, nil
}
