package main

import (
	"fmt"
	"net/http"
	"test/pkg/home"
	"test/pkg/router"
	"test/pkg/tools"
)

func main() {
	home.SetupFileserver()
	tools.Create_database()
	http.HandleFunc("/", router.RoutingHandler)
	fmt.Println("Running server on port 8080:")
	http.ListenAndServe(":8080", nil)
}
