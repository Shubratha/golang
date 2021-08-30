package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"net/http"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	imgname := "beach"
	c := getCookie(w, req, imgname)
	err := tpl.ExecuteTemplate(w, "index.gohtml", c.Value)
	if err != nil {
		fmt.Println(err)
		return 
	}
}

func getCookie(w http.ResponseWriter, req *http.Request, imgname string) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name: "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	i, err := req.Cookie("img")
	if err != nil {
		i = &http.Cookie{
			Name: "img",
			Value: imgname,
		}
		http.SetCookie(w, i)
	} else {
		if ! strings.Contains(i.Value, imgname) {
			i = &http.Cookie{
				Name: "img",
				Value: i.Value + "|" + imgname,
			}
			http.SetCookie(w, i)
		}
	}

	return c
}