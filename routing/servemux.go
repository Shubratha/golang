package main

import (
	"io"
	"net/http"
)

type foo int

func(d foo) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog = woof!")
}

type bar int

func(c bar) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat = meow!")
}

func main() {
	var d foo
	var c bar

	mux := http.NewServeMux()
	mux.Handle("/dog/", d)  // Handle takes handler interface as argument
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux) // mux acts as handler as it has a method attached to it
}