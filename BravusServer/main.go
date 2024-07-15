package main

import (
	"github.com/surrealdb/surrealdb.go"
)

type User struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func main() {
	db, err := surrealdb.New("ws://0.0.0.0:8000/rpc")
	if err != nil {
		panic(err)
	}

	if _, err = db.Signin(map[string]interface{}{
		"user": "root",
		"pass": "root",
	}); err != nil {
		panic(err)
	}

	if _, err = db.Use("test", "test"); err != nil {
		panic(err)
	}

	// Create user
	user := User{
		Name:    "John",
		Surname: "Doe",
	}

	// Insert user
	data, err := db.Create("user", user)
	if err != nil {
		panic(err)
	}

	// Unmarshal data
	createdUser := make([]User, 1)
	err = surrealdb.Unmarshal(data, &createdUser)
	if err != nil {
		panic(err)
	}

	// Get user by ID
	data, err = db.Select(createdUser[0].ID)
	if err != nil {
		panic(err)
	}

	// Unmarshal data
	selectedUser := new(User)
	err = surrealdb.Unmarshal(data, &selectedUser)
	if err != nil {
		panic(err)
	}

	// Change part/parts of user
	changes := map[string]string{"name": "Jane"}
	// Update user
	if _, err = db.Update(selectedUser.ID, changes); err != nil {
		panic(err)
	}

	if _, err = db.Query("SELECT * FROM $record", map[string]interface{}{
		"record": createdUser[0].ID,
	}); err != nil {
		panic(err)
	}

	// Delete user by ID
	if _, err = db.Delete(selectedUser.ID); err != nil {
		panic(err)
	}

}
