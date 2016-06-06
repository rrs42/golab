package main

import (
	"database/sql"
	"fmt"
	"strings"
)

/* type TestDB struct */
type TestDB struct {
	db   *sql.DB
	next int
}

/* New() */
func NewTestDB() (*TestDB, error) {
	var err error

	n := new(TestDB)

	n.db, err = sql.Open("sqlite3", "test.db")
	fmt.Println("Init DB")
	n.next = 0
	return n, err
}

func (db *TestDB) Close() {
	db.db.Close()
}

/* (*TestDB) Clear() error */
func (db *TestDB) Clear() error {
	stmt := "DROP TABLE IF EXISTS test"
	_, err := db.db.Exec(stmt)
	return err
}

func (db *TestDB) Init() error {
	err := db.Clear()
	if err != nil {
		return err
	}

	stmt := strings.TrimSpace(`CREATE TABLE IF NOT EXISTS test
			(
				id INTEGER,
				name VARCHAR(80),
				ts INTEGER
			)`)
	_, err = db.db.Exec(stmt)
	return err
}

/* (*TestDB) InsertNew(string, int64) error */
func (db *TestDB) InsertNew(name string, ts int64) error {
	stmt := "INSERT INTO test VALUES($1, $2, $3)"
	_, err := db.db.Exec(stmt, db.next, name, ts)
	if err == nil {
		db.next++
	}
	return err
}

func (db *TestDB) Count() int {
	stmt := "SELECT COUNT(*) FROM test"
	row := db.db.QueryRow(stmt)

	var count int
	err := row.Scan(&count)
	if err != nil {
		panic(err.Error())
	}

	return count
}

type row_iter func(id int, name string, ts int64)

func (db *TestDB) Iterate(fn row_iter) {
	stmt := "SELECT id, name, ts FROM test"
	rows, err := db.db.Query(stmt)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var id int
		var name string
		var ts int64

		err := rows.Scan(&id, &name, &ts)
		if err != nil {
			panic(err.Error())
		}

		fn(id, name, ts)
	}
}
