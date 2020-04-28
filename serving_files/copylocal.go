package main

import (
	"io"
	// "os"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/minnie.jpg", dogpic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `
			<img src="/minnie.jpg">
		`)
}

func dogpic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("minnie.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	//Method 1
	io.Copy(w, f) // copies img to responsewriter

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	// Method 2
	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)  // takes file details as parameter


	//Mehtod 3
	http.ServeFile(w, req, "minnie.jpg")
}