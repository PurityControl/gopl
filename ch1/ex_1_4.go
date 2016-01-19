// ex_1_4 prints the count and text of lines that appear more than once
// in the input It reads from stding or from a list of named files.
// Dup2 operates in a streaming mode potentially allowing an arbitary
// amount of data.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fnames := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fnames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fnames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, fnames[line])
		}
	}
}


func countLines(f *os.File, counts map[string]int, fnames map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if len(fnames[input.Text()]) > 0 {
			if !(fnames[input.Text()][len(fnames[input.Text()])-1] == f.Name()) {
				fnames[input.Text()] = append(fnames[input.Text()], f.Name())
			}
		} else {
			fnames[input.Text()] = append(fnames[input.Text()], f.Name())
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
