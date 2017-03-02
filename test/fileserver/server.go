package main

import (
	"flag"
	"net/http"
)

var addr = flag.String("http", ":8888", "http")
var dir = flag.String("path", ".", "path")

func main() {
	flag.Parse()
	http.ListenAndServe(*addr, http.FileServer(http.Dir(*dir)))
}
