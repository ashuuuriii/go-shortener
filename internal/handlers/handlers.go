package handlers

import (
	"fmt"
	"net/http"
)

func Shorten(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "shorten")
}
