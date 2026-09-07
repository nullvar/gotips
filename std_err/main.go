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
