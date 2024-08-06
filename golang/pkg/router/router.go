package router

import (
	"fmt"
	"net/http"
	"test/pkg/about"
	"test/pkg/home"
	"test/pkg/product"
	"test/pkg/signup"
	"test/pkg/tools"
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
	case "/tools":
		tools.GenerateCSRFTPageHandler(w, r)
	case "/setcookie":
		tools.SetCookieTwo(w, r)
	case "/readcookie":
		tools.ReadCookie(w, r)
	case "/expirecookie":
		tools.ExpireCookie(w, r)
	default:
		fmt.Fprintf(w, "Page not found")
	}
}
