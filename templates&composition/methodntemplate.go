package main
import (
	"os"
	"text/template"
	"log"
)
type person struct {
	Name string
	Age int
}

func (p person) ReturnSeven() int {
	return 7
}

func (p person) DoubleAge() int {
	return p.Age * 2
}

func (p person) ArgMethod(x int) int {
	return x * 3
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("methodtemp.gohtml"))
}

func main() {
	p := person {
		Name: "Sheldon",
		Age: 35,
	}

	err := tpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln(err)
	}
}