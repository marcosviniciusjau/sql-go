package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	_ = db

	createTable :=
		`CREATE TABLE foo (
	 id integer not null primary key,
	 name text);
	 `

	res, err := db.Exec(createTable)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.RowsAffected())

	insert := `INSERT INTO foo (name) VALUES (1,"foo")`
	res, err = db.Exec(insert)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.RowsAffected())

	type user struct {
		id   int64
		name string
	}
	query := `SELECT * FROM foo WHERE ID = ?`
	var u user

	if err := db.QueryRow(query, 1).Scan(&u.id, &u.name); err != nil {
		panic(err)
	}
	fmt.Println(u)
}
