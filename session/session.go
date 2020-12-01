package main
import (
	"fmt"
	"net/http"
	"github.com/satori/go.uuid"
	"html/template"
)

type user struct {
	UserName string
	First string
	Last string
}

var tpl *template.Template
var dbUsers = map[string]user{}  // userId, user
var dbSessions = map[string]string{}  // sessionId, userID
// or var dbSessions = make(map[string]string)

func init() {
	tpl =  template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session-id")
	if err != nil {
		sId, err1 := uuid.NewV4()
		if err1 != nil {
			fmt.Printf("err: ", err1)
		}
		c = &http.Cookie{
			Name: "session-id",
			Value: sId.String(),
		}
		http.SetCookie(w, c)
	}
	var u user
	if un, ok := dbSessions[c.Value]; ok { // ok return true or false based on if key exists
		u = dbUsers[un]
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session-id")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}