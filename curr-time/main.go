package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		rfcTime := currentTime.Format(time.RFC3339)

		var output []byte
		acceptHeader := r.Header.Get("Accept")
		if acceptHeader == "application/json" {
			w.Header().Set("Content-Type", "application/json")

			res := struct {
				DayOfWeek  string `json:"day_of_week"`
				DayOfMonth int    `json:"day_of_month"`
				Month      string `json:"month"`
				Year       int    `json:"year"`
				Hour       int    `json:"hour"`
				Minute     int    `json:"minute"`
				Second     int    `json:"second"`
			}{
				DayOfWeek:  currentTime.Weekday().String(),
				DayOfMonth: currentTime.Day(),
				Month:      currentTime.Month().String(),
				Year:       currentTime.Year(),
				Hour:       currentTime.Hour(),
				Minute:     currentTime.Minute(),
				Second:     currentTime.Second(),
			}

			out, err := json.Marshal(res)
			if err != nil {
				panic(err)
			}
			output = out
		} else {
			w.Header().Set("Content-Type", "text/plain")
			output = []byte(rfcTime)
		}
        w.WriteHeader(http.StatusOK)
		w.Write(output)
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
		Addr:    ":8080",
		Handler: mux,
	}

	slog.Info("Starting to listen on", "port", s.Addr)

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
