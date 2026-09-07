# Go Syslog Example

A simple Go program demonstrating how to send log messages to the system **syslog** service using the standard `log/syslog` package.

## Description

The program creates two syslog writers with different facilities:

- `LOG_LOCAL7` with `LOG_INFO` priority
- `LOG_MAIL`

It also demonstrates the difference between writing messages to syslog and writing directly to standard output (`stdout`).

## Requirements

- Go
- Linux or another Unix-like system with a running syslog service
- `rsyslog`, `syslog-ng`, or another compatible syslog daemon

Check your Go version:

```bash
go version
```

## Run

Run the program directly:

```bash
go run main.go
```

Or build it first:

```bash
go build -o syslog-example main.go
```

Then run:

```bash
./syslog-example
```

## Expected Output

The following message is written directly to standard output:

```text
Will you see this?
```

The other messages are sent to syslog:

```text
LOG_INFO + LOG_LOCAL7: Logging in Go
LOG_MAIL: Logging in GO
```

## How It Works

The first syslog writer is created with:

```go
syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, n)
```

The program name is taken from:

```go
filepath.Base(os.Args[0])
```

The standard Go logger is then redirected to syslog:

```go
log.SetOutput(l)
```

After that:

```go
log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go")
```

is sent to the system logging service.

The program then creates another syslog writer:

```go
syslog.New(syslog.LOG_MAIL, "Some program!")
```

and redirects the logger to it.

## Viewing Logs

Depending on your Linux distribution and syslog configuration, logs can usually be inspected with:

```bash
journalctl
```

or:

```bash
tail -f /var/log/syslog
```

You can search for the program:

```bash
grep "Logging in Go" /var/log/syslog
```

For systems using `journalctl`:

```bash
journalctl | grep "Logging in Go"
```

The exact destination of `LOG_LOCAL7` and `LOG_MAIL` messages depends on the system's syslog configuration.

## Important

The `log/syslog` package is available only on Unix-like operating systems.

The program requires access to a local syslog daemon. If no syslog service is available, `syslog.New()` may return an error and the program will terminate with `log.Fatal()`.

## Example

```bash
$ go run main.go
Will you see this?
```

The terminal shows only the `fmt.Println()` output, while the `log.Println()` messages are sent to syslog.

## License

This example is free to use for learning and testing purposes.