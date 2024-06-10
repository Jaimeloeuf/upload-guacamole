package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

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
