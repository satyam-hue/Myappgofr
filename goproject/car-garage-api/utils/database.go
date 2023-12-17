package utils
import "fmt"



import (
	"car-garage-api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=1234")
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.Car{})

	DB = db
	DB = db
    fmt.Println("Database connected and tables migrated successfully.")
}
