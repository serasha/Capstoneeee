package database

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "host=localhost user=postgres password=postgres dbname=latihan-capstone port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Gagal Connect ke database:", err)
    }

    DB = db
    log.Println("Berhasil Connect ke database.")
}
