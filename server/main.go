package main

import (
	"github.com/reedobrien/rescat"
	"log"
	"net/http"
)

var files = []string{"ignore.txt", "ignore2.txt", "ignore3.txt"}
var c []byte

func main() {
	http.HandleFunc("/", rescat.HandleFS)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Server", "SCRONK")
// 	for _, f := range files {
// 		fh, err := ioutil.ReadFile(f)
// 		if err != nil {
// 			log.Fatalln("No such file")
// 		}
// 		c = append(c, fh...)
// 	}
// 	w.Write(c)
// }
