package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Current PID:", os.Getpid())
	fmt.Println("Current Parent PID:", os.Getppid())

	fmt.Println("userID:", os.Getuid())
	fmt.Println("groupID:", os.Getgid())

	groups, err := os.Getgroups()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("GroupIDS:", groups)

	fmt.Println("Pages size:", os.Getpagesize())
}
