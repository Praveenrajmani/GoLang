package main

import "net/http"
import "fmt"
	

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Started")
	fmt.Fprintf(w, "Hello Bitches")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:9000", nil)
}

