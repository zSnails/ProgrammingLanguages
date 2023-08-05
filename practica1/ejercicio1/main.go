package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var filename string

func init() {
	flag.StringVar(&filename, "file", "", "The file to read")
	flag.Parse()
}

// getData returns the total characters, words and lines in a given string
func getData(input string) (chars int, words int, lines int) {
	chars = len(input)
	words = len(strings.Fields(input))
	lines = len(strings.Split(input, "\n"))
	return
}

func main() {
	if filename == "" {
		panic(errors.New("the file parameter must not be empty"))
	}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	chars, words, lines := getData(string(data))
	fmt.Printf("Characters: %d\nWords: %d\nLines: %d\n", chars, words, lines)
}
