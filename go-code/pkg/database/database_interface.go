package database

import (
    "database/sql"
)

type DBInterface interface {
    Exec(query string, args ...interface{}) (sql.Result, error)
    Query(query string, args ...interface{}) (*sql.Rows, error)
    QueryRow(query string, args ...interface{}) *sql.Row
    Prepare(query string) (*sql.Stmt, error)
}