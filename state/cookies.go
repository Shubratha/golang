package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	/* SetCookie takes response and pointer to cookie */
	http.SetCookie(w, &http.Cookie{
		Name: "my-cookie",
		Value: "foo",
	}) // passing composite literal ( Cookie struct)
	fmt.Fprintln(w, "Cookie Set - Check browser")
}

func read(w http.ResponseWriter, req *http.Request) {
	/* Cookie takes string(name of the cookie) and returns pointer to the cookie */
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err .Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "Your cookie: ", c)
}