// Echo2 is a beginner program to print cmdline args

package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
      	fmt.Println(s)
}
