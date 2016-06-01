package testdb

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
func New() (*TestDB, error) {
	var n TestDB
	var err error

	n.db, err = sql.Open("sqlite3", "test.db")
	fmt.Println("Init DB")
	n.next = 0
	return &n, err
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
