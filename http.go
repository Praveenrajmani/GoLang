package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hell", handler)
	//http.HanleFunc("")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
