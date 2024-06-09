package main

import (
	"errors"
	"fmt"
	"os"
)

func getPath() (string, error) {
	// First element in os.Args is always the program name,
	// So we need at least 2 arguments to have a file name argument.
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, please provide path")
		return "", errors.New("missing parameter, please provide path")
	}

	return os.Args[1], nil
}

func main() {
	path, err := getPath()

	if err != nil {
		return
	}

	fmt.Print(path)
}
