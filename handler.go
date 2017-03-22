package blankdefault

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", notFound)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not found", http.StatusNotFound)
}
