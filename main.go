package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Product Catalog API</h1>")
}

func main() {
	http.HandleFunc("/", handleFunc)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}
