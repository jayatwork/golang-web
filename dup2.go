//Dup3 is simplified rev of original 
//Continue to print count and text of files appearing
//more than once from named files

package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"os"
)

func main() {
	counts :=  make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3:%v\n", err)
			continue
		}
	  	for _, line :=range strings.Split(string(data), "\n") {
			counts[line]++
	  	}
	}
	for line, n := range counts {
		if n > 1 {
		 	fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
