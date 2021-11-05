package main

import (
	"fmt"
	"strings"
)

func staircase(n int32) {
	// Write your code here

	//fmt.Printf("%10s\n", "")
	//fmt.Printf("+%s+\n", strings.Repeat("-", 5))
	var i int32
	for i = 1; i <= n; i++ {
		space := strings.Repeat(" ", int(n-i))
		hash := strings.Repeat("#", int(i))
		stair := fmt.Sprintf("%s%s", space, hash)
		fmt.Println(stair)
	}
}

func main() {

	staircase(6)
}
