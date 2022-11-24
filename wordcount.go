package main

import (
	"fmt"
	"os"
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
func readInput() (src []string, err error) {
	src = os.Args[1:]
	/*	if len(src) == 0 {
			return src, errors.New("missing words to count")
		}
	*/
	return src, nil
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("wordcount:", err)
	os.Exit(1)
}
