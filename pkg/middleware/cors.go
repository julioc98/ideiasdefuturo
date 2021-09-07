// Package middleware http negroni
package middleware

import (
	"log"
	"net/http"
)

// Cors adiciona os headers para suportar o CORS nos navegadores.
func Cors(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Cors middleware")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Access-Control-Allow-Headers, Origin, X-Requested-With, Access-Control-Request-Method, Access-Control-Request-Headers")
	w.Header().Set("Access-Control-Request-Headers", "Accept, Authorization, Content-Type, Access-Control-Allow-Headers, Origin, X-Requested-With, Access-Control-Request-Method, Access-Control-Request-Headers")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		return
	}

	next(w, r)
}
