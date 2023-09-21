package dbpostgres

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

const (
    host     = "192.168.79.133"
    user     = "postgres"
    password = "postgres"
    dbname   = "postgres"
)
 
var DB *gorm.DB

func ConnectDB() *gorm.DB{
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
        host,
        user,
        password,
        dbname,
    )

    DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })

    if err != nil {
        log.Fatal("Failed to connect to database. \n", err)
        os.Exit(2)
    }

    
    log.Println("connected")
    DB.Logger = logger.Default.LogMode(logger.Info)

    return DB
}


func SelectCount(db *gorm.DB) int64 {

    var count int64
    db.Table("vpn").Count(&count)
    return count
}
