// Exercise 4.9
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "[error] Word Counts: No file provided.\n")
		os.Exit(1)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "[error] Word Counts: %v\n", err)
				continue
			}
			countWords(f, counts)
			f.Close()
		}
	}
	fmt.Println(counts)
}

func countWords(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}
}
