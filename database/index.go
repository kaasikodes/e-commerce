package database

import (
	"database/sql"
	"flag"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/database/migrations"
	"github.com/kaasikodes/e-commerce-go/database/seeders"
	"github.com/kaasikodes/e-commerce-go/utils"
)

// TODO: Add database config
// TODO: Add database connection
// TODO: Add database queries to create the needed tables

func MakeConnection() (*sql.DB, bool, error) {
	defer utils.Recover()
	seedDBData := flag.Bool("seed_db_data", false, "This determines whether to seed the database");


	dbUserName := flag.String("db_username", constants.DbUser, "This is the name of the database user");
	dbPassword := flag.String("db_password", constants.DbPassword, "This is the value of the database password");
	dbName := flag.String("db_name", constants.DbName, "This is the value of the database name");
	dbHostName := flag.String("db_host_name", constants.DbHost, "This is the value of the database host name");
	dbNet := flag.String("db_tcp", constants.DbNet, "This is the value of the database network protocol");

	flag.Parse();

	cfg := mysql.Config{
		User: *dbUserName,
		Passwd: *dbPassword,
		Net: *dbNet,
		Addr: *dbHostName,
		DBName: *dbName,
		AllowNativePasswords: true,
		ParseTime: true, //to ensure that the time is parsed correctly from the database
	}

	// http.ListenAndServe(":8000", http.DefaultServeMux)

	// val, err := time.ParseDuration(fmt.Sprint(*msg))
    db, err := sql.Open("mysql", cfg.FormatDSN())
	utils.ErrHandler(err)
	err = db.Ping()
	fmt.Println("Checking database connection .........")
	if(err != nil){
		return nil, *seedDBData,err
	}
	db.SetConnMaxLifetime(constants.DBMaxConnectionLifeTime)
	db.SetMaxOpenConns(constants.DBMaxOpenConnections)
	db.SetMaxIdleConns(constants.DBMaxIdleConnections)
	fmt.Println("Connected to database!")

	// createProductTable(db)
	// fmt.Println(cfg.FormatDSN())
	// fmt.Println(flag.Args())

	return db, *seedDBData,nil

}
func  DropTables(db *sql.DB)  {
	// TODO: Add drop tables queries
	
}


func SetupDB( ) (*sql.DB, error){
	db, seedDBData, err := MakeConnection()
	if err != nil {
		return nil, err
		
	}

	CreateTables(db)
	
	if(seedDBData){
		CreateSeedData(db)
	}
	
	return db, nil

}

func CreateSeedData(db *sql.DB)  {
	err := seeders.AddCountries(db)
	utils.ErrHandler(err,)
	err = seeders.AddStates(db)
	utils.ErrHandler(err)
	err = seeders.AddLGAs(db)
	utils.ErrHandler(err)
	err = seeders.AddAddresses(db)
	utils.ErrHandler(err)
	
}
func  CreateTables(db *sql.DB)  {
	


	err := migrations.CreateCategoryTable(db)
	utils.ErrHandler(err,)
	err = migrations.CreateUserTable(db)
	utils.ErrHandler(err)
	err = migrations.CreateSellerTable(db)
	utils.ErrHandler(err)
	err = migrations.CreateCustomerTable(db)
	utils.ErrHandler(err)
	err = migrations.CreatePasswordResetTokenTable(db)
	utils.ErrHandler(err)
	err = migrations.CreateVerificationTokenTable(db)
	utils.ErrHandler(err)
	err = migrations.CreateProductTable(db)
	utils.ErrHandler(err)
	err = migrations.CreateCartTable(db)
	utils.ErrHandler(err)
	err = migrations.CreateCartItemTable(db)
	utils.ErrHandler(err)
	err = migrations.CreateCountryTable(db)
	utils.ErrHandler(err)
	err = migrations.CreateStateTable(db)
	utils.ErrHandler(err)
	err = migrations.CreateLgaTable(db)
	utils.ErrHandler(err)
	err = migrations.CreateAddressTable(db)
	utils.ErrHandler(err)
	err = migrations.CreateOrderTable(db)
	utils.ErrHandler(err)
	err = migrations.CreateOrderItemTable(db)
	utils.ErrHandler(err)
	err = migrations.CreatePaymentTable(db)
	utils.ErrHandler(err)
	
}







