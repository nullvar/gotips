package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide at least on argument")
		os.Exit(1)
	}

	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("Arg[%d] %s \n", i, os.Args[i])
	}
}
