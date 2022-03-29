// Exercise 4.9 reports the frequency of each word in an input text file.
package exercise4_9

import (
	"bufio"
	"io"
	"strings"
)

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
