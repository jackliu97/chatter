package handlers

import (
	"net/http"
	"fmt"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		// always pass for now
		pass := true

		if(pass) {
			next.ServeHTTP(w, r)
			return
		}

		fmt.Println("bad token")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("bad token"))
	})
}
