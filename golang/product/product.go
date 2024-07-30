package product

import (
	"fmt"
	"net/http"
)

func ProductPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the product page")
}
