package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("hola mundo genial")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "IT'S BROKEN!!!", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Goodbyeeeee")
	})

	http.ListenAndServe(":9090", nil)
}
