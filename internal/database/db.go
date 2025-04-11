package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/DevAthhh/url-shortener/internal/lib/generateAlias"
	"github.com/DevAthhh/url-shortener/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func (d *Database) SaveURL(url string, size int) (string, error) {
	alias := generateAlias.GenerateStr(size)

	urlToDB := models.Url{
		Alias: alias,
		Root:  url,
	}

	if res := d.db.Create(&urlToDB); res.Error != nil {
		return "", res.Error
	}

	return alias, nil
}

func (d *Database) GetUrl(alias string) (string, error) {
	if alias == "" {
		return "", errors.New("alias cannot be empty")
	}
	var url models.Url
	if res := d.db.First(&url, "alias = ?", alias); res.Error != nil {
		return "", res.Error
	}
	return url.Root, nil
}

func LoadDatabase() *Database {
	db, err := loadDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := syncDB(db); err != nil {
		log.Fatal(err)
	}

	return &Database{
		db: db,
	}
}

func loadDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PWD"),
		os.Getenv("DATABASE_NAME"),
	)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func syncDB(db *gorm.DB) error {
	return db.AutoMigrate(&models.Url{})
}
