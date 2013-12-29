package main

import (
	"log"
	"net/http"
	"time"

	"github.com/reedobrien/rescat"
)

var files = []string{"ignore.txt", "ignore2.txt", "ignore3.txt"}
var c []byte

// Demonstrate using a file system handler at /testfiles/
func main() {
	http.Handle("/testfiles/", http.TimeoutHandler(&rescat.HandleFS{}, time.Nanosecond, ""))
	log.Fatal(http.ListenAndServe(":8888", nil))
}
