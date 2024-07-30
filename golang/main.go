package main

import (
	"fmt"
	"net/http"
	"test/router"
)

func main() {
	http.HandleFunc("/", router.RoutingHandler)
	fmt.Println("Running server on port 8080:")
	http.ListenAndServe(":8080", nil)
}
