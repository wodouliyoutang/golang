package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/login", login)
	http.ListenAndServe("localhost:8080", nil)
}

func handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "hello worlds")
}
func login(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "login")
}
