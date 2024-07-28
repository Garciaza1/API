package db

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Init() {
    var err error
    dsn := ("host=localhost user=postgres password=Gr112495#$ dbname=go port=5432 sslmode=disable")

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Println(err)
        panic("failed to connect to database")
    }
    fmt.Println("Database connection established")
}