package database

import "github.com/surrealdb/surrealdb.go"

func ConnectDb() surrealdb.DB {
	db, err := surrealdb.New("ws://172.24.195.132:3000/rpc")
	if err != nil {
		panic(err)
	}

	db.Signin(map[string]interface{}{
		"user": "ara",
		"pass": "ara2454",
	})

	db.Use("test", "test")

	return *db
}
