package main

import (
	"github.com/gorilla/mux"
	"net/http"
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
}
