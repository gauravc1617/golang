package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		data, error := ioutil.ReadAll(r.Body)
		if error != nil {
			http.Error(rw, "oops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello %s", data)
	})

	http.HandleFunc("/goodBye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Good bye World")
	})

	http.ListenAndServe(":9090", nil)

}
