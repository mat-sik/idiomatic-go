package main

import (
	"log/slog"
	"net/http"
	"time"
)

func main() {
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        currentTime := time.Now()
        rfcTime := currentTime.Format(time.RFC3339)
        w.Write([]byte(rfcTime))
    })

    loggingMiddleware := func(h http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            slog.Info("Received request", "from", r.RemoteAddr)
            h.ServeHTTP(w, r) 
        })
    }

    mux := http.NewServeMux()
    mux.Handle("GET /time", loggingMiddleware(handler))

    s := http.Server{
        Addr: ":8080",
        Handler: mux,
    }

    slog.Info("Starting to listen on", "port", s.Addr)

    err := s.ListenAndServe()
    if err != nil {
        panic(err)
    }
}
