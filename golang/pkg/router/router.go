package router

import (
	"fmt"
	"net/http"
	"test/pkg/about"
	"test/pkg/home"
	"test/pkg/product"
	"test/pkg/signup"
)

func RoutingHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/about":
		about.AboutPageHandler(w, r)
	case "/product":
		product.ProductPageHandler(w, r)
	case "/home":
		home.HomePageHandler(w, r)
	case "/signup":
		signup.SignupPageHandler(w, r)
	default:
		fmt.Fprintf(w, "Page not found")
	}
}
