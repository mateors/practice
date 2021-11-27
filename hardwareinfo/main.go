package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/user"

	"github.com/zcalusic/sysinfo"
)

func main() {

	current, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	if current.Uid != "0" {
		log.Fatal("requires superuser privilege")
	}

	var si sysinfo.SysInfo

	si.GetSysInfo()

	data, err := json.MarshalIndent(&si, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
