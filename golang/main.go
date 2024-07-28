package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePageHandler)
	fmt.Println("Runnin server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello This is the home page")
}
