package model

import "database/sql"

type User struct {
	Id	int
	Firstname string
	Lastname string
	Age sql.NullInt16
}