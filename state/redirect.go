package main

import (
	"html/template"
	"net/http"
	"fmt"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() { 
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("your req method at foo: ", req.Method, "\n\n")
}

func bar(w http.ResponseWriter, req *http.Request) { 
	fmt.Print("your req method at bar: ", req.Method, "\n\n")
	/*
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
	*/
	// or the following can be used to redirect:
	// http.Redirect(w, req, "/", http.StatusSeeOther) // 303  => method is always get

	// http.Redirect(w, req, "/", http.StatusTemporaryRedirect) // 307  => preserves method

	http.Redirect(w, req, "/", http.StatusMovedPermanently) // 301  => remembers new location
	
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Print("your req method at barred: ", req.Method, "\n\n")
	tpl.ExecuteTemplate(w, "index3.gohtml", nil)
}