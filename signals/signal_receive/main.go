package main

import (
	"log"
	"os"
	"os/signal"
)

//Graceful shutdowns
func main() {
	pid := os.Getpid()
	log.Println("Start application...", pid)
	c := make(chan os.Signal)

	signal.Notify(c)
	s := <-c

	log.Println("Exit with signal:", s)
}
