package p

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Handle ...
func Handle(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	if d.Message == "" {
		fmt.Fprint(w, "Hello World!")
		fmt.Fprint(w, os.Getenv("GCP_PROJECT"))
		return
	}
	// fmt.Fprint(w, html.EscapeString(d.Message))
	fmt.Fprint(w, os.Getenv("GCP_PROJECT"))
}
