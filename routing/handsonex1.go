package main

import (
	"io"
	"net/http"
)

func root(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>Index</h1>")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "My favourite dog breed is Golden Retriever")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hey, Shubratha!")
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}