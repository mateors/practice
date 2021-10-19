package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", indxHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func indxHandler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	fmt.Fprintln(w, "Hello web", path)

}
