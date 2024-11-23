package main

import (
	"net/http"
	"time"
)

func main() {
    h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        currentTime := time.Now()
        rfcTime := currentTime.Format(time.RFC3339)
        w.Write([]byte(rfcTime))
    })

    mux := http.NewServeMux()
    mux.Handle("GET /time", h)

    s := http.Server{
        Addr: ":8080",
        Handler: mux,
    }

    err := s.ListenAndServe()
    if err != nil {
        panic(err)
    }
}
