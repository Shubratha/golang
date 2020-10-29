package main

import (
	"net/http"
	"log"
	"io"
	"strconv"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/expire", expire)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("counter")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name: "counter",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)
	http.SetCookie(w, cookie)
	io.WriteString(w, cookie.Value)
}

func expire(w http.ResponseWriter, req *http.Request) {
	/* clears a cookie */
	cookie, err := req.Cookie("counter")
	if err != nil {
		log.Fatalln(err)
		// http.SetCookie(w, &http.Cookie{
		// 	Name: "counter",
		// 	Value: "1",
		// }) // doesnt work if cookie not found?!
	}
	cookie.MaxAge = -1  //0 or negative value
	http.SetCookie(w, cookie)
}
