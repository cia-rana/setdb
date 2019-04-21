package setbd

import (
	"database/sql"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type DBWithSQLite3 struct {
	db *sql.DB
}

func OpenWithSQLite3(dataSourceName string) (*DBWithSQLite3, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(`CREATE TABLE T(V text PRIMARY KEY)`); err != nil && !checkTableAlreadyExists(err) {
		return nil, err
	}

	return &DBWithSQLite3{db: db}, nil
}

func (db *DBWithSQLite3) Close() error {
	return db.db.Close()
}

func (db *DBWithSQLite3) Contain(v string) (bool, error) {
	row := db.db.QueryRow(`SELECT EXISTS (SELECT * FROM T WHERE V = ?)`, v)

	var n int
	err := row.Scan(&n)
	if err != nil {
		return false, err
	}

	return n == 1, nil
}

func (db *DBWithSQLite3) Insert(v string) error {
	_, err := db.db.Exec(`INSERT INTO T VALUES(?)`, v)
	if err != nil && !checkUniqueConstraintFailed(err) {
		return err
	}

	return nil
}

func (db *DBWithSQLite3) Erase(v string) error {
	_, err := db.db.Exec(`DELETE FROM T WHERE V = ?`, v)
	if err != nil {
		return err
	}

	return nil
}

func checkTableAlreadyExists(err error) bool {
	return strings.HasSuffix(err.Error(), "already exists")
}

func checkUniqueConstraintFailed(err error) bool {
	return strings.HasPrefix(err.Error(), "UNIQUE constraint faild")
}
