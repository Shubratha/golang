package main
import (
	"net/http"
	"fmt"
)

type hey int

func (m hey) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Whatever..!") // prints in the browser
	fmt.Println("Whatever..!") //prints in console

	w.Header().Set("hello", "world") //set custom response header cotent key
}

func main() {
	var d hey
	http.ListenAndServe(":8080", d)
}