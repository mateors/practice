package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {

	//os.StartProcess("",)
	//ec := exec.Command("ls")
	//ec.Start()
	//ec.Process.Pid

	//running process where we will send signal
	//get the process id using shell command
	//syntax: echo $(pgrep 'processName')
	//syntax: echo $(pgrep 'signal_reload')

	runningProcessId := 563290
	//strconv.Atoi()
	p, err := os.FindProcess(runningProcessId)
	if err != nil {
		fmt.Println(err.Error())
		//panic(err)
		return
	}

	err = p.Signal(syscall.SIGQUIT) //SIGUSR1 | SIGUSR2 | SIGALRM
	if err != nil {
		log.Fatal(err) //exit status 1
		//panic(err) //exit status 2
	}
	fmt.Println("Signal successfully sent")
}
