package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type cwriter struct {
	total uint64
}

func (c *cwriter) Write(p []byte) (n int, err error) {
	c.total += uint64(len(p)) //32 kb at a time
	progress := float64(c.total) / (1024 * 1024)
	fmt.Printf("\rDownloading %f MB..", progress)
	return len(p), nil
}

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

	reader := io.TeeReader(resp.Body, &cwriter{}) //showing progressbar
	written, err := io.Copy(pf, reader)
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
