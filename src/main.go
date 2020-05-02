package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/deal", hello)
	http.ListenAndServe(":3000", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello")
}
