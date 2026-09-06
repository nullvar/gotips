# CLI Arguments

A simple Go example that demonstrates how to work with command-line arguments using `os.Args`.

## Description

The program reads arguments passed from the command line and prints each argument with its index.

If no arguments are provided, the program prints an error message and exits with status code `1`.

## Run

```bash
go run main.go hello world
```

Output:

```text
Arg[1] hello
Arg[2] world
```

## No Arguments

If you run the program without arguments:

```bash
go run main.go
```

the program exits with an error:

```text
Please provide at least one argument
```

## Build

Build the executable:

```bash
go build -o cli-args main.go
```

Then run it:

```bash
./cli-args hello world
```

Output:

```text
Arg[1] hello
Arg[2] world
```

## Key Concept

Go provides command-line arguments through:

```go
os.Args
```

- `os.Args[0]` — program name
- `os.Args[1]` — first argument
- `os.Args[2]` — second argument
- etc.

The program starts iterating from index `1` because index `0` contains the program name.