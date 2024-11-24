package mid

import (
	"context"
	"net/http"
	"time"
)

func main() {

}

func timeoutMIddleware(ctx context.Context, duration time.Duration) func(http.Handler) http.Handler {
    return func(h http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx, cancel := context.WithTimeout(ctx, duration)
            defer cancel()
            r.WithContext(ctx)
            h.ServeHTTP(w, r)
        })
    }
}
