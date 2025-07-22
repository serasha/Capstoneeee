package database

import (
    "log"
    "os"
    "gorm.io/driver/postgres"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    dbType := os.Getenv("DB_TYPE")
    if dbType == "" {
        dbType = "sqlite" // default to sqlite
    }

    var db *gorm.DB
    var err error

    switch dbType {
    case "postgres":
        dsn := "host=localhost user=postgres password=postgres dbname=latihan-capstone port=5432 sslmode=disable"
        db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    case "sqlite":
        dbPath := os.Getenv("DB_PATH")
        if dbPath == "" {
            dbPath = "./app.db"
        }
        db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
    default:
        log.Fatal("Unsupported database type:", dbType)
    }

    if err != nil {
        log.Fatal("Gagal Connect ke database:", err)
    }

    DB = db
    log.Println("Berhasil Connect ke database.")
}
