//demo is awesome
package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"regexp"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile("^(.+)@golang.org$") // beginning of the line + whatever + "golang.org" + end of the line
	path := r.URL.Path[1:]                        // path of the request
	match := re.FindAllStringSubmatch(path, -1)

	if match != nil {
		fmt.Fprintf(w, "Hello, gopher %s\n", match[0][1]) //match[0][0]is the whole string, match[1]is the username part
		return
	}

	fmt.Fprintf(w, "Hello, dear %s\n", path)
}
