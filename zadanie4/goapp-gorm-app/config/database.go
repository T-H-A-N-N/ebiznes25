package config

import (
    "log"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
    var err error
    DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Błąd połączenia z bazą danych:", err)
    }
    return DB
}
