package about

import (
	"fmt"
	"net/http"
)

func AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page")
}
