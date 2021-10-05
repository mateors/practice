package main

import (
	"fmt"
	"os"
)

func main() {

	fsL, err := os.ReadDir(".") ///home/mostain/go/src/practice
	if err != nil {
		fmt.Println(err)
	}
	for i, v := range fsL {

		//v.Info()
		fmt.Println(i, v.IsDir(), v.Name(), v.Type().Perm())

	}

	//open the folder as as file, much faster/efficient solution
	fd, err := os.Open("..")
	if err != nil {
		fmt.Println(err)
	}

	fsList, err := fd.Readdir(1024)
	if err != nil {
		fmt.Println(err)
	}
	for i, file := range fsList {
		fmt.Println(i, file.Name(), file.IsDir(), file.Mode().Perm())
	}
}
