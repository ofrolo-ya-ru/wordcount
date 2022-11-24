package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	src, err := readInput()
	if err != nil {
		fail(err)
	}
	fmt.Println(len(src))
}

// readInput reads source string
// from command line arguments and returns them.
func readInput() ([]string, error) {
	src := make([]string, 0)
	params := os.Args
	if len(params) == 1 {
		return src, nil
	}
	for _, s := range params[1:] {
		for _, si := range strings.Split(s, " ") {
			if si != "" {
				src = append(src, si)
			}
		}
	}
	// fmt.Println(src)
	return src, nil
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("wordcount:", err)
	os.Exit(1)
}
