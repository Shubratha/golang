package main

import (
	"io"
	"log"
	"net/http"
	"html/template"
)

func root(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>Index</h1><br>Go to /dog/")
}

func dog(w http.ResponseWriter, req *http.Request) {
	// io.WriteString(w, "My favourite dog breed is Golden Retriever")
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	data := struct {
		Breed string
		Origin string
	} {
		"Golden Retriever",
		"UK, Scotland, England",
	}

	err = tpl.ExecuteTemplate(w, "dog.gohtml", data)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(root))
	http.Handle("/dog/", http.HandlerFunc(dog))

	http.ListenAndServe(":8080", nil)
}