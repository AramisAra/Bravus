package database

import "github.com/surrealdb/surrealdb.go"

func ConnectDb() surrealdb.DB {
	db, err := surrealdb.New("ws://0.0.0.0:80/rpc")
	if err != nil {
		panic(err)
	}

	db.Signin(map[string]interface{}{
		"user": "Dev",
		"pass": "Ara",
	})

	db.Use("Build", "Dev")

	return *db
}
