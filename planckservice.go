package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
)

var digit = regexp.MustCompile("^/([01])$")
var nand = regexp.MustCompile("/[01]/[01]")

func main() {
	var addr = flag.String("addr", ":8000", "Address to listen on")
	flag.Parse()
	http.HandleFunc("/", handler)
	http.ListenAndServe(*addr, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "%v", rand.Intn(2))
	} else if r.URL.Path == "/1/1" {
		fmt.Fprintf(w, "0")
	} else if digit.MatchString(r.URL.Path) {
		fmt.Fprintf(w, "%v", digit.FindAllStringSubmatch(r.URL.Path, 1)[0][1])
	} else if nand.MatchString(r.URL.Path) {
		fmt.Fprintf(w, "1")
	} else {
		http.NotFound(w, r)
	}
}
