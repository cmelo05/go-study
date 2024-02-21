package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/oi", Handler)
	http.ListenAndServe(":3000", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)

	w.Write([]byte("Server response oi!"))
}
