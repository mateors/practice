package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {

	//two dimension array
	var sections = make([][2]int, 10) //declaration & initialization with 0 0 value

	//assign value for the first or index 0
	sections[0][0] = 10
	sections[0][1] = 20

	//assign value for the second or index 1
	sections[1][0] = 30
	sections[1][1] = 40

	//assign value for the third or index 2
	sections[2][0] = 50
	sections[2][1] = 60

	//assign value for the fourth or index 3
	sections[3][0] = 70
	sections[3][1] = 80

	//assign value for the fifth or index 4
	sections[4][0] = 90
	sections[4][1] = 99

	fmt.Println(sections)
	//os.Exit(0)

}

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
