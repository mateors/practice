package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {

	s := time.Now()

	// expire_date=1279012036
	apiEndpoint := "https://test.httpapi.com/api/domains/renew.json?auth-userid=%s&api-key=%s&order-id=%s&years=1&exp-date=%s&invoice-option=%s&discount-amount=0.0"

	//Method-1
	contentType := "application/json"
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Mostain Billah",
		"email": "mostain@mateors.com",
	})

	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(apiEndpoint, contentType, responseBody)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//sb := string(body)

	unixtoDate("1633242596")
	etime, _ := epocDate("2021-10-03 12:29:56")
	fmt.Println(etime) //1633264196 1633264196

	fmt.Println("Response:", string(body))
	fmt.Println("TimeTaken:", time.Since(s).Milliseconds(), "ms")

	//Method-2
	// frmVals := make(url.Values)
	// frmVals.Add("name", "mostain")
	// frmVals.Add("age", "37")
	//http.PostForm(apiEndpoint, frmVals)
}

// /1633242596
func epocDate(dateTime string) (int64, error) {

	//Mon Jan _2 15:04:05 MST 2006
	//thetime, err := time.Parse("2006-01-02 15:04:05", dateTime)
	thetime, err := time.Parse("2006-01-02 15:04:05", dateTime)
	if err != nil {
		return 0, err
	}
	epoch := thetime.Unix()
	return epoch, nil
}

func unixtoDate(unixTimeStr string) {

	i, err := strconv.ParseInt(unixTimeStr, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	fmt.Println(tm.Format("2006-01-02 15:04:05"))

}
