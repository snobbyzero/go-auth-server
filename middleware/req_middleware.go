package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		reader := io.NopCloser(bytes.NewReader(body))
		req.Body = reader
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s %s %s", start.Format(time.RFC822), req.Method, req.RequestURI, time.Since(start), string(body))
	})
}
