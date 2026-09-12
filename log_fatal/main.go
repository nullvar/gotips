package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	l, e := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "Some program")
	if e != nil {
		log.Fatal(e)
	} else {
		log.SetOutput(l)
	}
	log.Panic(l)
	fmt.Println("Will you see this?")
}
