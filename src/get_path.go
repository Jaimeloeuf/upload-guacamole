package main

import (
	"errors"
	"os"
)

func getPath() (string, error) {
	// First element in os.Args is always the program name,
	// So we need at least 2 arguments to have a file name argument.
	if len(os.Args) < 2 {
		return "", errors.New("missing parameter, please provide path")
	}

	return os.Args[1], nil
}
