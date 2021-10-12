package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

	file, err := os.Open("mostain.pub")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	newstr := strings.TrimPrefix(string(bytes), "-----BEGIN PUBLIC KEY-----")
	onlykey := strings.TrimSuffix(newstr, "-----END PUBLIC KEY-----")

	oneline := strings.ReplaceAll(onlykey, "\n", "")

	fmt.Println(onlykey)
	fmt.Println()
	fmt.Println(oneline)

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
