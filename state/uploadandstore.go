package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
	"net/http"
	"html/template"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func foo(w http.ResponseWriter, req *http.Request) {
	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {
		f, h, err := req.FormFile("q") // catches the file
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		fmt.Println("\nFile: ", f, "\nHeader: ", h, "\nError: ", err)

		bs, err := ioutil.ReadAll(f) // reads file in byte slice
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)

		ds, err := os.Create(filepath.Join("./storedfiles/", h.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer ds.Close()

		_, err = ds.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(w, "index1.gohtml", s)
}