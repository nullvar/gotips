# Go Tip #1: Graceful Shutdown

This example shows how to gracefully shut down a Go application using `signal.NotifyContext`.

The application waits for `SIGINT` or `SIGTERM`. When a signal is received, the context is canceled and the application shuts down cleanly.

```go
package main

import (
    "context"
    "fmt"
    "os/signal"
    "syscall"
)

func main() {
    ctx, stop := signal.NotifyContext(
        context.Background(),
        syscall.SIGINT,
        syscall.SIGTERM,
    )
    defer stop()

    fmt.Println("running...")

    <-ctx.Done()

    fmt.Println("shutting down...")
}