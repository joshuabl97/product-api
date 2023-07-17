package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		// when the handler is envoked this log message appears on the server
		// however, it is not revealed to the client as it is not attached to the http.ResponseWriter
		log.Println("Hello World")
		// reading the data from the request body (returned as []byte)
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Failed to read request body", http.StatusBadRequest)
			rw.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(rw, "Failed to read request body %v\n", err)
			return
		}

		// prints back to the http.ResponseWrite (returning the string to the user)
		fmt.Fprintf(rw, "Hello %s\n", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	http.ListenAndServe(":8080", nil)
}
