package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/fatih/color"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func connectDB(driver string, args string) *sql.DB {
	db, err := sql.Open(driver, args)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func dropTable(db *sql.DB) error {
	stmt := `DROP TABLE IF EXISTS test`

	_, err := db.Exec(stmt)
	return err
}

func createTable(db *sql.DB) error {
	stmt := strings.TrimSpace(`
		CREATE TABLE IF NOT EXISTS
		test
		(
			id INTEGER,
			name VARCHAR(80),
			ts INTEGER
		)`)

	_, err := db.Exec(stmt)
	return err
}

func countRows(db *sql.DB) int {
	stmt := "SELECT COUNT(*) FROM test"
	row := db.QueryRow(stmt)

	var count int
	err := row.Scan(&count)
	if err != nil {
		panic(err.Error())
	}

	return count
}

func insertRows(db *sql.DB, id int, name string, ts int64) {
	stmt, err := db.Prepare("INSERT INTO test VALUES( $1, $2, $3)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, name, ts)
	if err != nil {
		panic(err.Error())
	}
}

func queryRows(db *sql.DB) {
	stmt := "SELECT id, name,ts FROM test"

	rows, err := db.Query(stmt)
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

func main() {
	red := color.New(color.FgRed).SprintFunc()

	now := time.Now()

	fmt.Println("Testing testdb")
	db1, err := NewTestDB()
	if err != nil {
		panic(err.Error())
	}
	defer db1.Close()
	db1.Init()
	db1.InsertNew("Jeb", now.Unix())
	db1.InsertNew("Bill", now.Unix())
	fmt.Println(db1.Count())

	i := func(id int, name string, ts int64) {
		type Entry struct {
			Id   int
			Name string
			Ts   int64
		}

		entry := Entry{id, name, ts}

		if entry.Id == 1 {
			entry.Name = red(entry.Name)
		}

		tmpl, err := template.New("test").Parse("ID: {{.Id}} Name: {{.Name}} Timestamp: {{.Ts}}\n")
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(os.Stdout, entry)
		// fmt.Println(id, name, ts)
	}

	db1.Iterate(i)
}
