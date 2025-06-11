package db

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "cache_API/config"
    "log"
    "time"
)

var db *sql.DB

func InitDB() (*sql.DB, error) {
    cfg, err := config.LoadPostgresConfig()
    if err != nil {
        log.Fatalf("Config load failed: %v", err)
    }

    dsn := fmt.Sprintf(
        "postgres://%s:%s@%s:%d/%s?sslmode=%s",
        cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode,
    )



    for i := 0; i < 10; i++ {
        var err error
        db, err = sql.Open("postgres", dsn)
        fmt.Println("ping :", db.Ping())
        if err == nil && db.Ping() == nil {
            break
        }
        log.Println("DB not ready, retrying...")
        time.Sleep(2 * time.Second)
    }

    fmt.Println("ping :", db.Ping())

    return db, nil
}

func GetDB() *sql.DB {
    return db
}