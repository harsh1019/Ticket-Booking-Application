package db

import (
	"log"
	"ticketbookingapp/config"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/postgres"
)

func GetDBConnectionString(config *config.EnvConfig, DBMigrator func(db *gorm.DB) error) *gorm.DB {
	uri:= fmt.Sprintf(`
		host=%s user=%s password=%s dbname=%s port=5432 sslmode=%s`,
	    config.DBHOST,
		config.DBUSER,
		config.DBPASSWORD,
		config.DBNAME,
		config.DBSSLMode,
	)


	db,err := gorm.Open(postgres.Open(uri),&gorm.Config{
		
		Logger: logger.Default.LogMode(logger.Info),
	})
    

	if err != nil {
		
		log.Fatalf("Error connecting to the database: %e",err)
	}

	log.Println("Connected to the database")

	if err := DBMigrator(db); err != nil {
		log.Fatalf("Error migrating the database: %e",err)
	}
	return db
}
