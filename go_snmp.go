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
    host     = "10.10.0.117"
    user     = "postgres"
    password = "krnllnx"
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

    
    log.Println("connected postgres")
    DB.Logger = logger.Default.LogMode(logger.Silent)

    return DB
}
