// Package middleware http negroni
package middleware

import (
	"net/http"
)

// Cors adiciona os headers para suportar o CORS nos navegadores.
func Cors(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")
	w.Header().Set("Content-Type", "application/json, application/x-www-form-urlencoded")

	if r.Method == "OPTIONS" {
		return
	}

	next(w, r)
}
