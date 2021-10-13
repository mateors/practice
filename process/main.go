package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
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

	//user info
	usr, err := user.LookupId(strconv.Itoa(os.Getuid()))
	if err != nil {
		fmt.Println(err)
	}

	ugrs, _ := usr.GroupIds()
	fmt.Println("userInfo: ", usr.Username, usr.HomeDir, usr.Gid, ugrs)

	//group info
	grp, err := user.LookupGroupId(strconv.Itoa(os.Getgid()))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("groupInfo:", grp.Name, grp.Gid)

	//obtain working directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("current working directory:", dir)
	adir := filepath.Join(dir, "process") // os.Args[0]
	fmt.Println("Application")

	//create a directory
	fmt.Println(os.ModeAppend, os.ModeDir, os.ModePerm.Perm())
	d := filepath.Join(adir, "test")
	if err := os.Mkdir(d, 0755); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Directory create", d)

	//child process, Accessing child properties
	// cmd := exec.Command("ls", "-l")
	// if err := cmd.Start(); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Cmd: ", cmd.Args[0])
	// fmt.Println("Args:", cmd.Args[1:])
	// fmt.Println("PID: ", cmd.Process.Pid)
	// cmd.Wait()

	//standard input
	b := bytes.NewBuffer(nil)
	cmd := exec.Command("ls")
	cmd.Stdin = b
	cmd.Stdout = os.Stdout
	//fmt.Fprintf(b, "Hello World! I'm using this memory address: %p\n", b)
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}

	//cmd.Output()
	fmt.Println("output", b.String())
	cmd.Wait()

	//fmt.Println("Pages size:", os.Getpagesize())
}
