package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func init() {

	size, err := getContentLength("https://automan.biz/resources/images/contact.png")
	fmt.Println(size, err)

	fmt.Println(fmt.Sprintf("%.2f", float64(size/20)))

	eachSectionByteSize := size / 10

	var sections = make([][2]int, 10)

	for i := range sections {

		if i == 0 {
			sections[i][0] = 0
			sections[i][1] = int(eachSectionByteSize)

		} else if i == 10-1 {

			first := sections[i-1][1] + 1
			sections[i][0] = first
			sections[i][1] = first + int(eachSectionByteSize) - 1

		} else {
			first := sections[i-1][1] + 1
			sections[i][0] = first
			sections[i][1] = first + int(eachSectionByteSize)

		}

		fmt.Println(i, sections[i][0], sections[i][1])

		// if i == 0 {
		// 	// starting byte of first section
		// 	sections[i][0] = 0
		// } else {
		// 	// starting byte of other sections
		// 	sections[i][0] = sections[i-1][1] + 1
		// }

		// if i < d.TotalSections-1 {
		// 	// ending byte of other sections
		// 	sections[i][1] = sections[i][0] + eachSize
		// } else {
		// 	// ending byte of other sections
		// 	sections[i][1] = size - 1
		// }

	}

	os.Exit(0)

}

func twoDimensionalArray() {

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

}

func getContentLength(fileUrl string) (uint64, error) {

	req, err := http.NewRequest("HEAD", fileUrl, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("User-Agent", "MateorsDownloader")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	//verifying is really body empty or not
	// bs, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(bs, err)
	length := resp.Header.Get("Content-Length")
	//fmt.Println(length)
	return strconv.ParseUint(length, 10, 64)
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
