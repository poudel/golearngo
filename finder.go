package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func parseArgs() (string, string) {
	/*
	   Parse the arguments and return the path and the pattern.
	*/

	args_length := len(os.Args)

	if args_length == 1 {
		log.Fatal("Usage: finder [pattern] path")
	}

	pattern := os.Args[1]
	fpath := "."

	if args_length == 3 {
		fpath = os.Args[2]
	}

	absolute_path, err := filepath.Abs(fpath)

	if err != nil {
		log.Fatal("Can't get absolute path: %v", err)
	}

	return absolute_path, pattern
}

func findFiles(dir, pattern string) {
	/*
		check if the filename contains the pattern and print the path if it does
	*/

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal("You can't handle the truth: %v", err)
		}

		if !info.IsDir() && strings.Contains(filepath.Base(path), pattern) {
			fmt.Println(path)
		}

		if info.IsDir() && dir != path {
			findFiles(path, pattern)
		}
		return err
	})

	if err != nil {
		log.Fatal("Error walking: %v", err)
	}
}

func main() {
	findFiles(parseArgs())
}
