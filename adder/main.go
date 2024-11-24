package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"time"
)

func main() {
    ctx := context.Background()
    target := 1234
    s, c, err := adder(ctx, target, 2 * time.Second)

    cause := fmt.Sprintf("number: %d was found", target)
    if err != nil {
        cause = err.Error()
    }

    slog.Info("adder results", "sum", s, "iterations", c, "cause", cause)
}

func adder(ctx context.Context, target int, timeout time.Duration) (int, int, error) {
    ctx, cancel := context.WithTimeout(ctx, timeout)
    defer cancel()

    iterCount := 0
    numSum := 0
    for {
        if err := ctx.Err(); err != nil {
            return numSum, iterCount, err
        }
        num := start + rand.Intn(end - start)
        numSum += num
        iterCount++
        if num == target {
            break
        }
    }

    return numSum, iterCount, nil
}

var (
    start = 0
    end = 100_000_000
)
