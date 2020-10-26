package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

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
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo (w http.ResponseWriter, req *http.Request) {
	bs := make([]byte, req.ContentLength)  // make byte slice; length & capacity is set to request content length
	req.Body.Read(bs)
	body := string(bs)

	err := tpl.ExecuteTemplate(w, "index2.gohtml", body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}