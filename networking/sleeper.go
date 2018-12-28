//Program directly from ch7 demos
package main

import (
	"fmt"
	"time"
	"flag"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v ...", *period)
	time.Sleep(*period)
	fmt.Printf("Done")
}
