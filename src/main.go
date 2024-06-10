package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)

func getPath() (string, error) {
	// First element in os.Args is always the program name,
	// So we need at least 2 arguments to have a file name argument.
	if len(os.Args) < 2 {
		return "", errors.New("missing parameter, please provide path")
	}

	return os.Args[1], nil
}

func main() {
	path, err := getPath()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(path)

	// Create a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer watcher.Close()
}
