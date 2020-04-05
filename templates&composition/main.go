package main
import (
	"os"
	"text/template"
	"log"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"C1", "Course 1", "4"},
				course{"C2", "Course 2", "4"},
				course{"C3", "Course 3", "3"},

			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"C4", "Course 4", "4"},
				course{"C5", "Course 5", "3"},
				course{"C6", "Course 6", "4"},

			},
		},
	}

	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}