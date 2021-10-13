package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
A more practical and common example of a clean shutdown is resource
cleanup. When using exit statements, deferred functions such as Flush for
a bufio.Writer struct are not executed. This can lead to information loss,
*/
func main() {

	//regular way
	// f, err := os.OpenFile("file.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// w := bufio.NewWriter(f)
	// defer w.Flush()
	// for i := 0; i < 10; i++ {
	// 	fmt.Fprintln(w, "hello", i)
	// 	log.Println(i)
	// 	time.Sleep(time.Second)
	// }

	//best way
	//capturing signals
	c := make(chan os.Signal, syscall.SIGTERM)
	signal.Notify(c)

	f, err := os.OpenFile("file.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	go func() {
		<-c
		w.Flush() //no defer because of os.Exit(0)
		os.Exit(0)
	}()

	for i := 0; i < 10; i++ {
		fmt.Fprintln(w, "hello", i)
		log.Println(i)
		time.Sleep(time.Second)
	}
}
