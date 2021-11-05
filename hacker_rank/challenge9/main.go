package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {

	// Write your code here
	slc := strings.SplitN(s, ":", 3)
	//fmt.Println(slc)
	hour := slc[0]
	minute := slc[1]
	second := slc[2][:2] //slice uptp PM
	AMPM := slc[2][2:]
	//fmt.Println(AMPM)
	//fmt.Println(hour, minute, second)

	if AMPM == "PM" {

		// slc := strings.SplitN(s, ":", 3)
		// fmt.Println(slc)
		// hour := slc[0]
		// minute := slc[1]
		// second := slc[2][:2] //slice uptp PM
		// fmt.Println(hour, minute, second)
		if hour != "12" {
			//hour = "01"
			hourn, _ := strconv.Atoi(hour)
			hour = fmt.Sprintf("%v", hourn+12)
		}

	} else if AMPM == "AM" {

		if hour == "12" {
			hour = "00"
		}
	}

	return fmt.Sprintf("%s:%s:%s", hour, minute, second)
}

func main() {

	//s := "12:01:00PM"
	//s := "12:01:00AM"
	//s := "07:01:00AM"
	//s := "07:01:00PM"
	//s := "11:01:00PM"
	s := "07:05:45PM"
	militaryFormat := timeConversion(s)
	fmt.Println(militaryFormat)

}
