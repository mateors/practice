package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

var cfgPath string

//We can launch the application in a Terminal
//and use another Terminal to send signals, as follows:

//kill -SIGUSR1 $(pgrep signal_reload)

//kill -SIGUSR2 $(pgrep signal_reload)

//kill -SIGINT $(pgrep signal_reload)

func init() {

	// u, err := user.Current()
	// if err != nil {
	// 	log.Fatalln("user:", err)
	// }

	wd, _ := os.Getwd()
	cfgPath = filepath.Join(wd, ".multi")
	fmt.Println("configurationPath:", cfgPath)
}

func main() {

	c := make(chan os.Signal, 1)
	d := time.Second * 5

	signal.Notify(c,
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT,
		syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGALRM)

	// initial load
	if err := handleSignal(syscall.SIGHUP, &d); err != nil &&
		!os.IsNotExist(err) {
		log.Fatal(err)
	}

	for {
		select {
		case s := <-c:
			if err := handleSignal(s, &d); err != nil {
				log.Printf("Error handling %s: %s", s, err)
				continue
			}
		default:
			time.Sleep(d)
			log.Println("After", d, "Executing action!")
		}
	}
}

func handleSignal(s os.Signal, d *time.Duration) error {

	switch s {
	case syscall.SIGHUP:
		return loadSettings(d)

	case syscall.SIGALRM:
		return saveSettings(d)

	case syscall.SIGINT:
		if err := saveSettings(d); err != nil {
			log.Println("Cannot save:", err)
			os.Exit(1)
		}
		fallthrough
	case syscall.SIGQUIT:
		os.Exit(0)
	case syscall.SIGUSR1:
		changeSettings(d, (*d)*2)
		return nil
	case syscall.SIGUSR2:
		changeSettings(d, (*d)/2)
		return nil
	}
	return nil
}

func changeSettings(d *time.Duration, v time.Duration) {
	*d = v
	log.Println("Changed", v)
}

func loadSettings(d *time.Duration) error {
	b, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return err
	}
	var v time.Duration
	if v, err = time.ParseDuration(string(b)); err != nil {
		return err
	}
	*d = v
	log.Println("Loaded", v)
	return nil
}
func saveSettings(d *time.Duration) error {
	f, err := os.OpenFile(cfgPath,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = fmt.Fprint(f, d); err != nil {
		return err
	}
	log.Println("Saved", *d)
	return nil
}
