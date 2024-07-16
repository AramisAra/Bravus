package database

import "github.com/surrealdb/surrealdb.go"

func connectDb() {

	db, err := surrealdb.New("ws://localhost:8000/rpc")
	if err != nil {
		panic(err)
	}

	db.Use()
}
