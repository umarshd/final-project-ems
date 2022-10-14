package config

import (
	"final-project-ems/entity"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
)

var DB *gorm.DB
var err error

type (
	dsn struct {
		Host     string
		User     string
		Password string
		Dbname   string
		Port     string
	}
)

func Database() {

	dsn := dsn{
		Host:     os.Getenv("PGHOST"),
		User:     os.Getenv("PGUSER"),
		Password: os.Getenv("PGPASSWORD"),
		Dbname:   os.Getenv("PGDATABASE"),
		Port:     os.Getenv("PGPORT"),
	}

	db_url := "host=" + dsn.Host + " user=" + dsn.User + " password=" + dsn.Password + " dbname=" + dsn.Dbname + " port=" + dsn.Port
	DB, err = gorm.Open(postgres.Open(db_url), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	log.Println("Database Connected")
}

func Migrate() {
	DB.AutoMigrate(
		&entity.Admin{},
		&entity.ProductCategory{},
		entity.UserProfile{},
		entity.User{},
		entity.UserAddress{},
		entity.Product{},
		entity.Cart{},
		entity.CartItem{},
		entity.Checkout{},
	)
}
