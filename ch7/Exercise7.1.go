package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	input1 := "Hello Go!"
	input2 := "Hello Go!\n This is a second line.\n This is a third line\n"
	var r WordCounter
	r.Read([]byte(input1))
	fmt.Println(r)

	var rl LineCounter
	rl.Read([]byte(input2))
	r.Read([]byte(input2))
	fmt.Println(r)
	fmt.Println(rl)
}

//Those types implement the io.Reader interface
type WordCounter int
type LineCounter int

func (wc *WordCounter) Read(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {
		count++
	}
	var err error
	if err = scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input: ", err)
	}
	*wc = WordCounter(count)
	return count, err
}

func (lc *LineCounter) Read(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	//no need for this line, ScanLines is the default split function for Scanner
	scanner.Split(bufio.ScanLines)

	countLines := 0
	for scanner.Scan() {
		countLines++
	}
	*lc = LineCounter(countLines)
	var err error
	if err = scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input: ", err)
	}
	return countLines, err
}
