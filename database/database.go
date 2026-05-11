package database

import(
	"fmt"
	"log"

	"github.com/ranggawaridat/belajar-golang/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=belajar port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database")
	}

	fmt.Println("Database connected")

	err =database.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatal(err)
	} 
	DB = database
	
}
