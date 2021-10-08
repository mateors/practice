package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	s := time.Now()
	httpRes, err := http.Get("https://mateors.com")
	if err != nil {
		fmt.Println(err.Error())
	}

	//method-1
	// var bs = make([]byte, 22398) //
	// httpres.Body.Read(bs)

	// for key, val := range httpres.Header {
	// 	fmt.Println(key, val)
	// }

	//method-2  TimeTaken: 400 ms
	bodyBuffer2 := new(bytes.Buffer)
	b, err := io.Copy(ioutil.Discard, httpRes.Body) //time: 300 ms
	//b, err := io.Copy(bodyBuffer2, httpRes.Body)
	fmt.Println(b, err)
	fmt.Println(bodyBuffer2)

	//method-3 (problemetic)
	// maxmem := 1000
	// var bfr bytes.Buffer
	// ioreader := io.TeeReader(httpRes.Body, &bfr)

	// readsize := func(r io.Reader) int {
	// 	bytes := make([]byte, maxmem)
	// 	var size int
	// 	for {
	// 		read, err := r.Read(bytes)
	// 		if err == io.EOF {
	// 			//fmt.Println(err)
	// 			break
	// 		}
	// 		size += read
	// 		//fmt.Println(string(bytes))
	// 	}
	// 	return size
	// }
	// fmt.Println("Size:", readsize(ioreader))

	//fmt.Println("ContentLength:", httpres.ContentLength)
	//fmt.Println("Status:", httpres.StatusCode)
	//fmt.Println("Body:", string(bs))
	fmt.Println("TimeTaken:", time.Since(s).Milliseconds(), "ms")
}
