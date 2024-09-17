package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/123", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("Hello 123"))
	})
	r.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("Hello World"))
	})
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
