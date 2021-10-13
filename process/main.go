package main

import (
	"fmt"
	"os"
	"os/user"
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

	fmt.Println("Pages size:", os.Getpagesize())
}
