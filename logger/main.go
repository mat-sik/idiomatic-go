package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {
}

func LogMiddleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        level := r.URL.Query().Get(QueryLogLevelKey)

        ctx := ContextWithLogLevel(r.Context(), Level(level))

        r = r.WithContext(ctx)

        h.ServeHTTP(w, r)
    })
}

func Log(ctx context.Context, level Level, message string) {
    var inLevel Level
    // TODO get a logging level out of the context and assign it to inLevel
    inLevel, ok := LogLevelFromContext(ctx)
    if !ok {
        return
    }
    if level == Debug && inLevel == Debug {
        fmt.Println(message)
    }
    if level == Info && (inLevel == Debug || inLevel == Info) {
        fmt.Println(message)
    }
}

func ContextWithLogLevel(ctx context.Context, logLevel Level) context.Context {
    return context.WithValue(ctx, LogLevelKey{}, logLevel) 
}

func LogLevelFromContext(ctx context.Context) (Level, bool) {
    logLevel, ok := ctx.Value(LogLevelKey{}).(Level)
    return logLevel, ok
}

type Level string

const (
    Debug Level = "debug"
    Info Level = "info"
)

type LogLevelKey struct{}

const QueryLogLevelKey = "log_level"
