package main
import (
	"fmt"
	"net"
	"log"
	"bufio"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request(conn)

	// write request
	request(conn)	
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			u := strings.Fields(ln)[1]
			m := strings.Fields(ln)[0]
			fmt.Println("***METHOD", m)
			fmt.Println("***URL", u)
		}
		if ln == "" {
			break
		}
		i++
	}
}