package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(h http.HandlerFunc) http.Handler {
	// TODO Currently only logs what was requested
	// would like to know what I responded with

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqTime := time.Now()

		h.ServeHTTP(w, r)

		// TODO Include POST Form values
		log.Printf("%s %s %s %s\n", r.RemoteAddr, r.Method, r.URL, time.Since(reqTime))
	})
}
