package main

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)

func main() {
	if len(os.Args) == 1 {
		helpMenu()
		return
	}

	path := os.Args[1]

	fmt.Print(path)

	// Create a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer watcher.Close()
}
