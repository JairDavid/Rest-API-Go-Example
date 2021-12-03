package connection

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	"github.com/rest-api-market/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
	lock     sync.Once
)

func Driver(name string) {

	switch strings.ToLower(name) {
	case "postgres":
		getDbInstance()
	}
}

func getDbInstance() {
	lock.Do(func() {

		environment := godotenv.Load()
		if environment != nil {
			log.Fatal(environment)
		}

		uri := os.Getenv("DATABASE_URI_DEV")
		db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})

		database = db
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("the database has been connected successfuly")
		}

	})
}

func GetConnection() *gorm.DB {
	return database
}

//database migrations
func StartMigrations() {
	db := GetConnection()
	db.AutoMigrate(&model.Category{}, &model.Product{}, &model.Customer{}, &model.CustomerProduct{})

	//Manual Mode
	// db.AutoMigrate(&model.Product{}).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")
	// db.AutoMigrate(&model.CustomerProduct{}).AddForeignKey("customer_id", "customers(id)", "RESTRICT", "RESTRICT")
	// db.AutoMigrate(&model.CustomerProduct{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")
}
