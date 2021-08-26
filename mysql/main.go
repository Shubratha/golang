package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
)

var db *sql.DB
var err error

const (
	username = "root"
	password = "password"
	hostname = "localhost:3306"
	dbname   = "mypro"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}
func main() {
	db, err := sql.Open("mysql", dsn(dbname))//; if err != nil {
	//	log.Println(err)
	//	panic(err)
	//}
	check(err)
	defer db.Close()

	err= db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", dele)
	http.HandleFunc("/drop", drop)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "Successfully completed.")
	check(err)
}

func check(err error) {
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		fmt.Println(err)
	}
}

func create_conn() *sql.DB {
	db, err := sql.Open("mysql", dsn(dbname))
	check(err)
	return db
}

func create(w http.ResponseWriter, req *http.Request) {
	//stmt, err := db.Prepare("create TABLE people(name varchar(20));")
	db = create_conn()
	stmt, err := db.Query(`create TABLE people(name varchar(20));`)
	if err != nil {
		fmt.Println(err)
	}
	check(err)
	defer stmt.Close()
	fmt.Println(stmt, err)
	//defer func(stmt *sql.Stmt) {
	//	err := stmt.Close()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}(stmt)
	//r, err := stmt.Exec()
	//check(err)
	//
	//n, err := r.RowsAffected()
	//check(err)
	//fprintln, err := fmt.Fprintln(w, "Created table People", n)
	//check(err)
	//fmt.Println(fprintln)
}

func insert(w http.ResponseWriter, req *http.Request) {
	db = create_conn()
	stmt, err := db.Prepare(`insert into people values ("Harry");`)
	check(err)

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Println(w, "Row inserted", n)
}

func read(w http.ResponseWriter, req *http.Request) {
	db = create_conn()
	stmt, err := db.Query(`select * from people;`)
	check(err)
	defer stmt.Close()
	var name string
	for stmt.Next() {
		err = stmt.Scan(&name)
		check(err)
		fmt.Println(w, "Retrieved record", name)
	}
}

func update(w http.ResponseWriter, req *http.Request) {
	db = create_conn()
	stmt, err := db.Prepare(`Update people set name ="Potter" where name="Harry";`)
	check(err)

	r, err := stmt.Exec()
	check(err)
	n, err := r.RowsAffected()
	check(err)
	fmt.Println(w, "Row uodated", n)
}

func dele(w http.ResponseWriter, req *http.Request) {
	db = create_conn()
	stmt, err := db.Prepare(`delete from people where name="Potter";`)
	check(err)
	r, err := stmt.Exec()
	check(err)
	n, err := r.RowsAffected()
	check(err)
	fmt.Println(w, "Row deleted", n)
}

func drop(w http.ResponseWriter, req *http.Request) {
	db = create_conn()
	stmt, err := db.Prepare(`drop table people;`)
	check(err)
	defer stmt.Close()
	_, err = stmt.Exec()
	check(err)
	fmt.Println(w, "Dropped table people")
}