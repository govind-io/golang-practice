package db

import (
	appUtil "api/src/utils"
	"errors"
	"fmt"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func NewDbConnect() (*DB, error) {
	dbConnStr := appUtil.GetEnv("db_connection")

	db, err := connect(dbConnStr)

	if err != nil {
		return nil, err
	}

	return &DB{
		db: db,
	}, nil
}

func connect(dsn string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	var dbName string = ""
	fmt.Println(dsn)
	fmt.Sscanf(dsn, "dbname=%s", &dbName)

	fmt.Println(dbName)
	fmt.Println(err)

	if err != nil {
		log.Panic(err)
		return nil, errors.New("cannot connect to postgres")
	}

	return db, nil
}
