package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy.. <3")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "kitty cat.. meow")
}

func main() {
	http.HandleFunc("/dog/", d)   // HandleFunc takes handler func
	http.HandleFunc("/cat", c) // attaches to DefaultServeMux
	http.ListenAndServe(":8080", nil) //handler is nil	
}