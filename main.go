package main
import "fmt"

// 3. struct
type person struct{
	fname string
	lname string
} 

type secretagent struct{
	person
	cankill bool
}


// Interface
type human interface{
	says()
}

// Functions
func (p person) says(){
	fmt.Println(p.fname, `says, "Hello World"`)
}

func (s secretagent) says(){
	fmt.Println(s.person.fname, `says, "Hands up!"`)
}

func speak(h human){
	h.says()
}


func main() {
	x := 7
	fmt.Printf("%T", x)

	//data structure
	// 1. slice
	xi := []int{1,2,3,4}
	fmt.Println(xi)

	// 2. map
	m := map[string]int {
		"me": 20,
		"marley": 5,
	}
	fmt.Println(m)

	// 3. struct
	p1 := person{
		"Shubratha",
		"Raichand",
	}
	fmt.Println(p1)

	p1.says()

	sa1 := secretagent{
		person{
			"Srikant",
			"Tiwari",
		},
		true,
	}
	sa1.says()
	sa1.person.says()

	speak(p1)
	speak(sa1)
}