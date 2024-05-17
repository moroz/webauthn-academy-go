package handler

import (
	"fmt"
	"net/http"
)

func ParseForm(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			w.Header().Add("content-type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Bad Request")
			return
		}

		next.ServeHTTP(w, r)
	})
}
