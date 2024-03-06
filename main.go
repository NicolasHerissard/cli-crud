package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "admin:admin@/franchise")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	select_stmt := flag.NewFlagSet("select", flag.ContinueOnError)
	id := select_stmt.Int("id", 1, "identifiant")

	insert_stmt := flag.NewFlagSet("insert", flag.ContinueOnError)
	lastnameValue := insert_stmt.String("lastname", "", "")
	firstnameValue := insert_stmt.String("firstname", "", "")

	delete_stmt := flag.NewFlagSet("delete", flag.ContinueOnError)
	id_delete := delete_stmt.Int("id", 1, "")

	switch os.Args[1] {
	case "select":
		select_stmt.Parse(os.Args[2:])
		if os.Args[2] == "all" {
			SelectStatement(db)
		} else {
			fmt.Println(SelectOneStatement(db, id))
		}

	case "insert":
		insert_stmt.Parse(os.Args[2:])
		if *lastnameValue != "" && *firstnameValue != "" {
			InsertStatement(db, lastnameValue, firstnameValue)
		} else {
			fmt.Println("Argument vide")
		}

	case "delete":
		delete_stmt.Parse(os.Args[2:])
		DeleteStatement(db, id_delete)
	}
}

func SelectStatement(db *sql.DB) {

	var (
		last_name  string
		first_name string
	)

	stmt, err := db.Query("SELECT last_name, first_name FROM employee")
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	for stmt.Next() {
		stmt.Scan(&last_name, &first_name)
		fmt.Println(last_name + " " + first_name)
	}
}

func SelectOneStatement(db *sql.DB, id *int) (string, string) {

	var (
		last_name  string
		first_name string
	)

	err := db.QueryRow("SELECT last_name, first_name FROM employee WHERE id = ?", id).Scan(&last_name, &first_name)
	if err != nil {
		panic(err.Error())
	} else if err == sql.ErrNoRows {
		panic("L'id n'existe pas")
	}

	return last_name, first_name
}

func InsertStatement(db *sql.DB, last_name *string, first_name *string) {
	res, err := db.Exec("INSERT INTO employee (last_name, first_name) VALUES (?, ?)", last_name, first_name)
	if err != nil {
		panic(err.Error())
	}

	rowCount, err := res.RowsAffected()

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Nombre de lignes affectées : %d", rowCount)
}

func DeleteStatement(db *sql.DB, id *int) {
	res, err := db.Exec("DELETE FROM employee WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	} else if err == sql.ErrNoRows {
		panic("L'id n'existe pas")
	}

	rowCount, err := res.RowsAffected()

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Nombre de lignes supprimées : %d", rowCount)
}
