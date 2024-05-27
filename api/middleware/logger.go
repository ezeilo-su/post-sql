package middleware

import (
	"log"
	"net/http"
)

func Logger(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v", r.Method, r.URL.Path)

		f(w, r)
	}
}
