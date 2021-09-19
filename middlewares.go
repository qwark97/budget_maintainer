package main

import (
	"net/http"
)

func headersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", ALLOW_ORIGIN)
		w.Header().Add("Access-Control-Allow-Methods", ALLOW_ORIGIN)
		w.Header().Add("Access-Control-Allow-Headers", ALLOW_ORIGIN)
		next.ServeHTTP(w, r)
	})
}

func optionsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
