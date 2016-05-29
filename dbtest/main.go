package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/fatih/color"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	red := color.New(color.FgBlue).SprintFunc()

	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	createStmt := strings.TrimSpace(`
    CREATE TABLE IF NOT EXISTS test ( id INTEGER, name VARCHAR(80), ts INTEGER )
    `)

	fmt.Print(red(createStmt))

	_, err = db.Exec(createStmt)
	if err != nil {
		log.Printf("%q: %s", err, createStmt)
		return
	}

	countStmt := "SELECT COUNT(*) FROM test"
	row := db.QueryRow(countStmt)

	var dbCount int
	err = row.Scan(&dbCount)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Rows in db: %d\n", dbCount)

	insertStmt, err := db.Prepare("INSERT INTO test VALUES(?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer insertStmt.Close()

	now := time.Now()
	_, err = insertStmt.Exec(10, "Bob", now.Unix())
	if err != nil {
		panic(err.Error())
	}

	qryStmt, err := db.Prepare("SELECT id, name, ts FROM test")
	if err != nil {
		panic(err.Error())
	}
	defer qryStmt.Close()

	rows, err := qryStmt.Query()
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var id int
		var name string
		var ts int64

		err = rows.Scan(&id, &name, &ts)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(id, name, time.Unix(ts, 0))
	}
}
