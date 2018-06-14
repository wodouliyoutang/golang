package main

import (
	"net/http"
	"log"
	)


func main() {
	s := &http.Server{
	Addr:           ":8099",
	//Handler:        myHandler,
	ReadTimeout:    10 ,
	WriteTimeout:   10 ,
	MaxHeaderBytes: 1 << 20,
}
	log.Fatal(s.ListenAndServe())
}