package database

import (
	"fmt"
	"log"
	"moddleware/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "admin"
	port     = "5432"
	dbname   = "db_go_sql"
	db       *gorm.DB
	err      error
)

func StartDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	// buat tabel baru
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "db_go_sql.",
			SingularTable: false,
		},
	})
	if err != nil {
		log.Fatal("error connectng to db :", err)
	}

	log.Println("successfully connected to")
	db.Debug().AutoMigrate(model.Usergo{}, model.Product{})

}

func GetDB() *gorm.DB {
	return db
}
