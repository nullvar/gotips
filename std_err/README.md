# Go Stdout / Stderr Example

A simple Go program that demonstrates how to write messages to **standard output (`stdout`)** and **standard error (`stderr`)**.

## What it does

The program:

- Writes a message to `stdout`.
- Reads the first command-line argument.
- Writes the argument to `stderr`.
- If no argument is provided, it writes an error message to `stderr`.

## Code

```go
package main

import (
    "io"
    "os"
)

func main() {
    e := ""

    if len(os.Args) == 1 {
        e = "Please give me one argument"
    } else {
        e = os.Args[1]
    }

    io.WriteString(os.Stdout, "This is Standard output\n")
    io.WriteString(os.Stderr, e+"\n")
}
```

## Run

Run the program with an argument:

```bash
go run main.go "Hello"
```

Output:

```text
This is Standard output
Hello
```

The first line is written to `stdout`, while `Hello` is written to `stderr`.

## Run without an argument

```bash
go run main.go
```

Output:

```text
This is Standard output
Please give me one argument
```

## Redirect stdout

You can redirect standard output to a file:

```bash
go run main.go "Hello" > stdout.txt
```

`stdout.txt`:

```text
This is Standard output
```

The `stderr` message will still be displayed in the terminal:

```text
Hello
```

## Redirect stderr

To redirect standard error:

```bash
go run main.go "Hello" 2> stderr.txt
```

`stderr.txt`:

```text
Hello
```

## Redirect both

```bash
go run main.go "Hello" > stdout.txt 2> stderr.txt
```

This creates:

```text
stdout.txt → This is Standard output
stderr.txt → Hello
```

## Build

```bash
go build -o stdout-example .
```

Then run:

```bash
./stdout-example "Hello"
```