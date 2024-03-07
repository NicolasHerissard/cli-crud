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

	update_stmt := flag.NewFlagSet("update", flag.ContinueOnError)
	id_update := update_stmt.Int("id", 1, "")
	updateLastname := update_stmt.String("lastname", "", "")
	updateFirstname := update_stmt.String("firstname", "", "")

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
			fmt.Println(InsertStatement(db, lastnameValue, firstnameValue))
		} else {
			fmt.Println("Argument vide")
		}

	case "delete":
		delete_stmt.Parse(os.Args[2:])
		fmt.Println(DeleteStatement(db, id_delete))

	case "update":
		update_stmt.Parse(os.Args[2:])
		fmt.Println(UpdateStatement(db, updateLastname, updateFirstname, id_update))
	}
}