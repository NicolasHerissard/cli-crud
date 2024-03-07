package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

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

func InsertStatement(db *sql.DB, last_name *string, first_name *string) string {
	res, err := db.Exec("INSERT INTO employee (last_name, first_name) VALUES (?, ?)", last_name, first_name)
	if err != nil {
		panic(err.Error())
	}

	rowCount, err := res.RowsAffected()

	if err != nil {
		panic(err.Error())
	}

	return "Nombre de lignes insérées : " + strconv.Itoa(int(rowCount))
}

func DeleteStatement(db *sql.DB, id *int) string {
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

	return "Nombre de lignes supprimées : " + strconv.Itoa(int(rowCount))
}

/*
 */
func UpdateStatement(db *sql.DB, lastname *string, firstname *string, id *int) string {
	res, err := db.Exec("UPDATE employee SET last_name = ?, first_name = ? WHERE id = ?", lastname, firstname, id)
	if err != nil {
		panic(err.Error())
	}

	rowCount, err := res.RowsAffected()

	if err != nil {
		panic(err.Error())
	}

	return "Nombre de lignes affectées : " + strconv.Itoa(int(rowCount))
}
