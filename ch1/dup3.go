// Dup3 prints the count and text of lines that appear more than once
// in the input It reads from stding or from a list of named files.
// Dup3 operates by reading everything into memory in one go.


package main

import (
       "fmt"
       "io/ioutil"
       "os"
       "strings"
)

func main() {
     counts := make(map[string]int)
     for _, filename := range os.Args[1:] {
     	 data, err := ioutil.ReadFile(filename)
	 if err != nil {
	    fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
	    continue
	 }
	 for _, line := range strings.Split(string(data), "\n") {
	     counts[line]++
	 }
     }
     for line, n := range counts {
     	 if n > 1 {
	    fmt.Printf("%d\t%s\n", n, line)
	 }
     }
}
