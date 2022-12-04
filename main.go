package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := &mux.Router{}

	http.ListenAndServe(":8005", r)
}
