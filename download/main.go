package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	downloadUrl := "https://automan.biz/resources/images/contact.png"

	resp, err := http.Get(downloadUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// for key, val := range resp.Header {
	// 	fmt.Println(key, val)
	// }
	// os.Exit(0)

	pf, err := os.OpenFile("./download/code.png", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}

	written, err := io.Copy(pf, resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// var bs = make([]byte, 100)
	// n, err := pf.Read(bs)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Println("bytes written:", written)
}
