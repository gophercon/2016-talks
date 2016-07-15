package main

import (
	"io"
	"log"
	"log/syslog"
	"os"
	"time"
)

func main() {
	// START OMIT
	sw, err := syslog.Dial("tcp", "logs4.papertrailapp.com:27193", 0, "gophercon2016")
	if err != nil {
		log.Println(err)
	} else {
		log.SetOutput(io.MultiWriter(sw, os.Stderr))
	}
	//END OMIT
	for {
		log.Println("Hello Gophers!")
		time.Sleep(1 * time.Second)
	}
}
