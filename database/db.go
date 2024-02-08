package database

import (
	"fmt"
	"os"

	// "github-zaki-learning-golang/database/migration"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func DBConnection() *gorm.DB {
	
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)
	// dsn := "host=localhost user=postgres password=root dbname=pustakaapigolang port=5432 sslmode=disable"
	DBPostgres ,err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}

	// migration.DatabaseMigration(DBPostgres)
	
	fmt.Println(" to database")
	return DBPostgres
}