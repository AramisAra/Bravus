package main

import (
	"Server/db"
	"context"

	"github.com/gofiber/fiber/v2"
)

// Database Connections

func connect() (*db.PrismaClient, error) {
	client := db.NewClient()
	if err:= client.Prisma.Connect(); err != nil {
		return nil, err
	}

	return client, nil
}
func disconnect(db *db.PrismaClient) error {
	if err := db.Disconnect(); err != nil {
		panic(err)
	}
	return nil
}

// Input Models

type InputUser struct {
	FirstName 	string  `json:"firstName"`
	LastName 		string	`json:"lastName"`
	PhoneNumber string	`json:"phoneNumber"`
	Email 			string	`json:"email"`
	Password 		string	`json:"password"`
}

// Serialized Struct

type ClientResponse struct {
	FirstName 	string	`json:"firstName"`
	LastName 		string	`json:"lastName"`
	PhoneNumber string	`json:"phoneNumber"`
	Email 			string	`json:"email"`
}

// Serializing Functions

func ResponseCreator(dbresp *db.UserModel) ClientResponse {
	return ClientResponse{
		FirstName: dbresp.FirstName, 
		LastName: dbresp.LastName, 
		PhoneNumber:  dbresp.PhoneNumber, 
		Email: dbresp.Email,
		}
}

//CRUD Endpoints

func CreateUser(c *fiber.Ctx) error {
	client, err := connect()
	if err != nil {
		return err
	}
	defer disconnect(client)

	inputHandler := InputUser{}
	if err := c.BodyParser(&inputHandler); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}
	ctx := context.Background()

	createUser, err := client.User.CreateOne(
		db.User.FirstName.Set(inputHandler.FirstName),
		db.User.LastName.Set(inputHandler.LastName),
		db.User.PhoneNumber.Set(inputHandler.PhoneNumber),
		db.User.Email.Set(inputHandler.Email),
		db.User.Password.Set(inputHandler.Password),
	).Exec(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error":"Failed to create user"})
	}

	response := ResponseCreator(createUser)

	return c.Status(fiber.StatusOK).JSON(response)
} 