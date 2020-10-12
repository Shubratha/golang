package main

import (
	"io"
	"net/http"
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName string
	LastName string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/f")
}