# Go Stdin Scanner

A simple Go application that reads input from `stdin` line by line and prints each line with a `>` prefix.

## How it works

The program uses `bufio.Scanner` to read data from standard input:

```go
scanner := bufio.NewScanner(os.Stdin)

for scanner.Scan() {
    fmt.Println(">", scanner.Text())
}
```

For example, if the input is:

```text
Hello
World
```

the output will be:

```text
> Hello
> World
```

## Requirements

- Go 1.20+

## Run

Clone the repository and run:

```bash
go run main.go
```

Then enter text:

```text
Hello
> Hello

Test message
> Test message
```

Press `Ctrl+D` on Linux/macOS to send EOF and stop the program.

## Pipe input

The application can also receive input from another command:

```bash
echo "Hello World" | go run main.go
```

Output:

```text
> Hello World
```

You can also read from a file:

```bash
cat input.txt | go run main.go
```

## Project structure

```text
.
├── main.go
└── README.md
```

## License

MIT