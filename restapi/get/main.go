package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func practice() {

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
	//b, err := io.Copy(ioutil.Discard, httpRes.Body) //time: 300 ms
	b, err := io.Copy(bodyBuffer2, httpRes.Body)
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
}

var apikey, userid string

func init() {

	fs, err := os.Open("/home/mostain/go/src/practice/credentials.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	bs, err := ioutil.ReadAll(fs)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	slc := strings.SplitN(string(bs), ":", -1)
	apikey, userid = slc[2], slc[3]
}

func main() {

	s := time.Now()

	//GetAllDomains()
	dom, err := domainInfo("fulermela.org")
	fmt.Println(dom, err)

	//customGetRequest("http://mateors.com")

	fmt.Println("TimeTaken:", time.Since(s).Milliseconds(), "ms")
}

func customGetRequest(apiEndpoint string) ([]byte, error) {

	//bytes.NewBuffer(jsonData)
	//var ctx context.Context
	seconds := 600
	//time.Duration(seconds)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*time.Duration(seconds)))
	defer cancel()

	//var url string = "http://mateors.com"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return []byte{}, err
	}

	query := req.URL.Query()
	query.Add("name", "mostain")
	req.URL.RawQuery = query.Encode()

	var client = new(http.Client)
	//client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		//fmt.Println("ERROR:", err.Error())
		return []byte{}, err
	}

	defer resp.Body.Close()
	// if _, err = io.Copy(ioutil.Discard, resp.Body); err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//fmt.Println(err)
		return []byte{}, err
	}

	//fmt.Println("**", req.URL.RequestURI(), req.URL.Host)
	return responseBody, nil
}

func GetAllDomains() error {

	perpage := "30"
	pageno := "1"
	status := "Active"
	apistr := "https://test.httpapi.com/api/domains/search.json?auth-userid=%s&api-key=%s&no-of-records=%s&page-no=%s&status=%s"
	apiEndpoint := fmt.Sprintf(apistr, userid, apikey, perpage, pageno, status)

	//fmt.Println(apiEndpoint)
	byteArray, err := httpRequestAndResponse(apiEndpoint)
	if err != nil {
		//fmt.Println(err.Error())
		return err
	}

	//fmt.Println(byteArray)
	keyValMap := make(map[string]interface{})
	err = parseResponse(byteArray, &keyValMap)
	if err != nil {
		return err
	}

	//domainList := []*domain{}
	var domainList = make([]*domain, 0)
	domainList = map2domain(keyValMap)

	fmt.Println("count:", len(domainList))
	for k, val := range domainList {
		fmt.Println(k, val.domainName, val.expireDate, val.orderid)
	}
	return nil
}

func parseResponse(byteArray []byte, ptrIntrfce interface{}) error {

	//keyValMap := make(map[string]interface{})
	err := byteToJson(byteArray, ptrIntrfce)
	if err != nil {
		//fmt.Println(err.Error())
		return err
	}
	return nil
}

//single domain info
func domainInfo(domainName string) (*domain, error) {

	perpage := "10"
	pageno := "1"
	apistr := "https://test.httpapi.com/api/domains/search.json?auth-userid=%s&api-key=%s&no-of-records=%s&page-no=%s&domain-name=%s"
	apiEndpoint := fmt.Sprintf(apistr, userid, apikey, perpage, pageno, domainName)

	byteArray, err := httpRequestAndResponse(apiEndpoint)
	if err != nil {
		return nil, err
	}
	//fmt.Println(byteArray)
	keyValMap := make(map[string]interface{})
	err = parseResponse(byteArray, &keyValMap)
	if err != nil {
		return nil, err
	}

	//var domainList = make([]*domain, 0)
	var domainInfo *domain
	domainList := map2domain(keyValMap)
	for _, dom := range domainList {
		domainInfo = dom
		break //only reading first element
	}
	return domainInfo, nil
}

func map2domain(keyValMap map[string]interface{}) []*domain {

	domainList := []*domain{}
	//fmt.Println(keyValMap, len(keyValMap))
	for key, val := range keyValMap {

		domainInfo := &domain{}
		if vmap, ok := val.(map[string]interface{}); ok {

			domainInfo.serial = key
			domainInfo.orderid, _ = vmap["orders.orderid"].(string)
			domainInfo.customerid, _ = vmap["entity.customerid"].(string)
			domainInfo.domainName, _ = vmap["entity.description"].(string)
			domainInfo.createDate, _ = vmap["orders.timestamp"].(string)
			domainInfo.expireDate, _ = vmap["orders.endtime"].(string)
			//fmt.Printf("%s => %T %v \n", key, val, vmap["orders.orderid"])
			// for k, v := range vmap {
			// 	fmt.Printf("%s ==> %s = %T %v \n", key, k, v, v)
			// }
			domainList = append(domainList, domainInfo)

		} //else {
		//fmt.Printf("%s => %T %v \n", key, val, val.(string)) //may panic if unable to assert
		//}
	}
	return domainList
}

type domain struct {
	serial     string `json:"serial"`
	orderid    string `json:"orders.orderid"`
	customerid string `json:"entity.customerid"`
	domainName string `json:"entity.description"` //domain name
	createDate string `json:"orders.timestamp"`
	expireDate string `json:"orders.endtime"`
}

// type domainList struct {
// 	RecordsInDB string `json:"recsindb"`
// }

func httpRequestAndResponse(apiEndpoint string) ([]byte, error) {

	httpRes, err := http.Get(apiEndpoint)
	if err != nil {
		return []byte{}, err
	}

	//method-1
	// var buf bytes.Buffer
	// _, err = io.Copy(&buf, httpRes.Body)
	// if err != nil {
	// 	return []byte{}, err
	// }
	// defer httpRes.Body.Close()
	// return buf.Bytes(), nil

	//method-2
	defer httpRes.Body.Close()
	return ioutil.ReadAll(httpRes.Body)

}

func byteToJson(data []byte, i interface{}) error {

	err := json.Unmarshal(data, i)
	if err != nil {
		return err
	}
	return nil
}
