package main

import (
	"github.com/kaasikodes/e-commerce-go/database"
	"github.com/kaasikodes/e-commerce-go/server"
	"github.com/kaasikodes/e-commerce-go/utils"
)

func main() {
	defer utils.Recover()
	// TODO: Add a flag to specify the port
	// TODO: Define errors constants
	// go get github.com/database/sql
	// https://golangbot.com/connect-create-db-mysql/
	// TODO: Create services that will be used in the controllers
	// TODO: Implement a try/catch (panic, recover) to ensure the server doesn't crash, and test to ensure this
	// TODO: Create auth routes
	// TODO: Create user routes
	// TODO: Create product routes
	// TODO: Create cart routes
	// TODO: Create order routes
	// TODO: Create payment routes




	
	// Connect to database
	db, err := database.SetupDB()
	utils.ErrHandler(err)
	defer db.Close()

	server.NewApiServer(db, ":8000").Start()

	
	




}


