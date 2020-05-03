package store

import (
	"github.com/gocraft/dbr/v2"
	"log"
)

type DB *dbr.Session

func NewDB() DB {
	conn, err := dbr.Open("postgres",
		"postgres://hoorzy:sahar67@localhost:5432/todoapp?sslmode=disable",
		nil,
	)
	if err != nil {
		log.Println("error connection db", err.Error())
		panic(err)
	}
	session := conn.NewSession(nil)

	return session
}
