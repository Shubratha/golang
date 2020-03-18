package main
import "fmt"
import "os"
import "log"
import "io"
import "strings"

// Template

func main() {
	// name := "hey" 
	name := os.Args[1]
	temp := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`)		

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	// fmt.Println(temp)

	io.Copy(nf, strings.NewReader(temp))
}

