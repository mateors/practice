package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func init() {

	//r.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", c[0], c[1]))

	//https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Range
	//Request: Range -> bytes=100-200
	//Content-Range: bytes 200-1000/67589
	//vrange := fmt.Sprintf("bytes=%v-%v", 0, 100)
	//fmt.Println(vrange)

	// clength, err := getContentLength("https://drive.google.com/uc?export=download&confirm=hF9Z&id=1FFM-WG5JSkg0a41ZTEH4oinGJQz5sKsm")
	// fmt.Println(err, clength)
	// os.Exit(0)

	downloadmedia("https://doc-10-bk-docs.googleusercontent.com/docs/securesc/ug6gumdndi1jq9dh2rv9sn6k67fcvu0k/0lk3ilt11pnomqpku42h20as680u38o1/1633836975000/04552009210644688537/06512679350586085468Z/1FFM-WG5JSkg0a41ZTEH4oinGJQz5sKsm?e=download&nonce=hl4t1jfcbba0u&user=06512679350586085468Z&hash=9fc30vc21volfs36n1asepurh0uuppro")
	os.Exit(0)

	size, err := getContentLength("https://automan.biz/resources/images/contact.png")
	fmt.Println(size, err)

	var numbverOfSection int = 10
	//fmt.Println(fmt.Sprintf("%.2f", float64(size/numbverOfSection)))

	eachSectionByteSize := size / uint64(numbverOfSection)

	var sections = make([][2]int, numbverOfSection)

	for i := range sections {

		if i == 0 {
			sections[i][0] = 0
			sections[i][1] = int(eachSectionByteSize)

		} else if i == numbverOfSection-1 {

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

	//HEAD
	req, err := http.NewRequest("GET", fileUrl, nil)
	if err != nil {
		return 0, err
	}
	//req.Header.Set("User-Agent", "MateorsDownloader")
	//req.Header.Set("NIC", "511=uy4ufW5-yIzQPEa_Wm3aa9pgfOp9v_m2yPx-jxn9E7lM5fZ-qRoiVWZaLgQkShxYhhd-rFdSg84lOeeJ9pOk--m7yGK3-NpKynzVZpuBK-t8Xn50vvnydyLICg5eIv8E2zZRaOrGIzl62vSq__QEWmJhM5JuF4vUfglDYnwxJks")

	req.AddCookie(&http.Cookie{Name: "NID", Value: "511=JudNtY18WynKObaxmdN1FPPiD-5E7dwmZQt5WwticA9QTI9iVRDzQj1yMFCQla-MNJqy8VBLBdqPFxgVIBTYI8G6loeqcwjv4UX5-elIvEfnHZPshP5BshRlqZ97tngUimLXBJyS2i7CilxGl5t2pRye9Vl4B2H1oqIYe8NLeLk"})
	req.AddCookie(&http.Cookie{Name: "download_warning_13058876669334088843_1FFM-WG5JSkg0a41ZTEH4oinGJQz5sKsm", Value: "hF9Z"})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	for key, val := range resp.Header {
		fmt.Println(">>", key, val)
	}

	fmt.Println("Status-Code:", resp.StatusCode)
	//verifying is really body empty or not
	bs, err := ioutil.ReadAll(resp.Body)
	fmt.Println("***", bs, err)

	length := resp.Header.Get("Content-Length")
	fmt.Println("length:", length)
	if length == "" {
		return 0, errors.New("no response from content-length header")
	}
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

	// var bs = make([]byte, 100)
	// n, err := pf.Read(bs)
	// if err != nil {
	// 	fmt.Println(err)
	// }

}

func downloadmedia(downloadUrl string) {

	//downloadUrl := "https://automan.biz/resources/images/contact.png"
	resp, err := http.Get(downloadUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	for key, val := range resp.Header {
		fmt.Println("##", key, val)
	}
	fmt.Println("StatusCode:", resp.StatusCode)

	pf, err := os.OpenFile("./download/code", os.O_CREATE|os.O_WRONLY, 0600)
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
	fmt.Println("bytes written:", written)
}
