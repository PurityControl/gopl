package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, word := range os.Args[1:] {
		fmt.Println(idx, word)
	}
}
