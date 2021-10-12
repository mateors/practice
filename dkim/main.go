package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"regexp"
)

func regexVersion() {

	// 	var re = regexp.MustCompile(`(?mis)^-----BEGIN PUBLIC KEY-----\s((.*)\s)*-----END PUBLIC KEY-----[\s]$`)
	// 	var str = `-----BEGIN PUBLIC KEY-----
	// MIIBCgKCAQEAyVdO+Cs1xW2BHYjvvsEF1C4FAVxn/xeMRDMOHoz9amKlezsdPEbW
	// 79bjOXIbuJceblE77UWPS1tU5Wzkd4YdW9Sf/8zREoP6trw4Md8zDzLYqUGhoH32
	// GykEaGXyAA8mKN4fHi3Cjq6d0rEIDZ3C6kq94j2slIhKfiLxgrhZtWKORDwBr+Yc
	// 14sbHN0wgtthGOqMIWeSsDyE7PL//EYEVFxZlISJSoq8rC5ZARRJbqm90FR21ee5
	// 4+H9oxe78PoAknFdQuriaQYCsblpExaLF5CKzg111bdEkqLvCLRyhIoWUDKpoFE7
	// ALmfBdQt8eWG0ZgKYGyz+WRkOH1W5tAbUQIDAQAB
	// -----END PUBLIC KEY-----
	// `

	// 	for i, match := range re.FindAllStringSubmatch(str, -1) {
	// 		//fmt.Println(match, "found at index", i)
	// 		//fmt.Println(i, match)
	// 		for k, v := range match {
	// 			fmt.Println(i, k, v)
	// 		}
	// 	}

	var re = regexp.MustCompile(`-----\s(((.*\s))*)-----`)
	var str = `-----BEGIN PUBLIC KEY-----
MIIBCgKCAQEAyVdO+Cs1xW2BHYjvvsEF1C4FAVxn/xeMRDMOHoz9amKlezsdPEbW
79bjOXIbuJceblE77UWPS1tU5Wzkd4YdW9Sf/8zREoP6trw4Md8zDzLYqUGhoH32
GykEaGXyAA8mKN4fHi3Cjq6d0rEIDZ3C6kq94j2slIhKfiLxgrhZtWKORDwBr+Yc
14sbHN0wgtthGOqMIWeSsDyE7PL//EYEVFxZlISJSoq8rC5ZARRJbqm90FR21ee5
4+H9oxe78PoAknFdQuriaQYCsblpExaLF5CKzg111bdEkqLvCLRyhIoWUDKpoFE7
ALmfBdQt8eWG0ZgKYGyz+WRkOH1W5tAbUQIDAQAB
-----END PUBLIC KEY-----
`
	for i, match := range re.FindAllStringSubmatch(str, -1) {
		for k, v := range match {
			fmt.Println(i, k, v)
		}
	}

	// if len(re.FindStringIndex(str)) > 0 {
	// 	fmt.Println(re.FindString(str), "found at index", re.FindStringIndex(str)[0])
	// }

}

func main() {

	//KeyGenerate("./mostain")

	// file, err := os.Open("mostain.pub")
	// if err != nil {
	// 	//handle error
	// 	fmt.Println(err)
	// 	return
	// }

	// defer file.Close()
	// s := bufio.NewScanner(file)
	// var text string

	// for s.Scan() {
	// 	readLine := s.Text()

	// 	text += strings.ReplaceAll(readLine, "\n", "")
	// 	fmt.Println(readLine)
	// 	// other code what work with parsed line...
	// }

	//method-1 string prefix and suffix
	// file, err := os.Open("mostain.pub")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// defer file.Close()

	// bytes, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// newstr := strings.TrimPrefix(string(bytes), "-----BEGIN PUBLIC KEY-----")
	// onlykey := strings.TrimSuffix(newstr, "-----END PUBLIC KEY-----")

	// oneline := strings.ReplaceAll(onlykey, "\n", "")

	// fmt.Println(onlykey)
	// fmt.Println()
	// fmt.Println(oneline)

	//method-2 regex
	regexVersion()

	// pb, rst := pem.Decode(bytes)
	// pubin, err := x509.ParsePKIXPublicKey(pb.Bytes)
	// pk := pubin.(*rsa.PublicKey)

	//fmt.Println(pb.Type, bytes, rst)
}

func KeyGenerate(path string) {
	reader := rand.Reader
	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}

	publicKey := key.PublicKey

	//saveGobKey("private.key", key)
	savePEMKey(path+".priv", key)

	//saveGobKey("public.key", publicKey)
	savePublicPEMKey(path+".pub", publicKey)
}

func savePEMKey(fileName string, key *rsa.PrivateKey) {

	outFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}

	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}
}

func savePublicPEMKey(fileName string, pubkey rsa.PublicKey) {

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&pubkey),
	}

	pemfile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}
}
