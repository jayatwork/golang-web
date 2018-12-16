# golang-web
Simple GO language web quickstart

@author jayAtWorkerBee

Source initially looks something like this below >>>

`$ cat server.go`
package main

import (
        "fmt"
        "html"
        "log"
        "net/http"
)

func main() {

        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
        })

        http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "Hi")
        })

        log.Fatal(http.ListenAndServe(":8081", nil))

}

To RUN ` $ go run server.go`
