package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	n := filepath.Base(os.Args[0])
	l, e := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, n)
	if e != nil {
		log.Fatal(e)
	} else {
		log.SetOutput(l)
	}
	log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go")

	l, e = syslog.New(syslog.LOG_MAIL, "Some program!")
	if e != nil {
		log.Fatal(e)
	} else {
		log.SetOutput(l)
	}
	log.Println("LOG_MAIL: Logging in GO")
	fmt.Println("Will you see this?")
}
