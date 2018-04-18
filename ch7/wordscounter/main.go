package main

import (
	"bufio"
	"fmt"
	"strings"
)

// WordsCounter 计数器
type WordsCounter int

func (w *WordsCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p[:])))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	*w = WordsCounter(count)
	return count, nil
}

func main() {
	var w WordsCounter
	s := "hello world"
	fmt.Fprintf(&w, "%s Steven", s)
	fmt.Println(w)

}
