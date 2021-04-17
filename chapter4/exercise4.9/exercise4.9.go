// Exercise 4.9 reports the frequency of each word in an input text file.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	words := WordFreq(os.Stdin)
	for word, count := range words {
		fmt.Printf("%d %s\n", count, word)
	}
}

// WordFreq returns the frequency of each word in input.
func WordFreq(input io.Reader) map[string]int {
	words := make(map[string]int)
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words[strings.ToLower(scanner.Text())]++
	}
	return words
}
