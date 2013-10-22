package main

import (
	"github.com/reedobrien/rescat"
	"log"
	"net/http"
	"time"
)

var files = []string{"ignore.txt", "ignore2.txt", "ignore3.txt"}
var c []byte

func main() {
	// This should be a
	http.Handle("/testfiles/", http.TimeoutHandler(&rescat.HandleFS{}, time.Nanosecond, ""))
	log.Fatal(http.ListenAndServe(":8888", nil))
}
