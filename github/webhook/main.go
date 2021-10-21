package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/github", github)
	log.Println("listening on port: 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Welcome coder.cf")

}

func github(w http.ResponseWriter, r *http.Request) {

	protocol := r.Proto
	log.Println(protocol, r.Method, r.Host)

	log.Println("-----------HEADER----------------")
	for key, vals := range r.Header {
		log.Println(key, vals)

	}

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	log.Println("-----------JSON BODY----------------")

	rMap := make(map[string]interface{})
	json := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err = json.Decode(&rMap)
	if err != nil {
		log.Println(err)
	}

	for key, vals := range rMap {

		log.Println(key, vals)

	}

	log.Println("Count:", len(r.Form), len(rMap))

	fmt.Fprintln(w, "github webhook 1.0.1")

}
